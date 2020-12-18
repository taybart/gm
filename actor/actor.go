package actor

import "github.com/taybart/gm/bag"

type Actor struct {
	S bag.Stats `json:"stats,omitempty"`
}

func New(s bag.Stats) Actor {
	return Actor{S: s}
}

func AssignAction(playerID, action string) error {
	return nil
}
