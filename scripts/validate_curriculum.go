//go:build ignore

package main

import (
	"bufio"
	"encoding/json"
	"errors"
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

type V2Section struct {
	ID            string   `json:"id"`
	Number        string   `json:"number"`
	Slug          string   `json:"slug"`
	Title         string   `json:"title"`
	PathPrefix    string   `json:"path_prefix"`
	EntryPoints   []string `json:"entry_points"`
	Outputs       []string `json:"outputs"`
	Prerequisites []string `json:"prerequisites"`
}

type V2Item struct {
	ID               string   `json:"id"`
	SectionID        string   `json:"section_id"`
	Slug             string   `json:"slug"`
	Title            string   `json:"title"`
	Type             string   `json:"type"`
	Subtype          string   `json:"subtype"`
	Level            string   `json:"level"`
	VerificationMode string   `json:"verification_mode"`
	Path             string   `json:"path"`
	Prerequisites    []string `json:"prerequisites"`
	RunCommand       string   `json:"run_command"`
	TestCommand      string   `json:"test_command"`
	StarterPath      string   `json:"starter_path"`
	NextItems        []string `json:"next_items"`
}

type V2Curriculum struct {
	SchemaVersion int         `json:"schema_version"`
	Sections      []V2Section `json:"sections"`
	Items         []V2Item    `json:"items"`
}

var runPathPattern = regexp.MustCompile(`\./[A-Za-z0-9._/\-]+`)

var (
	allowedItemTypes = map[string]bool{
		"lesson":       true,
		"drill":        true,
		"exercise":     true,
		"checkpoint":   true,
		"mini_project": true,
		"capstone":     true,
		"reference":    true,
	}
	allowedLessonSubtypes = map[string]bool{
		"concept":     true,
		"pattern":     true,
		"integration": true,
	}
	allowedLevels = map[string]bool{
		"foundation": true,
		"core":       true,
		"stretch":    true,
		"production": true,
	}
	allowedVerificationModes = map[string]bool{
		"run":    true,
		"test":   true,
		"rubric": true,
		"mixed":  true,
	}
)

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

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func extractCommandTarget(command string) (string, error) {
	command = strings.TrimSpace(command)
	if command == "" {
		return "", errors.New("command is empty")
	}

	match := runPathPattern.FindString(command)
	if match == "" || match == "./..." || isPlaceholderPath(match) {
		return "", fmt.Errorf("command does not contain a concrete ./path target: %q", command)
	}

	return filepath.Clean(strings.TrimPrefix(match, "./")), nil
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

func validateV2Curriculum() (int, int, int, bool) {
	if _, err := os.Stat("curriculum.v2.json"); os.IsNotExist(err) {
		return 0, 0, 0, false
	}

	data, err := os.ReadFile("curriculum.v2.json")
	if err != nil {
		fmt.Printf("Failed to read curriculum.v2.json: %v\n", err)
		os.Exit(1)
	}

	var cur V2Curriculum
	if err := json.Unmarshal(data, &cur); err != nil {
		fmt.Printf("Failed to parse curriculum.v2.json: %v\n", err)
		os.Exit(1)
	}

	errorsFound := 0
	sectionIDs := make(map[string]V2Section, len(cur.Sections))
	itemIDs := make(map[string]V2Item, len(cur.Items))

	for _, s := range cur.Sections {
		if s.ID == "" {
			fmt.Println("Invalid v2 section: missing id")
			errorsFound++
			continue
		}

		if _, exists := sectionIDs[s.ID]; exists {
			fmt.Printf("Duplicate v2 section id: %s\n", s.ID)
			errorsFound++
			continue
		}

		if s.Number == "" || s.Slug == "" || s.Title == "" || s.PathPrefix == "" {
			fmt.Printf("Invalid v2 section metadata: %s requires number, slug, title, and path_prefix\n", s.ID)
			errorsFound++
		}

		if s.PathPrefix != "" && !pathExists(s.PathPrefix) {
			fmt.Printf("Invalid v2 section path_prefix: %s -> %s\n", s.ID, s.PathPrefix)
			errorsFound++
		}

		sectionIDs[s.ID] = s
	}

	for _, item := range cur.Items {
		if item.ID == "" {
			fmt.Println("Invalid v2 item: missing id")
			errorsFound++
			continue
		}

		if _, exists := itemIDs[item.ID]; exists {
			fmt.Printf("Duplicate v2 item id: %s\n", item.ID)
			errorsFound++
			continue
		}

		if item.SectionID == "" || item.Slug == "" || item.Title == "" || item.Type == "" || item.Level == "" || item.VerificationMode == "" || item.Path == "" {
			fmt.Printf("Invalid v2 item metadata: %s requires section_id, slug, title, type, level, verification_mode, and path\n", item.ID)
			errorsFound++
		}

		if _, exists := sectionIDs[item.SectionID]; !exists {
			fmt.Printf("Invalid v2 section linkage: %s -> %s\n", item.ID, item.SectionID)
			errorsFound++
		}

		if !allowedItemTypes[item.Type] {
			fmt.Printf("Invalid v2 item type: %s -> %s\n", item.ID, item.Type)
			errorsFound++
		}

		if item.Type == "lesson" {
			if !allowedLessonSubtypes[item.Subtype] {
				fmt.Printf("Invalid v2 lesson subtype: %s -> %s\n", item.ID, item.Subtype)
				errorsFound++
			}
		} else if item.Subtype != "" {
			fmt.Printf("Unexpected v2 subtype for non-lesson item: %s -> %s\n", item.ID, item.Subtype)
			errorsFound++
		}

		if !allowedLevels[item.Level] {
			fmt.Printf("Invalid v2 item level: %s -> %s\n", item.ID, item.Level)
			errorsFound++
		}

		if !allowedVerificationModes[item.VerificationMode] {
			fmt.Printf("Invalid v2 verification mode: %s -> %s\n", item.ID, item.VerificationMode)
			errorsFound++
		}

		if !pathExists(item.Path) {
			fmt.Printf("Invalid v2 item path: %s -> %s\n", item.ID, item.Path)
			errorsFound++
		}

		if section, exists := sectionIDs[item.SectionID]; exists && section.PathPrefix != "" {
			itemPath := filepath.ToSlash(filepath.Clean(item.Path))
			sectionPrefix := filepath.ToSlash(filepath.Clean(section.PathPrefix))
			if itemPath != sectionPrefix && !strings.HasPrefix(itemPath, sectionPrefix+"/") {
				fmt.Printf("Invalid v2 section path alignment: %s -> %s (expected prefix %s)\n", item.ID, item.Path, section.PathPrefix)
				errorsFound++
			}
		}

		if item.RunCommand != "" {
			target, err := extractCommandTarget(item.RunCommand)
			if err != nil {
				fmt.Printf("Invalid v2 run command: %s -> %v\n", item.ID, err)
				errorsFound++
			} else if !pathExists(target) {
				fmt.Printf("Invalid v2 run command target: %s -> %s\n", item.ID, item.RunCommand)
				errorsFound++
			}
		}

		if item.TestCommand != "" {
			target, err := extractCommandTarget(item.TestCommand)
			if err != nil {
				fmt.Printf("Invalid v2 test command: %s -> %v\n", item.ID, err)
				errorsFound++
			} else if !pathExists(target) {
				fmt.Printf("Invalid v2 test command target: %s -> %s\n", item.ID, item.TestCommand)
				errorsFound++
			}
		}

		switch item.VerificationMode {
		case "run":
			if strings.TrimSpace(item.RunCommand) == "" {
				fmt.Printf("Invalid v2 run contract: %s requires run_command\n", item.ID)
				errorsFound++
			}
		case "test":
			if strings.TrimSpace(item.TestCommand) == "" {
				fmt.Printf("Invalid v2 test contract: %s requires test_command\n", item.ID)
				errorsFound++
			}
		case "mixed":
			if strings.TrimSpace(item.RunCommand) == "" && strings.TrimSpace(item.TestCommand) == "" {
				fmt.Printf("Invalid v2 mixed contract: %s requires run_command or test_command\n", item.ID)
				errorsFound++
			}
		}

		if item.StarterPath != "" && !pathExists(item.StarterPath) {
			fmt.Printf("Invalid v2 starter path: %s -> %s\n", item.ID, item.StarterPath)
			errorsFound++
		}

		itemIDs[item.ID] = item
	}

	for _, s := range cur.Sections {
		for _, prereqID := range s.Prerequisites {
			if _, exists := sectionIDs[prereqID]; !exists {
				fmt.Printf("Invalid v2 section prerequisite: %s -> %s\n", s.ID, prereqID)
				errorsFound++
			}
		}

		for _, entryID := range s.EntryPoints {
			if _, exists := itemIDs[entryID]; !exists {
				fmt.Printf("Invalid v2 section entry point: %s -> %s\n", s.ID, entryID)
				errorsFound++
			}
		}

		for _, outputID := range s.Outputs {
			if _, exists := itemIDs[outputID]; !exists {
				fmt.Printf("Invalid v2 section output: %s -> %s\n", s.ID, outputID)
				errorsFound++
			}
		}
	}

	for _, item := range cur.Items {
		for _, prereqID := range item.Prerequisites {
			if _, itemExists := itemIDs[prereqID]; itemExists {
				continue
			}
			if _, sectionExists := sectionIDs[prereqID]; sectionExists {
				continue
			}
			fmt.Printf("Invalid v2 prerequisite: %s -> %s\n", item.ID, prereqID)
			errorsFound++
		}

		for _, nextID := range item.NextItems {
			if _, itemExists := itemIDs[nextID]; itemExists {
				continue
			}
			if _, sectionExists := sectionIDs[nextID]; sectionExists {
				continue
			}
			fmt.Printf("Invalid v2 next item: %s -> %s\n", item.ID, nextID)
			errorsFound++
		}
	}

	return len(cur.Sections), len(cur.Items), errorsFound, true
}

func main() {
	lessonCount, pathErrors := validateCurriculumPaths()
	filesScanned, runErrors := validateRunPaths()
	v2SectionCount, v2ItemCount, v2Errors, hasV2 := validateV2Curriculum()
	errors := pathErrors + runErrors + v2Errors

	if errors == 0 {
		if hasV2 {
			fmt.Printf("Success! All %d lessons mapped, %d files with run commands validated, and %d v2 sections plus %d v2 items checked.\n", lessonCount, filesScanned, v2SectionCount, v2ItemCount)
			return
		}
		fmt.Printf("Success! All %d lessons mapped and %d files with run commands validated.\n", lessonCount, filesScanned)
		return
	}

	fmt.Printf("Found %d validation errors.\n", errors)
	os.Exit(1)
}
