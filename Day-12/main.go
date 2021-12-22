package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/rwhelan/coding-challenge/trains/Go/pkg/graph"
)

func ReadLines(filename string) ([][]byte, error) {
	rawInput, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	rawInput = bytes.TrimRight(rawInput, "\n")
	return bytes.Split(rawInput, []byte("\n")), nil
}

func ParseGraph(filename string) (*graph.Graph, error) {
	inputRows, err := ReadLines(filename)
	if err != nil {
		return nil, err
	}

	resp := graph.NewGraph("Tunnels")

	for _, row := range inputRows {
		p := bytes.Split(row, []byte{'-'})

		var leftNode, rightNode *graph.Node

		if leftNode = resp.GetNode(string(p[0])); leftNode == nil {
			leftNode = graph.NewNode(string(p[0]))
			resp.AddNode(leftNode)
		}

		if rightNode = resp.GetNode(string(p[1])); rightNode == nil {
			rightNode = graph.NewNode(string(p[1]))
			resp.AddNode(rightNode)
		}

		leftNode.AddEdge(rightNode, 1)
		rightNode.AddEdge(leftNode, 1)

	}

	return resp, nil
}

func isSmallCave(node *graph.Node) bool {
	for _, b := range []byte(node.Name) {
		if b < 97 || b > 122 {
			return false
		}
	}

	return true
}

func GraphWalk(path *graph.Path, nextNode *graph.Node) graph.WalkerInstruction {
	if path.CurrentNode().Name == "end" {
		return graph.PATH_STOP
	}

	for _, node := range path.Nodes {
		if nextNode == node && isSmallCave(node) {
			return graph.PATH_DROP
		}
	}

	return graph.PATH_CONTINUE
}

func problemOne(g *graph.Graph) int {
	walkfunc := func(path *graph.Path, nextNode *graph.Node) graph.WalkerInstruction {
		if path.CurrentNode().Name == "end" {
			return graph.PATH_STOP
		}

		for _, node := range path.Nodes {
			if nextNode == node && isSmallCave(node) {
				return graph.PATH_DROP
			}
		}

		return graph.PATH_CONTINUE
	}

	paths, err := g.Walk(g.GetNode("start"), walkfunc)
	if err != nil {
		panic(err)
	}

	return paths.Len()
}

func main() {
	g, err := ParseGraph("input")
	if err != nil {
		panic(err)
	}

	fmt.Println("Problem One:", problemOne(g))
}
