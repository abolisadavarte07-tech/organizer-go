package organizer

import (
	"os"
	"path/filepath"
)

// MoveFile moves a file into Organized/<Category>
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

	// Move file
	return os.Rename(source, destination)
}