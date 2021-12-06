package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

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

	subOne := &SubmarineOne{}
	subTwo := &SubmarineTwo{}

	for _, r := range rows {
		inst := bytes.Split(r, []byte{' '})
		value, err := strconv.Atoi(string(inst[1]))
		if err != nil {
			panic(err)
		}

		switch string(inst[0]) {
		case "forward":
			subOne.MoveForward(value)
			subTwo.MoveForward(value)
		case "down":
			subOne.MoveDown(value)
			subTwo.MoveDown(value)
		case "up":
			subOne.MoveUp(value)
			subTwo.MoveUp(value)
		}
	}

	fmt.Println("Puzzle One:", subOne.X*subOne.Y)
	fmt.Println("Puzzle Two:", subTwo.X*subTwo.Y)
}
