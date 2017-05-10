package player

import (
	"fmt"
	"github.com/frieser/camelrace/camel"
)

type Human Player

//Human constructor
func NewHuman(name string, transport camel.Movable) *Human {
	return &Human{
		Name:      name,
		Transport: transport,
		Distance:  0}
}

//Move the human in the game
func (human *Human) Play() float32 {
	movement := human.Transport.Move(human.Roll())
	human.Distance += movement

	return human.Distance
}

//Returns the distance traveled in a roll. In this case ask the human input
func (human Human) Roll() float32 {
	var i float32

	fmt.Print("\nType the movement for the human player: ")

	_, err := fmt.Scanf("%f", &i)

	if err != nil {
		return 0
	}

	return i
}

//Type format output
func (human Human) String() string {
	return fmt.Sprintf("Human: %s with a %s", human.Name, human.Transport)
}
