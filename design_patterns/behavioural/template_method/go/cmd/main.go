package main

import "template_method.example/internal/documents"

func main() {
	word := &documents.WordDocument{}
	documents.ProcessDocument(word)

	pdf := &documents.PDFDocument{}
	documents.ProcessDocument(pdf)

	html := &documents.HTMLDocument{}
	documents.ProcessDocument(html)
}
