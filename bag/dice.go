package bag

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
)

const (
  DefaultDie = "d20"
)

func Roll(die string) (res int, err error) {
	if die == "" {
		err = fmt.Errorf("no roll provided")
		return
	}
  size, err := parseDie(die)
  res = roll(size)
	return
}

func RollDefault() (res int) {
  size, _ := parseDie(DefaultDie) // should always work via tests
	return roll(size)
}

func parseDie(die string) (size int, err error){
	r := regexp.MustCompile("d?(4|6|12|20)")
	number := r.FindStringSubmatch(die)
	if len(number) == 0 {
		err = fmt.Errorf("unknown die size of %v", die)
		return
	}
	size, err = strconv.Atoi(number[1])
	if err != nil {
		err = fmt.Errorf("could not convert %v to number", number[1])
		return
	}
  return
}

func roll(max int) int {
	n := rand.Intn(max)
	return n + 1
}
