package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var defaultIconDir string

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	defaultIconDir = filepath.Join(home, ".icons")
}

func main() {
	var result []string
	files, err := findFiles(defaultIconDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	result = append(result, files...)
	files, _ = findFiles("/usr/share/icons")
	result = append(result, files...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	var m = make(map[string]bool)
	for index := range result {
		if isIcons(result[index]) {
			m[getIconName(result[index])] = true
		}
	}
	for icon := range m {
		fmt.Printf("%s\n", icon)
	}
}

func getIconName(filename string) string {
	return strings.TrimRight(filepath.Base(filename), ".png")
}

func isIcons(filename string) bool {
	return strings.HasSuffix(filename, ".png")
}

func findFiles(dir string) ([]string, error) {
	fh, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	files, err := fh.Readdir(0)
	if err != nil {
		return nil, err
	}
	var result []string
	for _, file := range files {
		path := filepath.Join(dir, file.Name())
		if file.IsDir() {
			subFile, _ := findFiles(path)
			result = append(result, subFile...)
		} else {
			result = append(result, path)
		}
	}
	return result, nil
}
