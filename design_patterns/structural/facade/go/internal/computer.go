package computer

import "fmt"

type cpu struct{}

func (c *cpu) Execute() {
	fmt.Println("CPU: Executing instruction...")
}

func (c *cpu) Freeze() {
	fmt.Println("CPU: Freezing...")
}

type ram struct{}

func (m *ram) Load() {
	fmt.Println("RAM: Loading...")
}

func (m *ram) Flush() {
	fmt.Println("RAM: Flushing...")
}

type Computer struct {
	CPU *cpu
	RAM *ram
}

func (computer *Computer) Start() {
	fmt.Println("Starting computer!")
	computer.RAM.Load()
	computer.CPU.Execute()
	fmt.Println("Started!")
}

func (computer *Computer) Stop() {
	fmt.Println("Stoping computer!")
	computer.RAM.Flush()
	computer.CPU.Freeze()
	fmt.Println("Stopped!")
}

func NewComputer() *Computer {
	return &Computer{
		CPU: &cpu{},
		RAM: &ram{},
	}
}
