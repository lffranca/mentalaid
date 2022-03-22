package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadTemplate(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, strings.Replace(path, fmt.Sprintf("%v/", root), "", -1))
		}

		return nil
	})

	return files, err
}
