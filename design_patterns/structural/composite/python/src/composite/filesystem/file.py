from .component import Component


class File(Component):
    def __init__(self, name: str, size: int) -> None:
        self._name = name
        self._size = size

    def Size(self) -> int:
        return self._size
