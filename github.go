package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

const cacheFile = "cache.json"
const cacheDuration = 10 * time.Minute

type Event struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	CreatedAt string `json:"created_at"`
}

func getUserActivity(username string) ([]Event, error) {
	if data, err := readCache(); err == nil {
		return data, nil
	}
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	req, _ := http.NewRequest("GET", url, nil)
	token := os.Getenv("GITHUB_TOKEN")

	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch user activitzy: %s", resp.Status)
	}

	var events []Event

	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, err
	}

	writeCache(events)
	fmt.Println("🔄 Remaining API Requests:", resp.Header.Get("X-RateLimit-Remaining"))
	return events, nil
}

func readCache() ([]Event, error) {
	info, err := os.Stat(cacheFile)
	if err != nil || time.Since(info.ModTime()) > cacheDuration {
		return nil, fmt.Errorf("cache expired or does not exist")
	}
	data, err := os.ReadFile(cacheFile)
	if err != nil {
		return nil, err
	}

	var events []Event

	if err := json.Unmarshal(data, &events); err != nil {
		return nil, err
	}

	fmt.Println("✅ Using cached data.")
	return events, nil
}

func writeCache(events []Event) {
	data, _ := json.Marshal(events)
	os.WriteFile(cacheFile, data, 0644)
}
