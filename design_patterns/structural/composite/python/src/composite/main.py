from filesystem import File, Folder


def main():
    root = Folder("root")
    documents = Folder("documents")
    downloads = Folder("downloads")

    readme = File("readme.txt", 100)
    notes = File("notes.txt", 200)
    image = File("image.png", 300)

    root.Add(documents)
    root.Add(downloads)

    root.Add(readme)

    documents.Add(notes)

    downloads.Add(image)

    print(f"Documents: {documents.Size()}")
    print(f"Downloads: {downloads.Size()}")
    print(f"Root: {root.Size()}")


if __name__ == "__main__":
    main()
