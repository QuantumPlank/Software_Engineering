from abc import ABC, abstractmethod

from charger import MicroUSBCharger


class Charger(ABC):
    @abstractmethod
    def charge(self) -> str:
        pass


class MicroUSBToUSBAdapter(Charger):
    def __init__(self, MicroUSB: MicroUSBCharger) -> None:
        super().__init__()
        self._MicroUSBCharger = MicroUSB

    def charge(self) -> str:
        return self._MicroUSBCharger.ChargeMicroUSB()


def main():
    microusb = MicroUSBCharger()
    usb = MicroUSBToUSBAdapter(MicroUSB=microusb)
    print(usb.charge())


if __name__ == "__main__":
    main()
