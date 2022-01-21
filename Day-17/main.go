package main

import (
	"fmt"
)

type Probe struct {
	xv, yv   int
	ixv, iyv int
	x, y     int
	path     [][]int
}

func NewProbe(xv, yv int) *Probe {
	return &Probe{
		xv:  xv,
		ixv: xv,
		yv:  yv,
		iyv: yv,
		x:   0,
		y:   0,
	}
}

type Target struct {
	xLeft, xRight, yBottom, yTop int
}

func (p *Probe) Step() {
	p.x += p.xv
	p.y += p.yv

	switch {
	case p.xv > 0:
		p.xv--
	case p.xv < 0:
		p.xv++
	}

	p.yv--

	p.path = append(p.path, []int{p.x, p.y})
}

func (p *Probe) HitsTarget(t *Target) bool {
	for {
		if p.IsBelow(t) || p.IsToFar(t) {
			return false
		}

		if p.IsInTarget(t) {
			return true
		}

		p.Step()
	}
}

func (p *Probe) IsInTarget(t *Target) bool {
	return p.x >= t.xLeft && p.x <= t.xRight &&
		p.y >= t.yBottom && p.y <= t.yTop
}

func (p *Probe) IsToFar(t *Target) bool {
	return p.x > t.xRight
}

func (p *Probe) IsBelow(t *Target) bool {
	return p.y < t.yBottom
}

func SearchDownForX(t *Target) *Probe {
	for yv := 1000; ; yv-- {
		for xv := 1; xv < 1000; xv++ {
			p := NewProbe(xv, yv)
			if p.HitsTarget(t) {
				return p
			}
		}
	}
}

func SearchUpForX(t *Target) *Probe {
	for yv := -1000; ; yv++ {
		for xv := 0; xv < 1000; xv++ {
			p := NewProbe(xv, yv)
			if p.HitsTarget(t) {
				return p
			}
		}
	}
}

func MaxY(i [][]int) int {
	max := 0
	for _, v := range i {
		if v[1] > max {
			max = v[1]
		}
	}

	return max
}

func FindAll(t *Target) [][]int {
	resp := make([][]int, 0)

	MaxY := SearchDownForX(t).iyv
	MinY := SearchUpForX(t).iyv

	for yv := MinY; yv <= MaxY; yv++ {
		for xv := 0; xv <= 300; xv++ {
			p := NewProbe(xv, yv)
			if p.HitsTarget(t) {
				resp = append(resp, []int{xv, yv})
			}
		}
	}

	return resp
}

func main() {
	// Sample Target
	// target := &Target{
	// 	xLeft:   20,
	// 	xRight:  30,
	// 	yTop:    -5,
	// 	yBottom: -10,
	// }

	// target area: x=139..187, y=-148..-89
	target := &Target{
		xLeft:   139,
		xRight:  187,
		yTop:    -89,
		yBottom: -148,
	}

	p := SearchDownForX(target)
	fmt.Println("Problem One:", MaxY(p.path))

	v := FindAll(target)
	fmt.Println("Problem Two:", len(v))
}
