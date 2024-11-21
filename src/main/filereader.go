/*
********************
Last name: Bernandino
Language: Go
Paradigm(s): Procedural, Multi-paradigm
********************
*/
package main

import (
	"encoding/csv"
	"os"
)

func ReadCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	reader := csv.NewReader(file)
	reader.Read()	// Skips the header
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
