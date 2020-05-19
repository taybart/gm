package main

import (
	"flag"

	"github.com/taybart/gm/dice"
	"github.com/taybart/log"
)

type config struct {
	reset  bool
	roll   string
	player string
}

var c config

func init() {
	flag.StringVar(&c.roll, "r", "", "Roll a die [d4, d6, d12, d18]")
	flag.StringVar(&c.player, "p", "", "Do player stuff")
}

func main() {
	flag.Parse()
	switch {
	// Dice shit
	case c.roll != "":
		res, err := dice.Roll(c.roll)
		if err != nil {
			log.Error(err)
		} else {
			log.Info("You rolled", res)
		}
	case c.player != "":
	}
}
