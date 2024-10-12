package util

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func Pause() {
	fmt.Println("\nНажмите Enter для продолжения...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func GetDocumentsPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	documentsPath := filepath.Join(homeDir, "Documents")
	return documentsPath, nil
}
