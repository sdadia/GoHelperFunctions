package GoHelperFunctions

import (
	"bufio"
	"fmt"
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
