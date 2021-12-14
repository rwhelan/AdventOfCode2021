package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
)

func ReadInput(filename string) ([]int, error) {
	rawInput, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if rawInput[len(rawInput)-1] == '\n' {
		rawInput = rawInput[:len(rawInput)-1]
	}

	sections := bytes.Split(rawInput, []byte{','})
	resp := make([]int, len(sections))

	for i, v := range sections {
		num, err := strconv.Atoi(string(v))
		if err != nil {
			return nil, err
		}

		resp[i] = num
	}

	sort.Ints(resp)
	return resp, nil
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

func sum(i []int) int {
	s := 0
	for _, v := range i {
		s += v
	}

	return s
}

func problemOne(input []int) int {
	total := 0
	median := input[len(input)/2]

	for _, v := range input {
		total += abs(median - v)
	}

	return total
}

func problemTwo(input []int) int {
	total := 0
	avg := sum(input) / len(input)

	for _, v := range input {
		dis := abs(avg - v)
		total += (dis * (dis + 1)) / 2
	}

	return total
}

func main() {
	input, err := ReadInput("input")
	if err != nil {
		panic(err)
	}

	fmt.Println("Problem One:", problemOne(input))
	fmt.Println("Problem Two:", problemTwo(input))
}
