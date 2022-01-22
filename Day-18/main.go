package main

import (
	"errors"
	"fmt"
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
	Depth       int
}

func (n *Node) Print() string {
	if n.Leaf {
		return fmt.Sprintf("%d", n.LeafValue)
	}

	return fmt.Sprintf("[%s,%s]", n.Left.Print(), n.Right.Print())
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
	for ; currentNode.Parent.Left == currentNode; currentNode = currentNode.Parent {
		if currentNode.Parent == nil {
			return nil
		}
	}
	currentNode = currentNode.Parent.Left

	for ; !currentNode.Leaf; currentNode = currentNode.Right {
	}

	return currentNode
}

func findNextRightNode(n *Node) *Node {
	currentNode := n
	for ; currentNode.Parent.Right == currentNode; currentNode = currentNode.Parent {
		if currentNode.Parent == nil {
			return nil
		}
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

func ParseNodes(bs *ByteStream, depth int, parent *Node) *Node {
	resp := &Node{
		Depth:  depth,
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
			resp.Left = ParseNodes(bs, depth+1, resp)
		case ',':
			bs.Pop()
			resp.Right = ParseNodes(bs, depth+1, resp)
		default:
			resp.Leaf = true
			resp.LeafValue = readNum(bs)
			return resp
		}
	}
}

func main() {
	// b := []byte("[[54,8],20]")
	b := []byte("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]")

	n := ParseNodes(NewByteStream(b), 0, nil)

	fmt.Println(n.Print())
	Explode(n.Left.Right.Right.Right)
	fmt.Println(n.Print())
	// fmt.Printf("%+v %p\n", n.Left.Right.Right.Right, n.Left.Right.Right.Right)

	// fmt.Printf("%+v\n", n.Left.Right.Right.Right.Left)
	// fmt.Printf("%+v\n", n.Left.Right.Right.Right.Right)

	// fmt.Println("==========")

	// n = findNextRightNode(n.Left.Right.Right.Right)
	// fmt.Printf("%+v\n", n)

}
