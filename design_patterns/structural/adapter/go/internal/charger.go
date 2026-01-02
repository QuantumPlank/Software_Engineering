package charger

type MicroUSBCharger struct{}

func (m *MicroUSBCharger) ChargeMicroUSB() string {
	return "Charging via Micro-USB"
}
