package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

var csvData = [][]string{
	{"a", "b", "c"},
	{"1", "10", "100"},
	{"2", "20", "200"},
}

func main() {
	csvWriter()
	csvReaderAll()
}

func csvWriter() {
	recordFile, err := os.Create("1c.csv")
	if err != nil {
		return
	}

	writer := csv.NewWriter(recordFile)
	err = writer.WriteAll(csvData)
	if err != nil {
		return
	}

	err = recordFile.Close()
	if err != nil {
		return
	}
}
func csvReaderAll() {
	recordFile, err := os.Open("1c.csv")
	if err != nil {
		return
	}

	reader := csv.NewReader(recordFile)
	allRecords, err := reader.ReadAll()
	if err != nil {
		return
	}
	fmt.Println(allRecords)

	err = recordFile.Close()
	if err != nil {
		return
	}
}
