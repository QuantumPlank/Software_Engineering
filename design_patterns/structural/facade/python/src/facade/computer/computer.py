class RAM:
    def load(self) -> None:
        print("RAM: Loading...")

    def flush(self) -> None:
        print("RAM: Flushing...")


class CPU:
    def execute(self) -> None:
        print("CPU: Executing...")

    def freeze(self) -> None:
        print("CPU: Freezing...")


class Computer:
    def __init__(self) -> None:
        self._cpu = CPU()
        self._ram = RAM()

    def start(self) -> None:
        print("Computer: Starting!")
        self._ram.load()
        self._cpu.execute()
        print("Computer: Started!")

    def stop(self) -> None:
        print("Computer: Stopping!")
        self._cpu.freeze()
        self._ram.flush()
        print("Computer: Stopped!")
