from .document import Document


class PDFDocument(Document):
    def Open(self):
        print("Opening PDF document")

    def Parse(self):
        print("Parsing PDF document")

    def Export(self):
        print("Exporting PDF document")

    def Close(self):
        print("Closing PDF document")
