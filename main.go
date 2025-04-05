package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	link := flag.String("link", "", "Link to be added")
	file := flag.String("file", "", "Path to markdown file")
	setDefault := flag.String("set-default", "", "Set default path of markdown file")
	flag.Parse()

	if *setDefault != "" {
		cfg := LoadOrCreateConfig()
		cfg.DefaultFile = *setDefault
		saveConfig(cfg)
		fmt.Println("Default file updated successfully!")
		os.Exit(0)
	}

	if *link == "" {
		fmt.Println("Usege:\n  garden -link https://example.com -file markdown.md\n  garden --set-default markdown.md")
		os.Exit(1)
	}

	cfg := LoadOrCreateConfig()

	mdFile := *file
	if mdFile == "" {
		mdFile = cfg.DefaultFile
	}

	if mdFile == "" {
		fmt.Println("Error: No default markdown file defined. Use -file or --set-default.")
		os.Exit(1)
	}

	err := AppendLinkByMonth(mdFile, *link)
	if err != nil {
		fmt.Println("Error adding link:", err)
		os.Exit(1)
	}

	err = runGitCommands(mdFile, *link)
	if err != nil {
		fmt.Println("Git error:", err)
		os.Exit(1)
	}

	fmt.Println("Link added successfully!")
}
