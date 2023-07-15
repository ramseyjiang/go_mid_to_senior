package gofpdf

import (
	"bufio"
	"log"
	"os"

	"github.com/jung-kurt/gofpdf"
)

const paperSize = "A4"
const headerWidth = 190
const headerHeight = 20
const headerFontSize = 30
const headerContentDistance = 20
const bold = "B"
const fontStyle = "Arial"
const fontSize = 14
const contentWidth = 190
const contentHeight = 10
const contentsDistance = 10

func TxtConvertPDF(fileName string) {
	pdf := gofpdf.New("P", "mm", paperSize, "")
	pdf.AddPage()

	// Header
	pdf.SetFont(fontStyle, bold, headerFontSize)
	pdf.Cell(headerWidth, headerHeight, fileName)
	pdf.Ln(headerContentDistance) // New line

	// Content
	pdf.SetFont(fontStyle, "", fontSize)

	// Open the text file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("Error closing file: %v", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pdf.Cell(contentWidth, contentHeight, scanner.Text())
		pdf.Ln(contentsDistance) // New line after each line of the text
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Save the file
	err = pdf.OutputFileAndClose(fileName + ".pdf")
	if err != nil {
		log.Printf("Error saving file: %v", err)
		return
	}
}
