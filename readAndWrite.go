package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Database struct {
	ItemDatabase string
}

func (f *Database) ReadDatabase() (string, error) {
	file, err := os.OpenFile(f.ItemDatabase, os.O_CREATE, 0644)
	if err != nil {
		return "", fmt.Errorf("error found %s \n", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var store strings.Builder
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", fmt.Errorf("error found %s \n", err)
		}
		store.WriteString(line)
	}
	return store.String(), nil
}

func (f *Database) WriteDatabase(content string) error {
	file, err := os.OpenFile(f.ItemDatabase, os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error found %s \n", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	contents := content
	_, err = writer.WriteString(contents)
	if err != nil {
		return fmt.Errorf("error found %s \n", err)
	}
	writer.Flush()
	return nil
}
