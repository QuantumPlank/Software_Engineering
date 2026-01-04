from abc import ABC, abstractmethod


class Service(ABC):
    @abstractmethod
    def GetUser(self, id: int) -> str:
        pass
