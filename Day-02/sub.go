package main

// Problem One
type SubmarineOne struct {
	X, Y int
}

func (s *SubmarineOne) MoveUp(steps int) *SubmarineOne {
	s.Y -= steps
	return s
}

func (s *SubmarineOne) MoveDown(steps int) *SubmarineOne {
	s.Y += steps
	return s
}

func (s *SubmarineOne) MoveForward(steps int) *SubmarineOne {
	s.X += steps
	return s
}

// Problem Two
type SubmarineTwo struct {
	X, Y, Aim int
}

func (s *SubmarineTwo) MoveUp(steps int) *SubmarineTwo {
	s.Aim -= steps
	return s
}

func (s *SubmarineTwo) MoveDown(steps int) *SubmarineTwo {
	s.Aim += steps
	return s
}

func (s *SubmarineTwo) MoveForward(steps int) *SubmarineTwo {
	s.X += steps
	s.Y += s.Aim * steps
	return s
}
