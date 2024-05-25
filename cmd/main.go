package main

import (
	"fmt"
	"nes-emulator/pkg/cpu"
)

func main() {
	bus := cpu.NewDefaultBus()
	cpu := cpu.NewCPU6502()
	cpu.ConnectBus(bus)
	cpu.Write(0x0000, 0xF5)
	fmt.Printf("Read from 0x0000: %d\n", cpu.Read(0x0000))
}
