package main

import (
	"gophergba/cpu"
)

func main() {
	c := cpu.NewCPU()
	c.Print()

	// test out the registers
	c.Registers["AF"].Increment()
	c.Registers["BC"].Decrement()
	c.Registers["DE"].Add(0x10)
	c.Registers["HL"].Subtract(0x20)
	c.Registers["SP"].SetValue(0x1234)
	c.Registers["PC"].SetValue(0x5678)

	c.Print()
}
