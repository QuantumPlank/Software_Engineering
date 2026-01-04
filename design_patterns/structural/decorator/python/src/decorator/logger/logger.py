from abc import ABC, abstractmethod
from datetime import datetime


class Logger(ABC):
    @abstractmethod
    def log(self, message: str) -> None:
        pass


class ConsoleLogger(Logger):
    def log(self, message: str) -> None:
        print(message)


class LoggerDecorator(Logger):
    def __init__(self, logger: Logger) -> None:
        self._logger = logger

    def log(self, message: str) -> None:
        self._logger.log(message)


class TimestampDecorator(LoggerDecorator):
    def log(self, message: str) -> None:
        return super().log(
            f"[{datetime.now().strftime('%Y-%m-%d %H:%M:%S')}] {message}"
        )


class LevelDecorator(LoggerDecorator):
    def __init__(self, logger: Logger, level: str) -> None:
        super().__init__(logger)
        self._level = level

    def log(self, message: str) -> None:
        return super().log(f"[{self._level}] {message}")


class FilenameDecorator(LoggerDecorator):
    def __init__(self, logger: Logger, filename: str) -> None:
        super().__init__(logger)
        self._filename = filename

    def log(self, message: str) -> None:
        return super().log(f"[{self._filename}] {message}")
