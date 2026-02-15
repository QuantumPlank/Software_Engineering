package documents

import "fmt"

type PDFDocument struct{}

func (d *PDFDocument) Open() {
	fmt.Println("Opening PDF document")
}

func (d *PDFDocument) Parse() {
	fmt.Println("Opening PDF document")
}

func (d *PDFDocument) Export() {
	fmt.Println("Export PDF document")
}

func (d *PDFDocument) Close() {
	fmt.Println("Close PDF document")
}
