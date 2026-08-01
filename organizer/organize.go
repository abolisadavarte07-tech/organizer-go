package organizer

import (
	"os"
	"path/filepath"
)

func ScanDirectory(path string) ([]FileInfo, error) {

	var files []FileInfo

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {

		if entry.IsDir() {
			continue
		}

		ext := filepath.Ext(entry.Name())

		files = append(files, FileInfo{
			Name:      entry.Name(),
			Extension: ext,
		})
	}

	return files, nil
}

func GetCategory(extension string) (string, bool) {
	category, found := ExtensionMap[extension]
	return category, found
}