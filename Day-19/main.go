package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Beacon struct {
	X, Y, Z int
}

type Scanner struct {
	X, Y, Z int
	ID      int
	Beacons []*Beacon
}

func ParseScanner(b []byte) (*Scanner, error) {
	rows := bytes.Split(b, []byte("\n"))
	var err error
	scanner := &Scanner{}
	if rows[0][13] == ' ' {
		if scanner.ID, err = strconv.Atoi(string(rows[0][12])); err != nil {
			return nil, err
		}

	} else {
		if scanner.ID, err = strconv.Atoi(string(rows[0][12:14])); err != nil {
			return nil, err
		}

	}

	scanner.Beacons = make([]*Beacon, len(rows[1:]))
	for i, b := range rows[1:] {
		scanner.Beacons[i], err = ParseBeacon(b)
		if err != nil {
			return nil, fmt.Errorf("unable to parse beacon data: %s", b)
		}
	}

	return scanner, nil

}

func ParseBeacon(b []byte) (*Beacon, error) {
	p := bytes.Split(b, []byte(","))
	if len(p) != 3 {
		return nil, fmt.Errorf("malformed beacon data: %s", b)
	}

	x, err := strconv.Atoi(string(p[0]))
	if err != nil {
		return nil, fmt.Errorf("unable to parse X from beacon row: %s", p)
	}

	y, err := strconv.Atoi(string(p[1]))
	if err != nil {
		return nil, fmt.Errorf("unable to parse Y from beacon row: %s", p)
	}

	z, err := strconv.Atoi(string(p[2]))
	if err != nil {
		return nil, fmt.Errorf("unable to parse Z from beacon row: %s", p)
	}

	return &Beacon{
		X: x,
		Y: y,
		Z: z,
	}, nil
}

func ReadScanners(filename string) ([]*Scanner, error) {
	rawInput, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	rawScannerText := bytes.Split(rawInput, []byte("\n\n"))
	scanners := make([]*Scanner, len(rawScannerText))

	for i, s := range rawScannerText {
		scanners[i], err = ParseScanner(s)
		if err != nil {
			return nil, err
		}
	}

	return scanners, nil
}

func main() {
	scanners, err := ReadScanners("input")
	if err != nil {
		panic(err)
	}

	for _, scanner := range scanners {
		fmt.Printf("%+v\n", scanner)
		for _, beacon := range scanner.Beacons {
			fmt.Printf("%+v\n", beacon)
		}
	}
}
