package gofpdf

import (
	"os"
	"testing"
)

func TestTxtConvertPDF(t *testing.T) {
	// Create a dummy "input.txt" file
	file, err := os.Create("input.txt")
	if err != nil {
		t.Fatalf("Failed to create input.txt: %v", err)
	}
	file.WriteString("This is a longer content that should be suitable for a width of 190.\n It includes several sentences to fill up the space.\n Remember that the exact number of characters that can fit in\n a width of 190 will depend on the font size and type.")
	file.Close()
	fileName := "Example"

	// Run the TxtConvertPDF function
	TxtConvertPDF(fileName)

	// Check if the PDF file is created
	if _, err := os.Stat(fileName + ".pdf"); os.IsNotExist(err) {
		t.Fatalf("Failed to create PDF: %v", err)
	}

	// Clean up
	os.Remove("input.txt")
	os.Remove(fileName + ".pdf")
}
