package gofpdf

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"rsc.io/pdf"
)

func TestTxtConvertPDF(t *testing.T) {
	expectedContent := "This is a longer content that should be suitable for a width of 190.\n It includes several sentences to fill up the space.\n Remember that the exact number of characters that can fit in\n a width of 190 will depend on the font size and type."
	outputFileName := "Example"
	inputFileName := "input.txt"

	t.Run("Test PDF Creation and Deletion", func(t *testing.T) {
		// Create a dummy "input.txt" file

		file, err := os.Create(inputFileName)
		if err != nil {
			t.Fatalf("Failed to create input.txt: %v", err)
		}

		file.WriteString(expectedContent)
		file.Close()

		// Run the TxtConvertPDF function
		TxtConvertPDF(outputFileName, inputFileName)

		// Check if the PDF file is created
		pdfFileName := outputFileName + ".pdf"
		if _, err := os.Stat(pdfFileName); os.IsNotExist(err) {
			t.Fatalf("Failed to create PDF: %v", err)
		}

		// Clean up
		os.Remove(inputFileName)
		os.Remove(pdfFileName)

		// Check if the PDF file is not deleted
		if _, err := os.Stat(pdfFileName); os.IsExist(err) {
			t.Fatalf("Failed to delete PDF: %v", err)
		}
	})

	t.Run("Test PDF Content", func(t *testing.T) {
		file, err := os.Create(inputFileName)
		if err != nil {
			t.Fatalf("Failed to create input.txt: %v", err)
		}

		file.WriteString(expectedContent)
		file.Close()

		// Run the TxtConvertPDF function
		TxtConvertPDF(outputFileName, inputFileName)

		// Check the content of the PDF file
		pdfFileName := outputFileName + ".pdf"
		file, err = os.Open(pdfFileName)
		if err != nil {
			t.Fatalf("Failed to open PDF: %v", err)
		}
		defer file.Close()

		fileInfo, err := file.Stat()
		if err != nil {
			t.Fatalf("Failed to get file info: %v", err)
		}

		r, err := pdf.NewReader(file, fileInfo.Size())
		if err != nil {
			t.Fatalf("Failed to create PDF reader: %v", err)
		}

		var pdfContent string
		totalPage := r.NumPage()
		for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
			p := r.Page(pageIndex)
			if p.V.IsNull() {
				continue
			}
			content := p.Content()
			for _, txt := range content.Text {
				pdfContent += txt.S
			}
		}

		pdfContent = strings.ReplaceAll(pdfContent, " ", "")
		pdfContent = strings.ReplaceAll(pdfContent, "\n", "")
		expectedContent = strings.ReplaceAll(expectedContent, " ", "")
		expectedContent = strings.ReplaceAll(expectedContent, "\n", "")
		expectedContent = "Example" + expectedContent

		if pdfContent != expectedContent {
			fmt.Printf("PDF Content: %s\n", pdfContent)
			fmt.Printf("Expected Content: %s\n", expectedContent)
			t.Fatalf("PDF content does not match expected content")
		}

		// Clean up
		os.Remove(inputFileName)
		os.Remove(pdfFileName)
	})
}
