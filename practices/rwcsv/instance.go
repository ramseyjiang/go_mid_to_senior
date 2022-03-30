package rwcsv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Employee struct {
	ID  string
	Age int
}

const filePath string = "practices/rwcsv/"
const fileName string = "records.csv"

// Trigger is the entry, run it in the go_mid_to_senior folder, if in the another folder, the filePath should be update
func Trigger() {
	writeCSV()
	readCSV()
}

func readCSV() {
	file, err := os.Open(filePath + fileName)
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	fmt.Println(records)
}

func writeCSV() {
	records := []Employee{
		{"E01", 25},
		{"E02", 26},
		{"E03", 24},
		{"E04", 26},
	}
	file, err := os.Create(filePath + fileName)
	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	defer func(file *os.File) {
		err1 := file.Close()
		if err1 != nil {
			fmt.Println(err1)
		}
	}(file)

	w := csv.NewWriter(file)
	defer w.Flush()

	// Using Write
	for _, record := range records {
		row := []string{record.ID, strconv.Itoa(record.Age)}
		if err3 := w.Write(row); err3 != nil {
			log.Println("error writing record to file", err3)
		}
	}

	// Using WriteAll
	var data [][]string
	for _, record := range records {
		row := []string{record.ID, strconv.Itoa(record.Age)}
		data = append(data, row)
	}
	err2 := w.WriteAll(data)
	if err2 != nil {
		return
	}
}
