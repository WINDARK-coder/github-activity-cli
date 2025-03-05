package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Event struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please enter a github username")
		return
	}

	username := os.Args[1]

	fmt.Printf("Fetching activity for username: %s\n", username)
	events, err := getUserActivity(username)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	displayActivity(events)
}

func getUserActivity(username string) ([]Event, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	resp, err := http.Get(url)
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

	return events, nil

}

func displayActivity(events []Event) {
	for _, event := range events {
		switch event.Type {
		case "PushEvent":
			fmt.Printf("🔄 Pushed to %s\n", event.Repo.Name)
		case "IssuesEvent":
			fmt.Printf("📝 Opened a new issue in %s\n", event.Repo.Name)
		case "WatchEvent":
			fmt.Printf("⭐ Starred %s\n", event.Repo.Name)
		default:
			fmt.Printf("🔍 Other event type: %s in %s\n", event.Type, event.Repo.Name)
		}
	}
}
