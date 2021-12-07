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

	for i, b := range data {
		value, err := strconv.ParseUint(string(b), 2, 12)
		if err != nil {
			return nil, err
		}

		resp[i] = int(value)
	}

	return resp, nil
}

func problemOne(ints []int) int {
	gamma := 0

	for i := 0; i < 12; i++ {
		ones := 0
		for _, row := range ints {
			if (row>>i)&1 == 1 {
				ones++
			}
		}

		if ones > len(ints)/2 {
			gamma |= (1 << i)
		}
	}

	return gamma * (gamma ^ 4095)
}

func main() {
	rows, err := ReadLines("input")
	if err != nil {
		panic(err)
	}

	ints, err := ReadInts(rows)
	if err != nil {
		panic(err)
	}

	fmt.Println("Problem One:", problemOne(ints))
}
