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

func Organize(path string) error {

	files, err := ScanDirectory(path)
	if err != nil {
		return err
	}

	for _, file := range files {

		category, found := GetCategory(file.Extension)

		if !found {
			continue
		}

		err := MoveFile(path, file, category)
		if err != nil {
			return err
		}
	}

	return nil
}