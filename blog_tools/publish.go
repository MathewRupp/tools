package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	// Set the directory for your Hugo blog repository
	repoDir := "/home/mathewrupp/programming/blog/mdrupp" // Change this to your actual repository path

	// Commit message with a timestamp
	now := time.Now()
	commitMessage := fmt.Sprintf("new post %s", now.Format("2006-01-02"))

	// Change to the repository directory
	err := os.Chdir(repoDir)
	if err != nil {
		fmt.Printf("Error changing directory: %v\n", err)
		return
	}

	// Run git add .
	err = runGitCommand("add", ".")
	if err != nil {
		return
	}

	// Run git commit -m "new post"
	err = runGitCommand("commit", "-m", commitMessage)
	if err != nil {
		return
	}

	// Run git push
	err = runGitCommand("push")
	if err != nil {
		return
	}

	fmt.Println("Blog successfully published!")
}

// runGitCommand runs a git command with the given arguments.
func runGitCommand(args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command and wait for it to finish
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running git command '%s': %v\n", args, err)
		return err
	}
	return nil
}



