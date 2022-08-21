package file

import (
	"bufio"
	"os"
)

// Read function to read file content.
// return array of string content split by lines.
func Read(fileName string) ([]string, error) {
	readFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = readFile.Close()
	}()

	var (
		response    = make([]string, 0)
		fileScanner = bufio.NewScanner(readFile)
	)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		response = append(response, fileScanner.Text())
	}
	return response, nil
}
