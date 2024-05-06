package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Prompt the user to enter a search query
	fmt.Print("Enter a search query: ")
	reader := bufio.NewReader(os.Stdin)
	query, _ := reader.ReadString('\n')
	query = strings.TrimSpace(query)

	// Run yt-dlp command to search for videos based on the query
	cmd := exec.Command("yt-dlp", "-g", "--get-title", "-f", "bestaudio", "ytsearch3:"+query)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Convert command output to string
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")

	// Extract titles and URLs from output
	var titles []string
	var urls []string
	for i, line := range lines {
		if i%2 == 0 {
			titles = append(titles, line)
		} else {
			urls = append(urls, line)
		}
	}

	// Display search results
	fmt.Println("Search Results:")
	for i, title := range titles {
		fmt.Printf("%d. %s\n", i+1, title)
	}

	// Prompt the user to select a video to play
	fmt.Print("Enter the number of the video to play: ")
	var selection int
	_, err = fmt.Scanln(&selection)
	if err != nil || selection < 1 || selection > len(titles) {
		fmt.Println("Invalid selection.")
		return
	}

	// Get the selected audio URL
	selectedURL := urls[selection-1]

	// Run ffplay to stream the selected audio
	ffplayCmd := exec.Command("ffplay", "-nodisp", "-autoexit", "-loglevel", "quiet", "-stats", selectedURL)
	ffplayCmd.Stdout = os.Stdout
	ffplayCmd.Stderr = os.Stderr

	// Start streaming audio
	fmt.Println("Streaming audio...")
	if err := ffplayCmd.Run(); err != nil {
		fmt.Println("Error:", err)
		return
	}
}
