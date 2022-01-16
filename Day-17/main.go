package main

import "fmt"

type Probe struct {
	xv, yv, x, y int
	path         [][]int
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

func FindX(givenYV int, t *Target) *Probe {
	for xv := 1; ; xv++ {
		p := &Probe{yv: givenYV, xv: xv}
		if p.HitsTarget(t) {
			return p
		}

		if p.IsToFar(t) {
			return nil
		}
	}
}

func FindMaxY(t *Target) *Probe {
	for yv := 1000; ; yv-- {
		// fmt.Println("Trying:", yv)
		if p := FindX(yv, t); p != nil {
			return p
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

func main() {
	// target area: x=139..187, y=-148..-89

	// target := &Target{
	// 	xLeft:   20,
	// 	xRight:  30,
	// 	yTop:    -5,
	// 	yBottom: -10,
	// }

	target := &Target{
		xLeft:   139,
		xRight:  187,
		yTop:    -89,
		yBottom: -148,
	}

	p := FindMaxY(target)
	fmt.Println("Problem One:", MaxY(p.path))
}
