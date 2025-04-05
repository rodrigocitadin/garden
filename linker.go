package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func AppendLinkByMonth(filePath, link string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	month := time.Now().Format("January")
	sectionHeader := fmt.Sprintf("## %s", month)
	lines := strings.Split(string(content), "\n")

	var output []string
	var sectionFound bool
	var inserted bool

	for i := 0; i < len(lines); {
		line := lines[i]

		if strings.HasPrefix(line, sectionHeader) {
			sectionFound = true
			output = append(output, line)
			i++

			for i < len(lines) && lines[i] == "" {
				i++
			}

			output = append(output, "")
			output = append(output, fmt.Sprintf("- %s", link))

			for i < len(lines) && strings.HasPrefix(lines[i], "- ") {
				output = append(output, lines[i])
				i++
			}

			for i < len(lines) && lines[i] == "" {
				i++
			}

			inserted = true
			continue
		}

		output = append(output, line)
		i++
	}

	if !sectionFound {
		newSection := []string{
			sectionHeader,
			"",
			fmt.Sprintf("- %s", link),
			"",
		}
		output = append(newSection, output...)
		inserted = true
	}

	if !inserted {
		return fmt.Errorf("Unable to insert link")
	}

	return os.WriteFile(filePath, []byte(strings.Join(output, "\n")), 0644)
}
