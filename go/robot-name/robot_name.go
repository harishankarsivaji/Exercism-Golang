package robotname

import (
	"errors"
	"math/rand"
)

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers = "0123456789"
const maxNames = 26 * 26 * 10 * 10 * 10

var takenNames = map[string]bool{}

//Robot name
type Robot struct {
	name string
}

// Creates a random name with letters and numbers
func createName() string {
	n := make([]byte, 5)
	for i := range n {
		if i < 2 {
			n[i] = letters[rand.Intn(len(letters))]
		} else {
			n[i] = numbers[rand.Intn(len(numbers))]
		}
	}
	return string(n)
}

//Name returns new name
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		if len(takenNames) == maxNames {
			return r.name, errors.New("no names available to assign")
		}
		r.name = createName()
		for takenNames[r.name] {
			r.name = createName()
		}
		takenNames[r.name] = true
	}

	return r.name, nil
}

//Reset the name
func (r *Robot) Reset() {
	r.name = ""
}
