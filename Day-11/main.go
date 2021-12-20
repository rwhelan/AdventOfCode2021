package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type point struct {
	x, y, value int
	flashed     bool
}

type row map[int]*point
type grid map[int]row

func (g grid) Print() string {
	resp := strings.Builder{}
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			v := g.GetPoint(x, y).value
			if v == 0 {
				resp.WriteString(
					fmt.Sprintf("\033[1m%d\033[0m", v),
				)
			} else {
				resp.WriteString(
					fmt.Sprintf("%d", v),
				)
			}
		}
		resp.WriteString("\n")
	}

	return resp.String()
}

func flash(g grid, p *point) {
	p.flashed = true

	for _, n := range g.GetNeighbors(p) {
		n.value++
		if !n.flashed && n.value > 9 {
			flash(g, n)
		}
	}
}

func (g grid) Step() int {
	allPoints := g.GetAllPoints()

	for _, p := range allPoints {
		p.value++
		p.flashed = false
	}

	for _, p := range allPoints {
		if p.value > 9 && !p.flashed {
			flash(g, p)
		}
	}

	flashed := 0
	for _, p := range allPoints {
		if p.value > 9 {
			flashed++
			p.value = 0
		}
	}

	return flashed

}

func (g grid) GetAllPoints() []*point {
	resp := make([]*point, 0, 100)
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			resp = append(resp, g.GetPoint(x, y))
		}
	}

	return resp
}

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
	resp := make([]*point, 0, 8)

	top := g.GetPoint(p.x, p.y-1)
	if top != nil {
		resp = append(resp, top)
	}

	topRight := g.GetPoint(p.x+1, p.y-1)
	if topRight != nil {
		resp = append(resp, topRight)
	}

	right := g.GetPoint(p.x+1, p.y)
	if right != nil {
		resp = append(resp, right)
	}

	bottomRight := g.GetPoint(p.x+1, p.y+1)
	if bottomRight != nil {
		resp = append(resp, bottomRight)
	}

	bottom := g.GetPoint(p.x, p.y+1)
	if bottom != nil {
		resp = append(resp, bottom)
	}

	bottomLeft := g.GetPoint(p.x-1, p.y+1)
	if bottomLeft != nil {
		resp = append(resp, bottomLeft)
	}

	left := g.GetPoint(p.x-1, p.y)
	if left != nil {
		resp = append(resp, left)
	}

	topLeft := g.GetPoint(p.x-1, p.y-1)
	if topLeft != nil {
		resp = append(resp, topLeft)
	}

	return resp
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
	totalFlashes := 0
	for i := 0; i < 100; i++ {
		totalFlashes += g.Step()
	}

	return totalFlashes
}

func problemTwo(g grid) int {
	step := 0
	for flashCount := -1; flashCount != 100; step++ {
		flashCount = g.Step()
	}

	return step
}

func main() {
	inputFilename := "input"

	g, err := LoadGrid(inputFilename)
	if err != nil {
		panic(err)
	}

	fmt.Println("Problem One:", problemOne(g))

	// recreate grid
	g, err = LoadGrid(inputFilename)
	if err != nil {
		panic(err)
	}

	fmt.Println("Problem Two:", problemTwo(g))
}
