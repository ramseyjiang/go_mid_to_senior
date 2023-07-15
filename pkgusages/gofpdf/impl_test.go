package gofpdf

import (
	"os"
	"testing"
)

func TestTxtConvertPDF(t *testing.T) {
	// Create a dummy "input.txt" file
	inputFileName := "input.txt"
	file, err := os.Create(inputFileName)
	if err != nil {
		t.Fatalf("Failed to create input.txt: %v", err)
	}
	file.WriteString("This is a longer content that should be suitable for a width of 190.\n It includes several sentences to fill up the space.\n Remember that the exact number of characters that can fit in\n a width of 190 will depend on the font size and type.")
	file.Close()
	outputFileName := "Example"

	// Run the TxtConvertPDF function
	TxtConvertPDF(outputFileName, inputFileName)

	// Check if the PDF file is created
	if _, err := os.Stat(outputFileName + ".pdf"); os.IsNotExist(err) {
		t.Fatalf("Failed to create PDF: %v", err)
	}

	// Clean up
	os.Remove("input.txt")
	os.Remove(outputFileName + ".pdf")
	// Check if the PDF file is not deleted
	if _, err := os.Stat(outputFileName + ".pdf"); os.IsExist(err) {
		t.Fatalf("Failed to delete PDF: %v", err)
	}
}
