package race

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Length              float32 `yaml:"distance"`
	Lanes               int     `yaml:"players"`
	Terrain             string  `yaml:"terrain"`
	RoundTime           int     `yaml:"round_time"`
	RoundDividerTime    int     `yaml:"round_divider"`
	MaxMovementPerRound int     `yaml:"movement_per_round"`
}

//Load the race config from a yaml file
func LoadFromFile(path string) Config {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Printf("Error #%v ", err)
	}
	config := Config{}

	err = yaml.Unmarshal(file, &config)

	if err != nil {
		fmt.Errorf("Unmarshal: %v", err)
	}

	return config
}
