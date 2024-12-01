package common

import (
	"os"
	"bufio"
	"strings"
)

func InputToStrings(input string) ([]string, error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lines = append(lines, line)
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
