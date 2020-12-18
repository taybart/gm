package bag

import (
	"encoding/json"
	"io/ioutil"
)

type Stats struct {
	HP         int `json:"hp,omitempty"`
	MaxHP      int `json:"max_hp,omitempty"`
	Toughness  int `json:"toughness,omitempty"`
	Dexterity  int `json:"dexterity,omitempty"`
	Strength   int `json:"strength,omitempty"`
	Perception int `json:"perception,omitempty"`
	Mind       int `json:"mind,omitempty"`
	Charisma   int `json:"charisma,omitempty"`
}

func BuildStats(filename string) (s Stats, err error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(f, &s)
	return
}
