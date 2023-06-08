package apis

import (
	"log"
	"os"
)

func GetCurrentPath() string {
	currentPath, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	return currentPath
}

func GitFolderExists() bool {
	files, err := os.ReadDir(GetCurrentPath())

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.Name() == ".git" {
			return true
		}
	}

	return false
}
