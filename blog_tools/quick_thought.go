package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	// Get current date in "Month Day Year" format
	now := time.Now()
	title := now.Format("January 2, 2006")

	// Define default tag
	tag := "Quick Thought"

	// Define your posts directory
	postsDir := "/home/mathewrupp/programming/blog/mdrupp/content/posts" // Update if needed

	// Generate timestamp for directory name (without suffix)
	timestamp := now.Format("2006-01-02")
	postDirBase := filepath.Join(postsDir, timestamp+"-"+generateSlug(title))

	// Find unique directory name in case of multiple posts on the same day
	postDir, err := getUniquePostDir(postDirBase)
	if err != nil {
		fmt.Printf("Error finding unique post directory: %v\n", err)
		return
	}

	// Generate the full file path for index.md
	postFile := filepath.Join(postDir, "index.md")

	// Create post directory
	err = os.MkdirAll(postDir, 0755)
	if err != nil {
		fmt.Printf("Error creating post directory: %v\n", err)
		return
	}

	// Create the initial index.md file with metadata if it doesn't exist
	_, err = os.Stat(postFile)
	if os.IsNotExist(err) {
		// Write basic metadata to index.md (we'll add content later)
		f, err := os.Create(postFile)
		if err != nil {
			fmt.Printf("Error creating post file: %v\n", err)
			return
		}
		defer f.Close()

		// Write metadata to index.md
		_, err = f.WriteString(fmt.Sprintf(`+++
title = "%s"
date = "%s"
tags = ["%s"]
+++

`, title, now.Format(time.RFC3339), tag))
		if err != nil {
			fmt.Printf("Error writing to post file: %v\n", err)
			return
		}
	}

	// Open vim editor with a command to append content to the file
	cmd := exec.Command("vim", "-c", "normal! G", postFile) // Move cursor to end of file
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Run vim editor
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error opening vim: %v\n", err)
		return
	}

	// Notify user when finished
	fmt.Printf("Post content appended successfully at %s\n", postFile)
}

// generateSlug converts the title into a slug with lowercase and dashes
func generateSlug(title string) string {
	slug := strings.ToLower(strings.ReplaceAll(title, " ", "-"))
	return slug
}

// getUniquePostDir ensures the directory is unique for multiple posts per day
func getUniquePostDir(baseDir string) (string, error) {
	dir := baseDir
	count := 1
	// Check if the directory already exists; if so, increment a counter
	for {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			break
		}
		// Directory exists, increment count
		dir = fmt.Sprintf("%s-%02d", baseDir, count)
		count++
	}
	return dir, nil
}

