from .document import Document


class HTMLDocument(Document):
    def Open(self):
        print("Opening HTML document")

    def Parse(self):
        print("Parsing HTML document")

    def Export(self):
        print("Exporting HTML document")

    def Close(self):
        print("Closing HTML document")
