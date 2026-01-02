package main

import (
	"fmt"

	filesystem "composite.example/internal/filesystem"
)

func main() {
	root := filesystem.NewFolder("root")
	documents := filesystem.NewFolder("documents")
	downloads := filesystem.NewFolder("downloads")

	readme := filesystem.NewFile("Readme.txt", 100)
	notes := filesystem.NewFile("Notes.txt", 200)
	image := filesystem.NewFile("Image.png", 300)

	root.Add(documents)
	root.Add(downloads)
	root.Add(readme)

	documents.Add(notes)

	downloads.Add(image)

	fmt.Printf("Downloads size: %d\n", downloads.Size())
	fmt.Printf("Documents size: %d\n", documents.Size())
	fmt.Printf("Root size: %d\n", root.Size())
}
