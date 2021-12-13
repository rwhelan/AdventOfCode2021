package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

type point struct {
	x, y int
}

func ReadLines(filename string) ([][]byte, error) {
	rawInput, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	rawInput = bytes.TrimRight(rawInput, "\n")
	return bytes.Split(rawInput, []byte("\n")), nil
}

func ParseEndpoints(row []byte) (point, point) {
	s := func(b []byte) (int, int) {
		t := bytes.Split(b, []byte{','})

		x, err := strconv.Atoi(string(t[0]))
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(string(t[1]))
		if err != nil {
			panic(err)
		}

		return x, y
	}

	p := bytes.Split(row, []byte(" -> "))
	left, right := p[0], p[1]

	lx, ly := s(left)
	rx, ry := s(right)

	return point{
			x: lx,
			y: ly,
		}, point{
			x: rx,
			y: ry,
		}

}

// I Hate this function.  If it was 'real' code,
// it would need a refactor
func ExpandLine(row []byte) []point {
	resp := make([]point, 0)
	start, end := ParseEndpoints(row)

	if start.x == end.x {
		if start.y < end.y {
			for i := start.y; i <= end.y; i++ {
				resp = append(resp, point{x: start.x, y: i})
			}

			return resp
		}

		for i := start.y; i >= end.y; i-- {
			resp = append(resp, point{x: start.x, y: i})
		}

		return resp
	}

	if start.y == end.y {
		if start.x < end.x {
			for i := start.x; i <= end.x; i++ {
				resp = append(resp, point{x: i, y: start.y})
			}

			return resp
		}

		for i := start.x; i >= end.x; i-- {
			resp = append(resp, point{x: i, y: start.y})
		}

		return resp
	}

	if start.x > end.x {
		start, end = end, start
	}

	if start.y < end.y {
		for i := 0; i+start.x <= end.x; i++ {
			resp = append(resp, point{x: start.x + i, y: start.y + i})
		}

		return resp
	}

	for i := 0; i+start.x <= end.x; i++ {
		resp = append(resp, point{x: start.x + i, y: start.y - i})
	}

	return resp
}

func IsHorizontal(points []point) bool {
	return points[0].x == points[len(points)-1].x
}

func IsVertical(points []point) bool {
	return points[0].y == points[len(points)-1].y
}

func IsDiagonal(points []point) bool {
	return !IsHorizontal(points) && !IsVertical(points)
}

func CountBoard(board [1000][1000]int) int {
	resp := 0
	for _, row := range board {
		for _, cell := range row {
			if cell >= 2 {
				resp++
			}
		}
	}

	return resp
}

func main() {
	boardOne := [1000][1000]int{}
	boardTwo := [1000][1000]int{}

	rows, err := ReadLines("input")
	if err != nil {
		panic(err)
	}

	for _, r := range rows {
		line := ExpandLine(r)
		if len(line) == 0 {
			panic("Ooops")
		}
		for _, p := range line {
			boardTwo[p.x][p.y]++
		}

		if !IsDiagonal(line) {
			for _, p := range line {
				boardOne[p.x][p.y]++
			}
		}
	}

	fmt.Println("Problem One:", CountBoard(boardOne))
	fmt.Println("Problem Two:", CountBoard(boardTwo))

}
