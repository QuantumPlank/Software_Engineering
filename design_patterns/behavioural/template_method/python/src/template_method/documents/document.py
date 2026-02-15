from abc import ABC, abstractmethod


class Document(ABC):
    def ProcessDocument(self):
        self.Open()
        self.Parse()
        self.Export()
        self.Close()

    @abstractmethod
    def Open(self):
        pass

    @abstractmethod
    def Parse(self):
        pass

    @abstractmethod
    def Export(self):
        pass

    @abstractmethod
    def Close(self):
        pass
