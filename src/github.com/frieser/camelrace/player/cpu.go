package player

import (
	"fmt"
	"github.com/frieser/camelrace/camel"
	"math/rand"
	"time"
)

type Cpu Player

//Cpu constructor
func NewCpu(name string, transport camel.Movable) *Cpu {
	return &Cpu{
		Name:      name,
		Transport: transport,
		Distance:  0}
}

//Move the cpu in the game
func (cpu *Cpu) Play() float32 {
	movement := cpu.Transport.Move(cpu.Roll())
	cpu.Distance += movement

	return cpu.Distance
}

//Returns the distance traveled in a roll. In this case is a range random number
func (cpu Cpu) Roll() float32 {
	rand.Seed(time.Now().UTC().UnixNano())

	min := 0
	max := 3

	return (float32)(min + rand.Intn(max-min))
}

//Type format output
func (cpu Cpu) String() string {
	return fmt.Sprintf("Cpu: %s with a %s", cpu.Name, cpu.Transport)
}
