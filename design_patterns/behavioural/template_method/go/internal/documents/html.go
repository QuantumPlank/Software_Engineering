package documents

import "fmt"

type HTMLDocument struct{}

func (d *HTMLDocument) Open() {
	fmt.Println("Opening HTML document")
}

func (d *HTMLDocument) Parse() {
	fmt.Println("Opening HTML document")
}

func (d *HTMLDocument) Export() {
	fmt.Println("Export HTML document")
}

func (d *HTMLDocument) Close() {
	fmt.Println("Close HTML document")
}
