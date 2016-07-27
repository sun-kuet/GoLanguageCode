package main

import (
	"github.com/jung-kurt/gofpdf"
	"fmt"
	"os"
)

var baseDir string = os.Getenv("HOME")
var imageDir string = baseDir + "/Github/GoLanguageCode/gofpdf-test/images/"

type PdfCreator struct {
	pdf *gofpdf.Fpdf
}

func NewPdfCreator() *PdfCreator {
	return &PdfCreator{
		pdf: gofpdf.New("P", "mm", "A4", ""),
	}
}

func (p *PdfCreator) addImage(filename string) {
	p.pdf.SetY(10)
	p.pdf.Image(imageDir + filename, 100, 60, 30, 0, false, "", 0, "")
}

func main() {
	p := NewPdfCreator()
	p.pdf.AddPage()
	p.addImage("sun.jpg")

	p.pdf.SetFont("Arial", "B", 16)
	p.pdf.Cell(40, 10, "Hello World!")
	fileStr := "generated.pdf"
	err := p.pdf.OutputFileAndClose(fileStr)
	fmt.Println(err)
}
