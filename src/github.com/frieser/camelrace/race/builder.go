package race

import (
	"github.com/frieser/camelrace/camel"
	"github.com/frieser/camelrace/event"
	"github.com/frieser/camelrace/lane"
	"github.com/frieser/camelrace/player"
	"github.com/frieser/camelrace/round"
	"github.com/frieser/camelrace/terrain"
	"strconv"
)

//Instantiate the race object, config it and register players and event listeners
func Build(config Config) Race {
	camelRace := new(Race)

	camelRace.Config = config

	camelRace.registerPlayers()
	camelRace.registerRaceEventsListeners()

	return *camelRace
}

func (camelRace *Race) registerRaceEventsListeners() {
	camelRace.Events = event.NewEventsMap()

	camelRace.Events.On("weapon", func(args ...interface{}) {
		//do actions to affect other players with the arguments passed
	})

	camelRace.Events.On("specialKeyStroke", func(args ...interface{}) {
		//do actions to affect other players with the arguments passed
	})
}

func (camelRace *Race) registerPlayers() {
	camelRace.Players = make(map[int]round.Actor, camelRace.Config.Lanes)

	for i := camelRace.Config.Lanes; i > 0; i-- {
		name := "Player #" + strconv.Itoa(i)
		transport := camel.NewBactrianCamel(lane.Lane{Terrain: terrain.Terrain{Material: camelRace.Config.Terrain}}, camelRace.Events)

		camelRace.Players[i] = player.NewCpu(name, transport)
	}

	transport := camel.NewDomesticCamel
	camelRace.Players[0] = player.NewHuman("Human Player", transport(lane.Lane{Terrain: terrain.Terrain{Material: camelRace.Config.Terrain}}, camelRace.Events))
}
