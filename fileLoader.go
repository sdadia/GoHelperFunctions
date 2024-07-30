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

func Ls(directoryPath string) ([]os.DirEntry, error) {

	// Go through the directory
	directoryEntries, err := os.ReadDir(directoryPath)
	if err != nil {
		slog.Error("Error listing in directory : %v. Error is %v\n", directoryPath, err)
		return nil, err
	}

	return directoryEntries, nil
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

func HelloWorld() {

	slog.Info("Hello from module!")
}
