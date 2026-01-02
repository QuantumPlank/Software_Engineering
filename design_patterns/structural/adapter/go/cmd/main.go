package main

import (
	"fmt"

	charger "adapter.example/internal"
)

type USBCharger interface {
	Charge() string
}

type MicroUSBToUsbAdapter struct {
	microUSB *charger.MicroUSBCharger
}

func (a *MicroUSBToUsbAdapter) Charge() string {
	return a.microUSB.ChargeMicroUSB()
}

func main() {
	microUSBCharger := &charger.MicroUSBCharger{}
	adapter := &MicroUSBToUsbAdapter{
		microUSB: microUSBCharger,
	}
	fmt.Println(adapter.Charge())
}
