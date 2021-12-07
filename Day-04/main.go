package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Space struct {
	X, Y, Value int
}

type Board struct {
	spaces [5][5]*Space
	idx    map[int]*Space
}

func ParseBoard(input []byte) *Board {
	resp := &Board{
		idx: make(map[int]*Space),
	}

	for x, row := range bytes.Split(input, []byte{'\n'}) {
		// remove double spaces
		row = bytes.ReplaceAll(row, []byte{' ', ' '}, []byte{' '})

		// remove any leading spaces
		if row[0] == ' ' {
			row = row[1:]
		}

		for y, space := range bytes.Split(row, []byte{' '}) {
			v, err := strconv.Atoi(string(space))
			if err != nil {
				panic(err)
			}

			space := &Space{
				X:     x,
				Y:     y,
				Value: v,
			}

			resp.idx[v] = space
			resp.spaces[x][y] = space
		}
	}

	return resp
}

func ParseInput(filename string) ([]int, []*Board, error) {
	rawInput, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, nil, err
	}

	sections := bytes.Split(rawInput, []byte{'\n', '\n'})

	var draws []int
	boards := make([]*Board, 0)
	for i, s := range sections {
		if i >= 1 {
			boards = append(boards, ParseBoard(s))
		}

		if i == 0 {
			ds := bytes.Split(s, []byte(","))
			draws = make([]int, len(ds))

			for j, dss := range ds {
				draws[j], err = strconv.Atoi(string(dss))
				if err != nil {
					return nil, nil, err
				}
			}
		}
	}

	return draws, boards, nil
}

func main() {
	draws, boards, err := ParseInput("input")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", boards)
	fmt.Println(draws)
}
