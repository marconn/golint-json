package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// GetJSONFiles gets the filenames that matches
// the glob pattern *.json in the given path
func getJSONFiles(path string) []string {
	globPattern := fmt.Sprintf("%s/*.json", path)
	files, err := filepath.Glob(globPattern)
	if err != nil {
		log.Fatal(err)
	}

	return files
}

func isValidDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, fmt.Errorf("directory %s does not exist", path)
	}
	if !fileInfo.IsDir() {
		return false, fmt.Errorf("path %s is not a valid directory", path)
	}

	return true, nil
}

func readFile(path string) (string, error) {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(file), nil
}