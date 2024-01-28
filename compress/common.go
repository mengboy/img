package compress

import (
	"path/filepath"
	"strings"
)

func GetExtWithoutPoint(filePath string) string {
	ext := strings.ToLower(filepath.Ext(filePath))
	if strings.HasPrefix(ext, ".") {
		return ext[1:]
	}
	return ext
}
