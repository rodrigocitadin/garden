package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func runGitCommands(filePath, link string) error {
	repoDir := findGitRoot(filepath.Dir(filePath))
	if repoDir == "" {
		return fmt.Errorf("Git repository directory not found for %s", filePath)
	}

	relPath, err := filepath.Rel(repoDir, filePath)
	if err != nil {
		return err
	}

	commands := [][]string{
		{"git", "-C", repoDir, "add", relPath},
		{"git", "-C", repoDir, "commit", "-m", fmt.Sprintf("link added: %s", link)},
		{"git", "-C", repoDir, "push"},
	}

	for _, cmd := range commands {
		if err := runCommand(cmd); err != nil {
			return err
		}
	}
	return nil
}

func runCommand(args []string) error {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func findGitRoot(startDir string) string {
	current := startDir
	for {
		if _, err := os.Stat(filepath.Join(current, ".git")); err == nil {
			return current
		}
		parent := filepath.Dir(current)
		if parent == current {
			return ""
		}
		current = parent
	}
}
