import time

from .service import Service


class RealService(Service):
    def GetUser(self, id: int) -> str:
        time.sleep(2)
        return f"User-{id}"
