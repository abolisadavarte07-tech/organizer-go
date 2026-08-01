package organizer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// MoveFile moves a file into Organized/<Category>.
// If a file with the same name already exists,
// it creates "filename - Copy.ext", "filename - Copy (2).ext", etc.
func MoveFile(basePath string, file FileInfo, category string) error {

	// Destination folder
	destFolder := filepath.Join(basePath, "Organized", category)

	// Create folder if it doesn't exist
	if err := os.MkdirAll(destFolder, os.ModePerm); err != nil {
		return err
	}

	// Source path
	source := filepath.Join(basePath, file.Name)

	// Destination path
	destination := filepath.Join(destFolder, file.Name)

	// If destination already exists, create a new filename
	if _, err := os.Stat(destination); err == nil {

		ext := filepath.Ext(file.Name)
		name := strings.TrimSuffix(file.Name, ext)

		i := 1

		for {
			var newName string

			if i == 1 {
				newName = fmt.Sprintf("%s - Copy%s", name, ext)
			} else {
				newName = fmt.Sprintf("%s - Copy (%d)%s", name, i, ext)
			}

			destination = filepath.Join(destFolder, newName)

			if _, err := os.Stat(destination); os.IsNotExist(err) {
				break
			}

			i++
		}
	}

	// Move file
	return os.Rename(source, destination)
}

func Contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}