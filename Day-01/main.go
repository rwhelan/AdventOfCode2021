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

func ReadInts(data [][]byte) ([]int, error) {
	resp := make([]int, len(data))
	var err error
	for i, b := range data {
		if resp[i], err = strconv.Atoi(string(b)); err != nil {
			return nil, err
		}
	}

	return resp, err
}

func problemOne(ints []int) int {
	increases := 0

	for i := 1; i < len(ints); i++ {
		if ints[i] > ints[i-1] {
			increases++
		}
	}

	return increases
}

func problemTwo(ints []int) int {
	subsum := func(i int) int {
		return ints[i] + ints[i+1] + ints[i+2]
	}

	currentValue := subsum(0)
	increases := 0

	for i := 1; i < len(ints)-2; i++ {
		if subsum(i) > currentValue {
			increases++
		}

		currentValue = subsum(i)
	}

	return increases
}

func main() {
	data, err := ReadLines("input")
	if err != nil {
		panic(err)
	}

	ints, err := ReadInts(data)
	if err != nil {
		panic(err)
	}

	fmt.Println("Problem One:", problemOne(ints))
	fmt.Println("Problem Two:", problemTwo(ints))
}
