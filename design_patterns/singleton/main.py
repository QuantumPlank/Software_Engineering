import threading
from typing import Self


class Singleton:
    __instance = None
    __lock = threading.Lock()

    data: str

    def __new__(cls) -> Self:
        if cls.__instance is None:
            with cls.__lock:
                if cls.__instance is None:
                    print("Creating singleton")
                    cls.__instance = super().__new__(cls)
                    cls.__instance.data = "This is a singleton"
        return cls.__instance


def getInstance(id: int):
    s = Singleton()
    print(f"From {id}: {s.data}")


def main():
    threads: list[threading.Thread] = []
    for i in range(10):
        t = threading.Thread(target=getInstance, args=[i])
        threads.append(t)
        t.start()

    for t in threads:
        t.join()


if __name__ == "__main__":
    main()
