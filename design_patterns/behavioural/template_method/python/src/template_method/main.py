from documents import HTMLDocument, PDFDocument, WORDDocument


def main():
    word = WORDDocument()
    word.ProcessDocument()

    pdf = PDFDocument()
    pdf.ProcessDocument()

    html = HTMLDocument()
    html.ProcessDocument()


if __name__ == "__main__":
    main()
