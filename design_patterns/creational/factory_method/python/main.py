from abc import ABC, abstractmethod


class Meal(ABC):
    @abstractmethod
    def serve(self) -> str:
        pass


class Pizza(Meal):
    def serve(self) -> str:
        return "Serving pizza"


class Hamburger(Meal):
    def serve(self) -> str:
        return "Serving hamburger"


class Kitchen:
    def create_meal(self, meal_type):
        if meal_type == "pizza":
            return Pizza()
        elif meal_type == "hamburger":
            return Hamburger()
        else:
            raise ValueError("Unknown meal type")


if __name__ == "__main__":
    k = Kitchen()
    meal1 = k.create_meal("pizza")
    print(meal1.serve())
    meal2 = k.create_meal("hamburger")
    print(meal2.serve())
