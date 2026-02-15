from .document import Document


class WORDDocument(Document):
    def Open(self):
        print("Opening WORD document")

    def Parse(self):
        print("Parsing WORD document")

    def Export(self):
        print("Exporting WORD document")

    def Close(self):
        print("Closing WORD document")
