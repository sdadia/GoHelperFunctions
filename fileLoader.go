package GoHelperFunctions

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

func ReadFile(fileName string) []string {

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened file")
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var newLine = scanner.Text()
		ans := strings.Split(newLine, "Search :")
		println(ans)
		os.Exit(-1)
	}

	return lines
}

func ls(directoryPath string) ([]string, []string, error) {

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
