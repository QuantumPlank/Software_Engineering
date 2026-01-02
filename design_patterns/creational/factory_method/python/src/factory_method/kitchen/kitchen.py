from abc import ABC, abstractmethod


class Meal(ABC):
    @abstractmethod
    def serve(self) -> str:
        pass

    def dummy(self) -> None:
        pass


class Pizza(Meal):
    def serve(self) -> str:
        return "Serving pizza"


class Hamburger(Meal):
    def serve(self) -> str:
        return "Serving hamburger"


class Kitchen:
    def create_meal(self, meal_type) -> Meal:
        if meal_type == "pizza":
            return Pizza()
        elif meal_type == "hamburger":
            return Hamburger()
        else:
            raise ValueError("Unknown meal type")
