package scene

import (
	"fmt"
	"testing"

	"github.com/matryer/is"
)

func TestLoad(t *testing.T) {
	is := is.New(t)
	s, err := New("../test/scenes/1.json")
	is.NoErr(err)
	fmt.Println(s)
}
