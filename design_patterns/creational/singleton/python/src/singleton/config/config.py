import json
import threading
from typing import Self


class Config:
    _instance = None
    _lock = threading.Lock()

    def __new__(cls) -> Self:
        if cls._instance is None:
            with cls._lock:
                if cls._instance is None:
                    instance = super().__new__(cls)
                    instance._load()
                    cls._instance = instance
        return cls._instance

    def _load(self) -> None:
        with open("config.json") as f:
            data: dict = json.load(f)
            self._AppName = data.get("AppName", "")
            self._Version = data.get("Version", "")

    @property
    def AppName(self) -> str:
        return self._AppName

    @property
    def Version(self) -> str:
        return self._Version
