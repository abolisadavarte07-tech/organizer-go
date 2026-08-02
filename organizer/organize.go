package organizer

import (
    "fmt"
    "os"
    "path/filepath"
)

func ScanDirectory(path string) ([]FileInfo, error) {

	var files []FileInfo

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var IgnoredFiles = map[string]bool{
    	"Moved-Files-Log.txt": true,
    }

	for _, entry := range entries {

    	if entry.IsDir() {
	    	continue
	    }

    	if IgnoredFiles[entry.Name()] {
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

func Organize(path string, command string) error {

	files, err := ScanDirectory(path)
	if err != nil {
		return err
	}

	for _, file := range files {

		category, found := GetCategory(file.Extension)

		if !found {
			continue
		}

		allowed, ok := Commands[command]

        if !ok {
	        return fmt.Errorf("unsupported command: %s", command)
        }

        if !Contains(allowed, category) {
	        continue
        }

		err := MoveFile(path, file, category)
		if err != nil {
			return err
		}

		err = LogMove(path, file.Name, category)
        if err != nil {
        	return err
        }
	}

	return nil
}

func CountMovableFiles(path string, command string) (int, error) {
	files, err := ScanDirectory(path)
	if err != nil {
		return 0, err
	}

	count := 0

	allowed, ok := Commands[command]
	if !ok {
		return 0, fmt.Errorf("unsupported command: %s", command)
	}

	for _, file := range files {
		category, found := GetCategory(file.Extension)
		if !found {
			continue
		}

		if Contains(allowed, category) {
			count++
		}
	}

	return count, nil
}