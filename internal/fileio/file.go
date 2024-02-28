package fileio

import (
	"path/filepath"
	"strings"
)

func AddSuffixToFilename(imgPath, suffix string) string {
	ext := filepath.Ext(imgPath)
	base := strings.TrimSuffix(imgPath, ext)
	newImgPath := base + suffix + ext
	return newImgPath
}

func ChangeFilenameExtension(imgPath, newExt string) string {
	ext := filepath.Ext(imgPath)
	base := strings.TrimSuffix(imgPath, ext)
	newImgPath := base + newExt
	return newImgPath
}
