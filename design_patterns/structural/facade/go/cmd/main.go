package main

import computer "facade.example/internal"

func main() {
	computer := computer.NewComputer()
	computer.Start()
	computer.Stop()
}
