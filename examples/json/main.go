package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Data struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func LoadFromFile(data any, filePath string) error {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	err = json.Unmarshal(fileContent, data)
	if err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	return nil
}

func main() {

}
