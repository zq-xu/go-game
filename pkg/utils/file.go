package utils

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// LoadFilesWithSuffix
func LoadFilesWithSuffix(dir string, suffix string) ([]string, error) {
	list := make([]string, 0)
	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && strings.HasSuffix(strings.ToLower(d.Name()), suffix) {
			list = append(list, path)
		}
		return nil
	})
	return list, err
}

// LoadFilesInCurrentDirWithSuffix
func LoadFilesInCurrentDirWithSuffix(dir string, suffix string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	list := make([]string, 0)
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(strings.ToLower(entry.Name()), strings.ToLower(suffix)) {
			list = append(list, filepath.Join(dir, entry.Name()))
		}
	}

	return list, nil
}
