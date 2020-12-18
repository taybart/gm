package main

import (
	"flag"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/taybart/gm/bag"
	"github.com/taybart/log"
	"golang.org/x/crypto/ssh/terminal"
)


type config struct {
	reset  bool
	roll   string
}

var c config

func init() {
	flag.StringVar(&c.roll, "r", "", "Roll a die [d4, d6, d12, d18, d20]")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	flag.Parse()
	switch {
	// Dice shit
	case c.roll != "":
		res, err := bag.Roll(c.roll)
		if err != nil {
			log.Error(err)
		} else {
			log.Info("You rolled", res)
		}
	default:
		dorepl()
	}
}

func dorepl() {
	fd := int(os.Stdin.Fd())
	termState, err := terminal.MakeRaw(fd)
	if err != nil {
		log.Error(err)
		return
	}
	defer terminal.Restore(fd, termState)

	n := terminal.NewTerminal(os.Stdin, "-> ")
	n.SetSize(int(^uint(0)>>1), 0)

	for {
		ln, err := n.ReadLine()
		if err != nil {
			log.Error(err)
			break
		}

		terminal.Restore(int(os.Stdin.Fd()), termState)
		if err != nil {
			log.Error(err)
			break
		}

		quit := parseCmd(ln)
		if quit {
			break
		}

		termState, err = terminal.MakeRaw(int(os.Stdin.Fd()))
		if err != nil {
			log.Error(err)
			break
		}
	}
}

func parseCmd(cmd string) bool {
	args := strings.Split(cmd, " ")

	cmd = args[0]
	args = args[1:]

	switch cmd {
	case "r", "roll":
    d := bag.DefaultDie
    if len(args) > 0 {
      d = args[0]
    }
		res, err := bag.Roll(d)
		if err != nil {
			log.Error(err)
		} else {
			log.Info("You rolled", res)
		}
	case "q", "quit", "exit":
		return true
	}
	return false
}
