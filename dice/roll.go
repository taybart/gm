package dice

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

func Roll(roll string) (res int, err error) {
	if roll == "" {
		err = fmt.Errorf("no roll provided")
		return
	}
	r := regexp.MustCompile("d?(4|6|12|20)")
	number := r.FindStringSubmatch(roll)
	if len(number) == 0 {
		err = fmt.Errorf("unknown die size of %v", roll)
		return
	}
	s, err := strconv.Atoi(number[1])
	if err != nil {
		err = fmt.Errorf("could not convert %v to number", number[1])
		return
	}
	rand.Seed(time.Now().UnixNano())
	max := s
	n := rand.Intn(max)
	res = n + 1
	return
}
