package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/chzyer/readline"
)

func main() {
	// Prompt for post title
	fmt.Print("Enter post title: ")
	title, err := readInputWithBackspace("Post title: ")
	if err != nil {
		fmt.Printf("Error reading title: %v\n", err)
		return
	}
	title = strings.TrimSpace(title)

	// Remove date-related part from title
	titleWithoutDate := removeDateFromTitle(title)

	// Prompt for tags
	fmt.Print("Enter tags (comma-separated): ")
	tagsInput, err := readInputWithBackspace("Tags (comma-separated): ")
	if err != nil {
		fmt.Printf("Error reading tags: %v\n", err)
		return
	}

	// Split tags by comma
	tags := strings.Split(tagsInput, ",")

	// Define your posts directory
	postsDir := "/home/mathewrupp/programming/blog/mdrupp/content/posts" // Update if needed

	// Generate timestamp for directory name, but not for title
	now := time.Now()
	timestamp := now.Format("2006-01-02")
	postDir := filepath.Join(postsDir, timestamp+"-"+generateSlug(title))
	postFile := filepath.Join(postDir, "index.md")

	// Create post directory
	err = os.MkdirAll(postDir, 0755)
	if err != nil {
		fmt.Printf("Error creating post directory: %v\n", err)
		return
	}

	// Create and write to index.md file
	f, err := os.Create(postFile)
	if err != nil {
		fmt.Printf("Error creating post file: %v\n", err)
		return
	}
	defer f.Close()

	// Write metadata and title to index.md
	_, err = f.WriteString(fmt.Sprintf(`+++
title = "%s"
date = "%s"
tags = [%s]
+++

`, titleWithoutDate, now.Format(time.RFC3339), formatTags(tags)))
	if err != nil {
		fmt.Printf("Error writing to post file: %v\n", err)
		return
	}

	fmt.Printf("Post created successfully at %s\n", postFile)
}

// generateSlug converts the title into a slug with lowercase and dashes
func generateSlug(title string) string {
	slug := strings.ToLower(strings.ReplaceAll(title, " ", "-"))
	return slug
}

// removeDateFromTitle removes the date (YYYY-MM-DD) from the title if present
func removeDateFromTitle(title string) string {
	// Regular expression to match and remove date in YYYY-MM-DD format from the beginning of the title
	re := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}[-\s]+`)
	return re.ReplaceAllString(title, "")
}

// formatTags formats the tags into the correct format for front matter
func formatTags(tags []string) string {
	formatted := ""
	for _, tag := range tags {
		formatted += fmt.Sprintf(`"%s", `, strings.TrimSpace(tag))
	}
	// Remove trailing comma and space
	formatted = strings.TrimSuffix(formatted, ", ")
	return formatted
}

// readInputWithBackspace uses readline to handle backspace properly
func readInputWithBackspace(prompt string) (string, error) {
	// Create a new readline instance with a dynamic prompt
	rl, err := readline.New(prompt)
	if err != nil {
		return "", err
	}
	defer rl.Close()

	// Read user input
	line, err := rl.Readline()
	if err != nil {
		return "", err
	}

	// Return the input with backspace handled correctly
	return line, nil
}

