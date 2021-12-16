package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
)

type Stack struct {
	s   []byte
	idx int
}

func NewStack(l int) *Stack {
	return &Stack{
		s:   make([]byte, l),
		idx: -1,
	}
}

func (s *Stack) Push(b byte) {
	s.idx++
	s.s[s.idx] = b
}

func (s *Stack) Pop() byte {
	b := s.s[s.idx]
	s.idx--
	return b
}

func (s *Stack) Peek() byte {
	return s.s[s.idx]
}

func ReadLines(filename string) ([][]byte, error) {
	rawInput, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	rawInput = bytes.TrimRight(rawInput, "\n")
	return bytes.Split(rawInput, []byte("\n")), nil
}

func rev(b byte) byte {
	m := map[byte]byte{
		'(': ')',
		')': '(',
		'{': '}',
		'}': '{',
		'[': ']',
		']': '[',
		'<': '>',
		'>': '<',
	}

	r, ok := m[b]
	if !ok {
		panic("You f'ed up")
	}

	return r
}

func Validate(row []byte) (bool, int, byte) {
	stack := NewStack(len(row) / 2)

	for i, c := range row {
		if c == '<' || c == '(' || c == '{' || c == '[' {
			stack.Push(c)

		} else if stack.Pop() != rev(c) {
			return false, i, c

		}
	}

	return true, 0, 0
}

func Complete(row []byte) []byte {
	stack := NewStack(len(row) / 2)

	for _, c := range row {
		if c == '<' || c == '(' || c == '{' || c == '[' {
			stack.Push(c)

		} else {
			// input rows should already be validated
			stack.Pop()
		}
	}

	resp := make([]byte, stack.idx+1)
	for i := 0; i <= stack.idx; i++ {
		resp[stack.idx-i] = rev(stack.s[i])
	}

	return resp
}

func problemTwo(rows [][]byte) int {
	scoreChar := func(b byte) int {
		switch b {
		case ')':
			return 1
		case ']':
			return 2
		case '}':
			return 3
		case '>':
			return 4
		}

		return 0
	}

	scores := []int{}
	for _, r := range rows {
		rowScore := 0
		if ok, _, _ := Validate(r); !ok {
			// invalid set of instructions
			continue
		}

		for _, c := range Complete(r) {
			rowScore = (rowScore * 5) + scoreChar(c)
		}

		scores = append(scores, rowScore)
	}

	sort.Ints(scores)
	return scores[len(scores)/2]
}

func problemOne(rows [][]byte) int {
	scoreChar := func(b byte) int {
		switch b {
		case ')':
			return 3
		case ']':
			return 57
		case '}':
			return 1197
		case '>':
			return 25137
		}

		return 0
	}

	scores := make(map[byte]int)
	for _, r := range rows {
		valid, _, char := Validate(r)
		if !valid {
			scores[char] += scoreChar(char)
		}
	}

	total := 0
	for _, v := range scores {
		total += v
	}

	return total
}

func main() {
	rawRows, err := ReadLines("input")
	if err != nil {
		panic(err)
	}

	fmt.Println("Problem One:", problemOne(rawRows))
	fmt.Println("Problem Two:", problemTwo(rawRows))
}
