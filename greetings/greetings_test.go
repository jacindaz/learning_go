package greetings

import (
	"testing"
)

func TestHello(t *testing.T) {
    total := Hello("Jason")
    if total != "Hi, Jason. Welcome!" {
       t.Errorf("error!")
    }
}
