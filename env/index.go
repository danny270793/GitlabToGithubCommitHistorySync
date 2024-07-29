package env

import (
	"bufio"
	"os"
	"strings"
)

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func LoadFromFile(path string) error {
	if !fileExists(path) {
		return nil
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")
		if strings.HasPrefix(line, "#") {
			continue
		}

		if strings.Trim(line, " ") == "" {
			continue
		}

		keyValue := strings.SplitN(line, "=", 2)
		os.Setenv(keyValue[0], keyValue[1])
	}

	return nil
}
