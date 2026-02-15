package documents

type DocumentProcessor interface {
	Open()
	Parse()
	Export()
	Close()
}

func ProcessDocument(p DocumentProcessor) {
	p.Open()
	p.Parse()
	p.Export()
	p.Close()
}
