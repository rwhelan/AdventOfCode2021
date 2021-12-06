package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Submarine struct {
	X, Y int
}

func (s *Submarine) MoveUp(steps int) *Submarine {
	s.Y -= steps
	return s
}

func (s *Submarine) MoveDown(steps int) *Submarine {
	s.Y += steps
	return s
}

func (s *Submarine) MoveForward(steps int) *Submarine {
	s.X += steps
	return s
}

func (s *Submarine) MoveBack(steps int) *Submarine {
	s.X -= steps
	return s
}

func ReadLines(filename string) ([][]byte, error) {
	rawInput, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	rawInput = bytes.TrimRight(rawInput, "\n")
	return bytes.Split(rawInput, []byte("\n")), nil
}

func main() {
	rows, err := ReadLines("input")
	if err != nil {
		panic(err)
	}

	sub := &Submarine{}

	for _, r := range rows {
		inst := bytes.Split(r, []byte{' '})
		value, err := strconv.Atoi(string(inst[1]))
		if err != nil {
			panic(err)
		}

		switch string(inst[0]) {
		case "forward":
			sub.MoveForward(value)
		case "down":
			sub.MoveDown(value)
		case "up":
			sub.MoveUp(value)
		}
	}

	fmt.Println("Puzzle One:", sub.X*sub.Y)
}
