package cmd

import (
	"os"
	"path/filepath"
	"strings"
)

func IsDir(filePath string) (bool, error) {
	f, err := os.Stat(filePath)
	if err != nil {
		return false, err
	}
	return f.IsDir(), nil
}

func IsExist(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	}
	if os.IsExist(err) {
		return true, nil
	}
	return false, err
}

func GetFileNameWithoutExt(filePath string) string {
	baseName := filepath.Base(filePath)
	baseFilePart := strings.Split(baseName, ".")
	if len(baseFilePart) > 1 {
		return strings.Join(baseFilePart[:len(baseFilePart)-1], ".")
	}
	return baseName
}
