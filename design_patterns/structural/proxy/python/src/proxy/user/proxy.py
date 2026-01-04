import threading

from .service import Service


class ProxyService(Service):
    def __init__(self, realService: Service) -> None:
        self._realservice = realService
        self._cache = {}
        self._lock = threading.Lock()

    def GetUser(self, id: int) -> str:
        with self._lock:
            if id in self._cache:
                print(f"Cache hit for the user {id}")
                return self._cache[id]

        print(f"Cache miss, calling real service for user {id}")
        user = self._realservice.GetUser(id)

        with self._lock:
            self._cache[id] = user

        return user
