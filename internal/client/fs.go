package client

import (
	"fmt"
	"os"
	"path/filepath"
)

func getDownloadsDir() string {
	if dir, ok := os.LookupEnv("USERPROFILE"); ok {
		return filepath.Join(dir, "Downloads")
	}
	if dir, ok := os.LookupEnv("HOME"); ok {
		return filepath.Join(dir, "Downloads")
	}
	return "./downloads"
}

func getUniqueFilepath(dir, filename string) string {
	ext := filepath.Ext(filename)
	name := filename[:len(filename)-len(ext)]

	finalPath := filepath.Join(dir, filename)

	i := 1
	for {
		_, err := os.Stat(finalPath)
		if os.IsNotExist(err) {
			return finalPath
		}

		newName := fmt.Sprintf("%s(%d)%s", name, i, ext)
		finalPath = filepath.Join(dir, newName)
		i++
	}
}
