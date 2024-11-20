package main

import (
	"encoding/csv"
	"os"
)

func ReadCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)

	defer file.Close()

	if err != nil {
		return nil, err
	}
	
	reader := csv.NewReader(file)
	reader.Read()	// Skips the header
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
