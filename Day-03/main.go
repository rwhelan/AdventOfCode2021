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

func problemTwo(ints []int) int {
	mostCommon := func(ints []int, pos int) int {
		ones, zeros := 0, 0
		for _, row := range ints {
			if (row>>pos)&1 == 1 {
				ones++
			} else {
				zeros++
			}
		}

		if ones >= zeros {
			return 1
		}

		return 0
	}

	leastCommon := func(ints []int, pos int) int {
		mc := mostCommon(ints, pos)

		if mc == 1 {
			return 0
		}

		return 1
	}

	filter := func(ints []int, pos int, val int) []int {
		resp := make([]int, 0)

		for _, row := range ints {
			if (row>>pos)&1 == val {
				resp = append(resp, row)
			}
		}

		return resp
	}

	oxygenGenRating := func() int {
		list := filter(ints, 11, mostCommon(ints, 11))
		for i := 10; i >= 0; i-- {
			mc := mostCommon(list, i)

			list = filter(list, i, mc)
			if len(list) == 1 {
				return list[0]
			}
		}

		panic("End Of List")
	}

	c02scrubber := func() int {
		list := filter(ints, 11, leastCommon(ints, 11))
		for i := 10; i >= 0; i-- {
			lc := leastCommon(list, i)

			list = filter(list, i, lc)
			if len(list) == 1 {
				return list[0]
			}
		}

		panic("End Of List")
	}

	return oxygenGenRating() * c02scrubber()
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
	fmt.Println("Problem Two:", problemTwo(ints))
}
