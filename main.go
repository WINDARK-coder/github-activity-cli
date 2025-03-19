package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if os.Getenv("GITHUB_TOKEN") == "" {
		fmt.Println("⚠️ Warning: No GitHub token detected. API requests will be limited to 60 per hour.")
		fmt.Println("Set your token with 'export GITHUB_TOKEN=your_token' (Mac/Linux) or '$env:GITHUB_TOKEN=\"your_token\"' (Windows)")
	}
	fmt.Println("Enter a github username")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	username := strings.TrimSpace(scanner.Text())

	if username == "" {
		fmt.Println("❌ Username cannot be empty!")
		return
	}

	fmt.Printf("\nFetching activity for username: %s\n", username)
	events, err := getUserActivity(username)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	uniqueEvents := getUniqueEventTypes(events)
	fmt.Println("\nAvailable event types:")
	for _, eventType := range uniqueEvents {
		fmt.Printf("- %s\n", eventType)
	}

	fmt.Println("Enter event type to filter (or press Enter to skip): ")
	scanner.Scan()
	filter := strings.TrimSpace(scanner.Text())

	displayActivity(events, filter)
}
