package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func ReadLines(filename string) ([][]byte, error) {
	rawInput, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	rawInput = bytes.TrimRight(rawInput, "\n")
	return bytes.Split(rawInput, []byte("\n")), nil
}

func problemOne(rows [][]byte) int {
	total := 0
	for _, k := range rows {
		d := k[61:]
		for _, s := range bytes.Split(d, []byte{' '}) {
			l := len(s)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				total++
			}
		}
	}

	return total
}

func main() {
	rows, err := ReadLines("input")
	if err != nil {
		panic(err)
	}

	fmt.Println("Problem One:", problemOne(rows))
}
