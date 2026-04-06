//go:build ignore

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Lesson struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

type Section struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Lessons []Lesson `json:"lessons"`
}

type Curriculum struct {
	Sections []Section `json:"sections"`
}

var runPathPattern = regexp.MustCompile(`\./[A-Za-z0-9._/\-]+`)

func isPlaceholderPath(path string) bool {
	placeholders := []string{
		"./SECTION/LESSON",
		"./path/to/",
		"./my-folder",
		"./my-messy-folder",
	}

	for _, placeholder := range placeholders {
		if strings.Contains(path, placeholder) {
			return true
		}
	}

	return false
}

func validateCurriculumPaths() (int, int) {
	data, err := os.ReadFile("curriculum.json")
	if err != nil {
		fmt.Printf("Failed to read curriculum.json: %v\n", err)
		os.Exit(1)
	}

	var cur Curriculum
	if err := json.Unmarshal(data, &cur); err != nil {
		fmt.Printf("Failed to parse curriculum.json: %v\n", err)
		os.Exit(1)
	}

	errors := 0
	lessonCount := 0
	for _, s := range cur.Sections {
		for _, l := range s.Lessons {
			lessonCount++
			if l.Path == "" {
				fmt.Printf("Unmapped lesson: %s - %s\n", l.ID, l.Name)
				errors++
				continue
			}

			if _, err := os.Stat(l.Path); os.IsNotExist(err) {
				fmt.Printf("Path does not exist: %s (%s - %s)\n", l.Path, l.ID, l.Name)
				errors++
			}
		}
	}

	return lessonCount, errors
}

func shouldScanRunPaths(path string) bool {
	if filepath.Ext(path) == ".go" {
		return true
	}

	if filepath.Base(path) != "README.md" {
		return false
	}

	first := strings.Split(filepath.ToSlash(path), "/")[0]
	if len(first) < 2 {
		return false
	}

	return first[0] >= '0' && first[0] <= '9' && first[1] >= '0' && first[1] <= '9'
}

func validateRunPaths() (int, int) {
	filesScanned := 0
	errors := 0

	walkErr := filepath.WalkDir(".", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			switch d.Name() {
			case ".git", "vendor":
				return filepath.SkipDir
			default:
				return nil
			}
		}

		cleanPath := filepath.Clean(path)
		if !shouldScanRunPaths(cleanPath) {
			return nil
		}

		file, err := os.Open(cleanPath)
		if err != nil {
			return err
		}
		defer file.Close()

		filesScanned++
		scanner := bufio.NewScanner(file)
		lineNo := 0
		for scanner.Scan() {
			lineNo++
			line := scanner.Text()
			if !strings.Contains(line, "go run ") && !strings.Contains(line, "go test ") {
				continue
			}

			for _, match := range runPathPattern.FindAllString(line, -1) {
				if match == "./..." || isPlaceholderPath(match) {
					continue
				}

				target := filepath.Clean(strings.TrimPrefix(match, "./"))
				alternateTarget := filepath.Clean(filepath.Join(filepath.Dir(cleanPath), strings.TrimPrefix(match, "./")))

				if _, err := os.Stat(target); err == nil {
					continue
				}

				if _, err := os.Stat(alternateTarget); err == nil {
					continue
				}

				if _, err := os.Stat(target); os.IsNotExist(err) {
					fmt.Printf("Invalid run path: %s:%d -> %s\n", cleanPath, lineNo, match)
					errors++
				}
			}
		}

		return scanner.Err()
	})
	if walkErr != nil {
		fmt.Printf("Failed to scan run paths: %v\n", walkErr)
		os.Exit(1)
	}

	return filesScanned, errors
}

func main() {
	lessonCount, pathErrors := validateCurriculumPaths()
	filesScanned, runErrors := validateRunPaths()
	errors := pathErrors + runErrors

	if errors == 0 {
		fmt.Printf("Success! All %d lessons mapped and %d files with run commands validated.\n", lessonCount, filesScanned)
		return
	}

	fmt.Printf("Found %d validation errors.\n", errors)
	os.Exit(1)
}
