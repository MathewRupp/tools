package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Prompt for post metadata
	fmt.Print("Enter post title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Enter tags (comma-separated): ")
	tagsInput, _ := reader.ReadString('\n')
	tags := strings.Split(strings.TrimSpace(tagsInput), ",")

	// Define your posts directory
	postsDir := "/Users/mathewrupp/programming/blog/mdrupp.com/content/posts"

	// Generate timestamp for directory name and filename
	now := time.Now()
	timestamp := now.Format("2006-01-02")
	postDir := filepath.Join(postsDir, timestamp+"-"+generateSlug(title))
	postFile := filepath.Join(postDir, "index.md")

	// Create post directory
	err := os.MkdirAll(postDir, 0755)
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

`, title, now.Format(time.RFC3339), formatTags(tags)))
	if err != nil {
		fmt.Printf("Error writing to post file: %v\n", err)
		return
	}

	fmt.Printf("Post created successfully at %s\n", postFile)
}

func generateSlug(title string) string {
	// Replace spaces with dashes and lowercase
	slug := strings.ToLower(strings.ReplaceAll(title, " ", "-"))
	return slug
}

func formatTags(tags []string) string {
	// Format tags for front matter
	formatted := ""
	for _, tag := range tags {
		formatted += fmt.Sprintf(`"%s", `, strings.TrimSpace(tag))
	}
	// Remove trailing comma and space
	formatted = strings.TrimSuffix(formatted, ", ")
	return formatted
}
