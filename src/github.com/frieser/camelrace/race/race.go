//Package to manage all aspects of the camel race
package race

import (
	"fmt"
	"github.com/frieser/camelrace/event"
	"github.com/frieser/camelrace/round"
)

type Race struct {
	Rounds  map[int]round.Round
	Players map[int]round.Actor
	Events  event.EventDispatcher
	Config  Config
}

//Start the race rounds and proclaim the winner
func (camelRace Race) Start() {
	for {
		winner := camelRace.playRound()
		victory := camelRace.checkWinner(winner)

		if victory {
			return
		}
	}
}

func (camelRace Race) checkWinner(winner round.Actor) bool {
	if winner != nil {
		return camelRace.proclaimVictory(winner)
	}

	return false
}

func (camelRace Race) playRound() round.Actor {
	for _, player := range camelRace.Players {
		distance := player.Play()

		fmt.Printf("%v has traveled %.2f\n", player, distance)

		if distance >= camelRace.Config.Length {
			return player
		}
	}

	return nil
}

func (camelRace Race) proclaimVictory(actor round.Actor) bool {
	fmt.Printf("\nWinner %v", actor)

	return true
}
