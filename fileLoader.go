package GoHelperFunctions

import (
	"log/slog"
	"os"
	"path/filepath"
)

func ReadFile(fileName string) (string, error) {

	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	return string(bytes), nil

}

func listDirectory(directoryPath string) ([]string, []string, error) {

	var files = []string{}
	var directories = []string{}

	directoryEntries, err := os.ReadDir(directoryPath)
	if err != nil {
		slog.Error("Error listing in directory : %v. Error is %v\n", directoryPath, err)
		return nil, nil, err
	}

	for _, entry := range directoryEntries {
		if entry.IsDir() {
			directories = append(directories, entry.Name())
		} else {
			files = append(files, entry.Name())
		}
	}

	return files, directories, nil
}

func ListFilesRecursively(dir string) ([]string, error) {

	var files = []string{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
