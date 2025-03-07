package cpu

import (
	"fmt"
)

type CPU struct {
	Registers map[string]*Register
}

func (c *CPU) Print() {
	for key, value := range c.Registers {
		fmt.Printf("%s: ", key)
		value.Print()
		fmt.Println()
	}
}

func NewCPU() *CPU {
	AF := Register{}
	BC := Register{}
	DE := Register{}
	HL := Register{}
	StackPointer := Register{}
	ProgramCounter := Register{}

	return &CPU{
		Registers: map[string]*Register{"AF": &AF, "BC": &BC, "DE": &DE, "HL": &HL, "SP": &StackPointer, "PC": &ProgramCounter},
	}
}
