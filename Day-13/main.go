package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type point struct {
	x, y   int
	marked bool
}

type foldInstruction struct {
	direction string
	value     int
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

func (g *grid) maxXY() (maxX, maxY int) {
	for y := range *g {
		for x := range (*g)[y] {
			if p := g.GetPoint(x, y); p != nil {
				if p.marked && y > maxY {
					maxY = y
				}

				if p.marked && x > maxX {
					maxX = x
				}
			}
		}
	}

	return
}

func (g *grid) Mark(x, y int) *grid {
	if p := g.GetPoint(x, y); p == nil {
		if _, ok := (*g)[y]; !ok {
			(*g)[y] = make(row)
		}

		(*g)[y][x] = &point{
			x:      x,
			y:      y,
			marked: true,
		}

	} else {
		(*g)[y][x].marked = true

	}

	return g
}

func (g *grid) Print() string {
	maxX, maxY := g.maxXY()
	resp := strings.Builder{}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			p := g.GetPoint(x, y)
			if p == nil || !p.marked {
				resp.WriteByte('.')
			} else {
				resp.WriteByte('#')
			}
		}
		resp.WriteByte('\n')
	}

	return resp.String()
}

func (g *grid) FoldY(r int) *grid {
	maxX, maxY := g.maxXY()

	for y := r; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if p := g.GetPoint(x, y); p != nil && p.marked {
				g.Mark(x, r-(y-r))
			}
		}

		delete(*g, y)
	}

	return g
}

func (g *grid) FoldX(c int) *grid {
	maxX, maxY := g.maxXY()

	for y := 0; y <= maxY; y++ {
		for x := c; x <= maxX; x++ {
			if p := g.GetPoint(x, y); p != nil && p.marked {
				g.Mark(c-(x-c), y)
			}

			delete((*g)[y], x)
		}
	}

	return g
}

func LoadFile(filename string) ([]byte, []byte, error) {
	rawInput, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, nil, err
	}

	split := bytes.Split(rawInput, []byte("\n\n"))
	return split[0], split[1], nil
}

func LoadGrid(gridData []byte) grid {
	resp := make(grid)

	for _, rowBytes := range bytes.Split(gridData, []byte{'\n'}) {
		rowPieces := bytes.Split(rowBytes, []byte{','})

		x, err := strconv.Atoi(string(rowPieces[0]))
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(string(rowPieces[1]))
		if err != nil {
			panic(err)
		}

		if _, ok := resp[y]; !ok {
			resp[y] = make(row)
		}

		resp[y][x] = &point{
			x:      x,
			y:      y,
			marked: true,
		}

	}

	return resp
}

func (g *grid) Fold(instr *foldInstruction) *grid {
	switch instr.direction {
	case "y":
		g.FoldY(instr.value)
	case "x":
		g.FoldX(instr.value)
	default:
		panic("unknown instruction type")
	}

	return g
}

func LoadFoldInstructions(instData []byte) []*foldInstruction {
	rows := bytes.Split(instData, []byte{'\n'})
	resp := make([]*foldInstruction, len(rows))

	for i, rowBytes := range rows {
		value, err := strconv.Atoi(string(rowBytes[13:]))
		if err != nil {
			panic(err)
		}

		resp[i] = &foldInstruction{
			direction: string(rowBytes[11]),
			value:     value,
		}
	}

	return resp
}

func main() {
	gridData, instData, err := LoadFile("input")
	if err != nil {
		panic(err)
	}

	g := LoadGrid(gridData)
	instructions := LoadFoldInstructions(instData)

	g.Fold(instructions[0])

	count := 0
	maxX, maxY := g.maxXY()
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if p := g.GetPoint(x, y); p != nil && p.marked {
				count++
			}
		}
	}

	fmt.Println("Problem One:", count)

	for _, instr := range instructions[1:] {
		g.Fold(instr)
	}

	fmt.Println("\nProblem Two:")
	fmt.Println(g.Print())
}
