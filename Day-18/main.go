package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

// == Byte Streaming

type ByteStream struct {
	b   []byte
	idx int
}

type ByteStreamEOF struct{}

func (*ByteStreamEOF) Error() string {
	return "EOF"
}

func (bs *ByteStream) Pop() (byte, error) {
	if bs.idx > len(bs.b) {
		return 0, &ByteStreamEOF{}
	}

	bs.idx++
	return bs.b[bs.idx-1], nil
}

func (bs *ByteStream) Peek() (byte, error) {
	if bs.idx > len(bs.b) {
		return 0, &ByteStreamEOF{}
	}

	return bs.b[bs.idx], nil
}

func NewByteStream(b []byte) *ByteStream {
	return &ByteStream{
		b: b,
	}
}

// === /ByteStreaming

type Node struct {
	Parent      *Node
	Left, Right *Node
	Leaf        bool
	LeafValue   int
}

func (n *Node) Magnitude() int {
	if n.Leaf {
		return n.LeafValue
	}

	return 3*n.Left.Magnitude() + 2*n.Right.Magnitude()
}

func (n *Node) Print() string {
	if n.Leaf {
		return fmt.Sprintf("%d", n.LeafValue)
	}

	return fmt.Sprintf("[%s,%s]", n.Left.Print(), n.Right.Print())
}

func (n *Node) Depth() int {
	c := n
	depth := 0
	for ; c.Parent != nil; c = c.Parent {
		depth++
	}

	return depth
}

func Add(n, o *Node) *Node {
	resp := &Node{
		Left:  n,
		Right: o,
	}

	n.Parent = resp
	o.Parent = resp

	return resp
}

func Walk(n *Node, f func(*Node) bool) bool {
	if n.Left != nil {
		if Walk(n.Left, f) {
			return true
		}
	}

	if n.Right != nil {
		if Walk(n.Right, f) {
			return true
		}
	}

	return f(n)
}

func Process(n *Node) {
	exp := func(n *Node) bool {
		if n.Depth() >= 4 && !n.Leaf {
			Explode(n)
			return true
		}
		return false
	}

	spl := func(n *Node) bool {
		if n.Leaf && n.LeafValue > 9 {
			Split(n)
			return true
		}
		return false
	}

	for {
		if Walk(n, exp) {
			continue
		}

		if Walk(n, spl) {
			continue
		}

		break
	}
}

func readNum(bs *ByteStream) int {
	v := strings.Builder{}
	for {
		nextByte, err := bs.Peek()
		if err != nil {
			panic("Should Never End On Number " + v.String())
		}

		if nextByte == ',' || nextByte == ']' {
			value, err := strconv.ParseUint(v.String(), 10, 64)
			if err != nil {
				panic("OOPS " + v.String() + " is not a number")
			}

			return int(value)
		}

		v.WriteByte(nextByte)
		bs.Pop()
	}
}

func findNextLeftNode(n *Node) *Node {
	currentNode := n
	for {
		if currentNode.Parent == nil {
			return nil
		}

		if currentNode.Parent.Left != currentNode {
			break
		}

		currentNode = currentNode.Parent
	}

	currentNode = currentNode.Parent.Left

	for ; !currentNode.Leaf; currentNode = currentNode.Right {
	}

	return currentNode
}

func findNextRightNode(n *Node) *Node {
	currentNode := n
	for {
		if currentNode.Parent == nil {
			return nil
		}

		if currentNode.Parent.Right != currentNode {
			break
		}

		currentNode = currentNode.Parent
	}

	currentNode = currentNode.Parent.Right

	for ; !currentNode.Leaf; currentNode = currentNode.Left {
	}

	return currentNode
}

func Explode(n *Node) {
	left := findNextLeftNode(n)
	right := findNextRightNode(n)

	if left != nil {
		left.LeafValue += n.Left.LeafValue
	}

	if right != nil {
		right.LeafValue += n.Right.LeafValue
	}

	n.Right = nil
	n.Left = nil

	n.Leaf = true
	n.LeafValue = 0
}

func Split(n *Node) {
	leftVal := int(math.Floor(float64(n.LeafValue) / 2))
	rightVal := int(math.Ceil(float64(n.LeafValue) / 2))

	n.Left = &Node{Leaf: true, LeafValue: leftVal, Parent: n}
	n.Right = &Node{Leaf: true, LeafValue: rightVal, Parent: n}

	n.Leaf = false
	n.LeafValue = 0
}

func ParseNodes(bs *ByteStream, parent *Node) *Node {
	resp := &Node{
		Parent: parent,
	}

	for {
		c, err := bs.Peek()
		if errors.Is(err, &ByteStreamEOF{}) {
			return resp
		}

		if c == ']' {
			bs.Pop()
			return resp
		}

		switch c {
		case '[':
			bs.Pop()
			resp.Left = ParseNodes(bs, resp)

		case ',':
			bs.Pop()
			resp.Right = ParseNodes(bs, resp)

		default: // Leaf Node Number
			resp.Leaf = true
			resp.LeafValue = readNum(bs)

			return resp
		}
	}
}

func ParseNum(b []byte) *Node {
	return ParseNodes(NewByteStream(b), nil)
}

func ReadLines(filename string) ([][]byte, error) {
	rawInput, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	rawInput = bytes.TrimRight(rawInput, "\n")
	return bytes.Split(rawInput, []byte("\n")), nil
}

func main() {
	rows, err := ReadLines("input")
	if err != nil {
		panic(err)
	}

	numbers := make([]*Node, len(rows))
	for i, r := range rows {
		numbers[i] = ParseNum(r)
	}

	root := numbers[0]
	for _, n := range numbers[1:] {
		root = Add(root, n)
		Process(root)
	}

	fmt.Println("Problem One:", root.Magnitude())

	max := 0
	for _, i := range rows {
		for _, j := range rows {
			n := Add(ParseNum(i), ParseNum(j))
			Process(n)

			if n.Magnitude() > max {
				max = n.Magnitude()
			}
		}
	}

	fmt.Println("Problem Two:", max)

}
