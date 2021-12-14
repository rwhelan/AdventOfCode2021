package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
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

func strMatch(str string, subStr string) bool {
	for i := 0; i < len(subStr); i++ {
		if !strings.Contains(str, string(subStr[i])) {
			return false
		}
	}
	return true
}

func orderBytes(b []byte) []byte {
	ps := strings.Split(string(b), "")
	sort.Strings(ps)
	return []byte(strings.Join(ps, ""))
}

func sortInput(row []byte) [][]byte {
	resp := make([][]byte, 0, 4)
	pieces := bytes.Split(row, []byte{' '})

	for _, p := range pieces {
		resp = append(resp, orderBytes(p))
	}

	return resp
}

func generateLookupKey(b [][]byte) map[string]int {
	sort.Slice(b, func(i, j int) bool {
		return len(b[i]) < len(b[j])
	})

	resp := make(map[string]int)

	var oneKey string
	var fourFrag string
	for _, i := range b {
		switch len(i) {
		case 2:
			resp[string(i)] = 1
			oneKey = string(i)

		case 3:
			resp[string(i)] = 7

		case 4:
			resp[string(i)] = 4
			four := string(i)
			for i := 0; i < len(oneKey); i++ {
				four = strings.ReplaceAll(four, string(oneKey[i]), "")
			}
			fourFrag = four

		case 5:
			if strMatch(string(i), oneKey) {
				resp[string(i)] = 3
			} else if strMatch(string(i), fourFrag) {
				resp[string(i)] = 5
			} else {
				resp[string(i)] = 2
			}

		case 6:
			if !strMatch(string(i), fourFrag) {
				resp[string(i)] = 0
			} else if strMatch(string(i), oneKey) {
				resp[string(i)] = 9
			} else {
				resp[string(i)] = 6
			}

		case 7:
			resp[string(i)] = 8

		}
	}

	return resp

}

func solveRow(row []byte) (int, error) {
	key := sortInput(row[:58])
	question := sortInput(row[61:])

	answerKey := generateLookupKey(key)
	var resp string
	for _, q := range question {
		a, ok := answerKey[string(q)]
		if !ok {
			panic("This Shouldn't happen")
		}

		resp += strconv.Itoa(a)
	}

	return strconv.Atoi(resp)
}

func problemTwo(rows [][]byte) int {
	total := 0
	for _, r := range rows {
		a, _ := solveRow(r)
		total += a
	}

	return total
}

func main() {
	rows, err := ReadLines("input")
	if err != nil {
		panic(err)
	}

	fmt.Println("Problem One:", problemOne(rows))
	fmt.Println("Problem Two:", problemTwo(rows))
}
