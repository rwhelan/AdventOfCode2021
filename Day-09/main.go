package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"unsafe"
)

type point struct {
	x, y, value int
}

type row map[int]*point
type grid map[int]row

func (g grid) GetPoint(x, y int) *point {
	row, ok := g[y]
	if !ok {
		return nil
	}

	p, ok := row[x]
	if !ok {
		return nil
	}

	return p
}

func (g grid) GetNeighbors(p *point) []*point {
	resp := make([]*point, 0, 4)

	above := g.GetPoint(p.x, p.y-1)
	if above != nil {
		resp = append(resp, above)
	}

	below := g.GetPoint(p.x, p.y+1)
	if below != nil {
		resp = append(resp, below)
	}

	left := g.GetPoint(p.x-1, p.y)
	if left != nil {
		resp = append(resp, left)
	}

	right := g.GetPoint(p.x+1, p.y)
	if right != nil {
		resp = append(resp, right)
	}

	return resp
}

func pointSeen(p *point, points *map[int]*point) bool {
	_, ok := (*points)[int(uintptr(unsafe.Pointer(p)))]
	return ok
}

func walkBasin(g grid, p *point, points *map[int]*point) {
	(*points)[int(uintptr(unsafe.Pointer(p)))] = p
	for _, n := range g.GetNeighbors(p) {
		if n.value != 9 && !pointSeen(n, points) {
			walkBasin(g, n, points)
		}
	}
}

func (g grid) FindBasinSize(p *point) int {
	points := make(map[int]*point)
	walkBasin(g, p, &points)

	return len(points)
}

func (g grid) IsLowPoint(p *point) bool {
	if p == nil {
		panic("Unknown Point")
	}

	if p.value == 9 {
		return false
	}

	for _, npoint := range g.GetNeighbors(p) {
		if npoint.value < p.value {
			return false
		}
	}

	return true
}

func ReadLines(filename string) ([][]byte, error) {
	rawInput, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	rawInput = bytes.TrimRight(rawInput, "\n")
	return bytes.Split(rawInput, []byte("\n")), nil
}

func LoadGrid(inputfile string) (grid, error) {
	inputRows, err := ReadLines(inputfile)
	if err != nil {
		return nil, err
	}

	resp := make(grid)
	for y, r := range inputRows {
		resp[y] = make(row)
		for x, v := range r {
			value, err := strconv.Atoi(string(v))
			if err != nil {
				return nil, err
			}

			resp[y][x] = &point{
				x:     x,
				y:     y,
				value: value,
			}
		}
	}

	return resp, nil
}

func problemOne(g grid) int {
	score := 0
	for _, row := range g {
		for _, p := range row {
			if p == nil {
				continue
			}

			if g.IsLowPoint(p) {
				score += p.value + 1
			}
		}
	}

	return score
}

func problemTwo(g grid) int {
	basinSizes := make([]int, 0)
	for _, row := range g {
		for _, p := range row {
			if p == nil {
				continue
			}

			if g.IsLowPoint(p) {
				basinSizes = append(basinSizes, g.FindBasinSize(p))
			}
		}
	}

	sort.Ints(basinSizes)
	return basinSizes[len(basinSizes)-1] *
		basinSizes[len(basinSizes)-2] *
		basinSizes[len(basinSizes)-3]
}

func main() {
	grid, err := LoadGrid("input")
	if err != nil {
		panic(err)
	}

	fmt.Println("Problem One:", problemOne(grid))
	fmt.Println("Problem Two:", problemTwo(grid))
}
