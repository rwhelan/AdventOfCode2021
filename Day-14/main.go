package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func LoadFile(filename string) ([]byte, []byte, error) {
	rawInput, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, nil, err
	}

	split := bytes.Split(rawInput, []byte("\n\n"))
	return split[0], split[1], nil
}

func LoadPairs(b []byte) map[string]int {
	resp := make(map[string]int)
	for idx := 1; idx < len(b); idx++ {
		key := string(b[idx-1]) + string(b[idx])
		resp[key]++
	}

	return resp
}

func LoadRules(b []byte) map[string]string {
	resp := make(map[string]string)
	for _, row := range bytes.Split(b, []byte{'\n'}) {
		pieces := bytes.Split(row, []byte(" -> "))
		resp[string(pieces[0])] = string(pieces[1])

	}

	return resp
}

func applyRules(pairs map[string]int, rules map[string]string) map[string]int {
	resp := make(map[string]int)

	for p, c := range pairs {
		if insert, ok := rules[p]; ok {
			resp[string(p[0])+insert] += c
			resp[insert+string(p[1])] += c
		}
	}

	return resp
}

func RulesRunner(rules map[string]string) func(map[string]int) map[string]int {
	return func(p map[string]int) map[string]int {
		return applyRules(p, rules)
	}
}

func Count(pairs map[string]int) map[string]int {
	resp := make(map[string]int)
	for k, v := range pairs {
		resp[string(k[0])] += v
	}

	return resp
}

func MaxMin(counts map[string]int) (int, int) {
	max := 0
	min := (1 << 63) - 1

	for _, v := range counts {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	return max, min
}

func main() {
	p, r, err := LoadFile("input")
	if err != nil {
		panic(err)
	}

	rules := LoadRules(r)
	pairs := LoadPairs(p)
	step := RulesRunner(rules)

	for i := 0; i < 40; i++ {
		pairs = step(pairs)
	}

	counts := Count(pairs)
	counts[string(p[len(p)-1])]++

	max, min := MaxMin(counts)
	fmt.Println("Problem Two:", max-min)
}
