package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type point struct {
	x, y              int
	value, pcost, pqi int
	visited, finalp   bool
	prev              *point
}

type PointList []*point

func (p *PointList) Len() int {
	return len(*p)
}

func (p *PointList) Less(i, j int) bool {
	return (*p)[i].pcost < (*p)[j].pcost
}

func (p *PointList) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

type row map[int]*point
type grid map[int]row

func (g *grid) maxXY() (int, int) {
	return len((*g)[0]) - 1, len(*g) - 1
}

func (g grid) Print() string {
	resp := strings.Builder{}

	maxX, maxY := g.maxXY()

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			p := g.GetPoint(x, y)
			if p.finalp {
				resp.WriteString(
					fmt.Sprintf("\033[31m%d\033[0m", p.value),
				)
				continue
			}

			if p.visited {
				resp.WriteString(
					fmt.Sprintf("\033[1m%d\033[0m", p.value),
				)
			} else {
				resp.WriteString(
					fmt.Sprintf("%d", p.value),
				)
			}
		}
		resp.WriteString("\n")
	}

	return resp.String()
}

func (g grid) GetUnvisitedNeighbors(p *point) PointList {
	resp := make([]*point, 0, 4)

	above := g.GetPoint(p.x, p.y-1)
	if above != nil && !above.visited {
		resp = append(resp, above)
	}

	right := g.GetPoint(p.x+1, p.y)
	if right != nil && !right.visited {
		resp = append(resp, right)
	}

	bottom := g.GetPoint(p.x, p.y+1)
	if bottom != nil && !bottom.visited {
		resp = append(resp, bottom)
	}

	left := g.GetPoint(p.x-1, p.y)
	if left != nil && !left.visited {
		resp = append(resp, left)
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

func LoadGrid(gridData []byte) grid {
	resp := make(grid)

	for y, rowBytes := range bytes.Split(gridData, []byte{'\n'}) {
		resp[y] = make(row)
		for x, v := range rowBytes {
			value, err := strconv.Atoi(string(v))
			if err != nil {
				panic(err)
			}

			resp[y][x] = &point{
				x:       x,
				y:       y,
				value:   value,
				visited: false,
				pcost:   4294967296,
				pqi:     -1,
			}
		}

	}

	return resp
}

func main() {
	rawInput, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	g := LoadGrid(rawInput)
	pq := make(pqueue, 0, 1000)

	p := g.GetPoint(0, 0)
	p.pcost = 0

	pq.Push(p)

	maxX, maxY := g.maxXY()

	for {
		p = pq.Pop()

		if p.x == maxX && p.y == maxY {
			break
		}
		p.visited = true
		// fmt.Println(g.Print())

		for _, np := range g.GetUnvisitedNeighbors(p) {
			if np.pqi == -1 {
				pq.Push(np)
			}

			nc := p.pcost + np.value
			if nc < np.pcost {
				np.pcost = nc
				pq.Update(np)
				np.prev = p
			}
		}

		// time.Sleep(time.Millisecond * 10)
	}

	pathCost := p.pcost
	for p != nil {
		p.finalp = true
		p = p.prev
	}

	fmt.Println(g.Print())
	fmt.Println(pathCost)

}

// pqueue

type pqueue []*point

func (q *pqueue) swap(i, j int) {
	(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
	(*q)[i].pqi = i
	(*q)[j].pqi = j
}

func (q *pqueue) up(idx int) {
	for {
		parent := (idx - 1) / 2
		if (*q)[parent].pcost < (*q)[idx].pcost ||
			(*q)[parent].pcost == (*q)[idx].pcost {
			break
		}

		q.swap(parent, idx)
		idx = parent
	}
}

func (q *pqueue) down(idx int) bool {
	i := idx
	for {
		left := 2*i + 1
		right := left + 1

		if left >= len(*q) || right >= len(*q) {
			break
		}

		j := left
		if (*q)[right].pcost < (*q)[left].pcost {
			j = right
		}

		if !((*q)[j].pcost < (*q)[i].pcost) {
			break
		}

		q.swap(i, j)

		i = j
	}

	return i > idx
}

func (q *pqueue) Push(p *point) {
	p.pqi = len(*q)
	*q = append(*q, p)
	q.up(len(*q) - 1)
}

func (q *pqueue) Update(p *point) {
	if !q.down(p.pqi) {
		q.up(p.pqi)
	}
}

func (q *pqueue) Pop() *point {
	resp := (*q)[0]

	q.swap(0, len(*q)-1)
	(*q)[len(*q)-1] = nil
	*q = (*q)[:len(*q)-1]
	q.down(0)

	return resp
}
