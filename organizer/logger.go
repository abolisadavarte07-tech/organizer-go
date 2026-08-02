package organizer

import (
	"fmt"
	"os"
	"path/filepath"
)

func LogMove(basePath string, fileName string, category string) error {

	logFile := filepath.Join(basePath, "Moved-Files-Log.txt")

	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprintf(f, "%s -> %s\n", fileName, category)
	return err
}