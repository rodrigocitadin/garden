package main

import (
	"fmt"
	"os"
	"os/exec"
)

func runGitCommands(filePath, link string) error {
	commands := [][]string{
		{"git", "add", filePath},
		{"git", "commit", "-m", fmt.Sprintf("link added: %s", link)},
		{"git", "push"},
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
