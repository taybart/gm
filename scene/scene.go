package scene

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/taybart/gm/actor"
	"github.com/taybart/gm/bag"
)

type Action int

type Scene struct {
	Name       string   `json:"name,omitempty"`
	ActorNames []string `json:"actors,omitempty"`

	players map[string]actor.Actor
	enemies map[string]actor.Actor
	actions []Action
}

func New(filename string) (s Scene, err error) {
	s, err = loadScene(filename)
	if err != nil {
		return
	}
	base := "../test/characters"
	// Scene{
	// 	players: make(map[string]player.Player),
	// }
	for _, a := range s.ActorNames {
		var stats bag.Stats
		stats, err = bag.BuildStats(fmt.Sprintf("%s/%s.json", base, a))
		if err != nil {
			return
		}
		p := actor.New(stats)
		s.players[a] = p

	}
	return
}

func loadScene(filename string) (s Scene, err error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	s.players = make(map[string]actor.Actor)
	s.enemies = make(map[string]actor.Actor)
	err = json.Unmarshal(f, &s)
	return
}

func (s *Scene) Action(a Action) {}
