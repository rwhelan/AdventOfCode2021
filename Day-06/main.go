package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

func ReadInput(filename string) (map[int]int, error) {
	rawInput, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if rawInput[len(rawInput)-1] == '\n' {
		rawInput = rawInput[:len(rawInput)-1]
	}

	sections := bytes.Split(rawInput, []byte{','})
	resp := make(map[int]int)

	for _, v := range sections {
		num, err := strconv.Atoi(string(v))
		if err != nil {
			return nil, err
		}

		resp[num]++
	}

	return resp, nil
}

func solve(initial map[int]int, gen int) int {
	fish := make(map[int]int)
	for k, v := range initial {
		fish[k] = v
	}

	for i := 0; i < gen; i++ {
		nfish := make(map[int]int)
		for k, v := range fish {
			if k == 0 {
				nfish[8] += fish[0]
				nfish[6] += fish[0]
			} else {
				nfish[k-1] += v
			}
		}

		fish = nfish
	}

	sum := 0
	for _, v := range fish {
		sum += v
	}

	return sum
}

func main() {
	// fish := []int{3, 4, 3, 1, 2}
	fish, err := ReadInput("input")
	if err != nil {
		panic(err)
	}

	fmt.Println("Problem One:", solve(fish, 80))
	fmt.Println("Problem Two:", solve(fish, 256))
}
