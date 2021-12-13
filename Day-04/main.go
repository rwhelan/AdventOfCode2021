package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Space struct {
	X, Y, Value int
	Marked      bool
}

type Board struct {
	spaces  [5][5]*Space
	idx     map[int]*Space
	histroy []int
}

func (b *Board) Print() string {
	resp := strings.Builder{}
	for _, row := range b.spaces {
		resp.WriteString(
			fmt.Sprintf(
				"%2d %2d %2d %2d %2d\n",
				row[0].Value, row[1].Value, row[2].Value, row[3].Value, row[4].Value,
			),
		)
	}

	return resp.String()
}

func (b *Board) EvalDraw(value int) *Board {
	if b.IsWinner() {
		return b
	}

	space, ok := b.idx[value]
	if ok && !space.Marked {
		space.Marked = true
		b.histroy = append(b.histroy, value)
	}

	return b
}

func (b *Board) UnmarkedNumbers() []int {
	resp := make([]int, 0)
	for _, space := range b.idx {
		if !space.Marked {
			resp = append(resp, space.Value)
		}
	}

	return resp
}

func (b *Board) IsWinner() bool {
	all := func(spaces ...*Space) bool {
		for _, s := range spaces {
			if !s.Marked {
				return false
			}
		}

		return true
	}

	for i := 0; i < 5; i++ {
		if all(
			b.spaces[i][0],
			b.spaces[i][1],
			b.spaces[i][2],
			b.spaces[i][3],
			b.spaces[i][4],
		) {
			return true
		}

		if all(
			b.spaces[0][i],
			b.spaces[1][i],
			b.spaces[2][i],
			b.spaces[3][i],
			b.spaces[4][i],
		) {
			return true
		}
	}

	return false
}

func IntSum(values []int) int {
	resp := 0
	for _, v := range values {
		resp += v
	}

	return resp
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

	boardScore := func(b *Board) int {
		return b.histroy[len(b.histroy)-1] * IntSum(b.UnmarkedNumbers())
	}

	boardsInWinningOrder := make([]*Board, 0)

	boardAlreadyWon := func(b *Board) bool {
		for _, winner := range boardsInWinningOrder {
			if b == winner {
				return true
			}
		}

		return false
	}

	for _, d := range draws {
		for _, b := range boards {
			b.EvalDraw(d)
			if b.IsWinner() && !boardAlreadyWon(b) {
				boardsInWinningOrder = append(boardsInWinningOrder, b)
			}
		}
	}

	fmt.Println("Problem One:", boardScore(boardsInWinningOrder[0]))
	fmt.Println("Problem Two:", boardScore(boardsInWinningOrder[len(boardsInWinningOrder)-1]))
}
