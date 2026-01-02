from typing import List

from .component import Component


class Folder(Component):
    def __init__(self, name: str) -> None:
        self._name = name
        self._components: List[Component] = []

    def Add(self, component: Component) -> None:
        self._components.append(component)

    def Size(self) -> int:
        total = 0
        for component in self._components:
            total += component.Size()
        return total
