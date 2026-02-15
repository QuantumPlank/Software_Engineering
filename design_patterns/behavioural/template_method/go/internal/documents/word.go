package documents

import "fmt"

type WordDocument struct{}

func (d *WordDocument) Open() {
	fmt.Println("Opening WORD document")
}

func (d *WordDocument) Parse() {
	fmt.Println("Opening WORD document")
}

func (d *WordDocument) Export() {
	fmt.Println("Export WORD document")
}

func (d *WordDocument) Close() {
	fmt.Println("Close WORD document")
}
