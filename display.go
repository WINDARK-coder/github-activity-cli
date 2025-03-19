package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/fatih/color"
)

func displayActivity(events []Event, filter string) {
	filteredCount := 0

	pushColor := color.New(color.FgGreen).SprintFunc()
	issueColor := color.New(color.FgCyan).SprintFunc()
	watchColor := color.New(color.FgYellow).SprintFunc()
	forkColor := color.New(color.FgMagenta).SprintFunc()
	createColor := color.New(color.FgBlue).SprintFunc()
	prColor := color.New(color.FgHiRed).SprintFunc()
	deleteColor := color.New(color.FgHiBlack).SprintFunc()
	releaseColor := color.New(color.FgHiGreen).SprintFunc()

	for _, event := range events {
		if filter != "" && event.Type != filter {
			continue
		}

		timestamp, _ := time.Parse(time.RFC3339, event.CreatedAt)
		formattedTime := timestamp.Format("Jan 02 15:04")

		switch event.Type {
		case "PushEvent":
			fmt.Printf("%s %s Pushed to %s\n", formattedTime, pushColor("🔄"), event.Repo.Name)
		case "IssuesEvent":
			fmt.Printf("%s %s Opened a new issue in %s\n", formattedTime, issueColor("📝"), event.Repo.Name)
		case "WatchEvent":
			fmt.Printf("%s %s Starred %s\n", formattedTime, watchColor("⭐"), event.Repo.Name)
		case "ForkEvent":
			fmt.Printf("%s %s Forked %s\n", formattedTime, forkColor("🍴"), event.Repo.Name)
		case "CreateEvent":
			fmt.Printf("%s %s Created %s\n", formattedTime, createColor("📂"), event.Repo.Name)
		case "PullRequestEvent":
			fmt.Printf("%s %s Opened a pull request in %s\n", formattedTime, prColor("🔀"), event.Repo.Name)
		case "DeleteEvent":
			fmt.Printf("%s %s Deleted something in %s\n", formattedTime, deleteColor("❌"), event.Repo.Name)
		case "ReleaseEvent":
			fmt.Printf("%s %s Published a release in %s\n", formattedTime, releaseColor("🚀"), event.Repo.Name)
		default:
			fmt.Printf("%s 🔍 Other event type: %s in %s\n", formattedTime, event.Type, event.Repo.Name)
		}
		filteredCount++
	}
	if filter != "" && filteredCount == 0 {
		fmt.Printf("❌ No events found for filter: %s\n", filter)
	}
}

func getUniqueEventTypes(events []Event) []string {
	eventMap := make(map[string]bool)

	for _, event := range events {
		eventMap[event.Type] = true
	}

	var uniqueEvents []string
	for eventType := range eventMap {
		uniqueEvents = append(uniqueEvents, eventType)
	}
	sort.Strings(uniqueEvents)
	return uniqueEvents
}
