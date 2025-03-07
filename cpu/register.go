package cpu

import (
	"fmt"
)

type Register struct {
	value uint16
}

func (r *Register) ZeroFlag() bool {
	// zero flag stored in bit 7
	return r.value>>7 == 1
}

func (r *Register) SubtractionFlag() bool {
	// subtraction flag stored in bit 6
	return r.value>>6 == 1
}

func (r *Register) HalfCarryFlag() bool {
	// half carry flag stored in bit 5
	return r.value>>5 == 1
}

func (r *Register) CarryFlag() bool {
	// carry flag stored in bit 4
	return r.value>>4 == 1
}

func (r *Register) Value() uint16 {
	return r.value
}

func (r *Register) SetValue(value uint16) {
	r.value = value
}

func (r *Register) Increment() {
	r.value++
}

func (r *Register) Decrement() {
	r.value--
}

func (r *Register) Add(value uint16) {
	r.value += value
}

func (r *Register) Subtract(value uint16) {
	r.value -= value
}

func (r *Register) Print() {
	fmt.Printf("%04x", r.value)
}
