package main

import (
	"fmt"
	"math"
)

type Beacon struct {
	X, Y, Z int
}

type Scanner struct {
	X, Y, Z int
	ID      int
	Beacons []*Beacon
}

func (s Scanner) Clone() Scanner {
	b := make([]*Beacon, len(s.Beacons))
	for i, beacon := range s.Beacons {
		b[i] = &Beacon{
			X: beacon.X,
			Y: beacon.Y,
			Z: beacon.Z,
		}
	}

	return Scanner{
		X:       s.X,
		Y:       s.Y,
		Z:       s.Z,
		ID:      s.ID,
		Beacons: b,
	}
}

func BeaconDist(one, two *Beacon) float64 {
	xc := math.Pow(float64(two.X-one.X), 2)
	yc := math.Pow(float64(two.Y-one.Y), 2)
	zc := math.Pow(float64(two.Z-one.Z), 2)

	return math.Sqrt(xc + yc + zc)
}

func ScannerDist(s *Scanner) []float64 {
	distDeltas := make([]float64, 0)
	for i, b1 := range s.Beacons {
		for _, b2 := range s.Beacons[i+1 : len(s.Beacons)] {
			distDeltas = append(distDeltas, BeaconDist(b1, b2))
		}
	}

	return distDeltas
}

func AreNeighborScanners(one, two *Scanner) bool {
	oneDist := ScannerDist(one)
	twoDist := ScannerDist(two)

	dists := make(map[float64]int)
	for _, d := range oneDist {
		dists[d]++
	}

	for _, d := range twoDist {
		dists[d]++
	}

	for k, v := range dists {
		if v == 1 {
			delete(dists, k)
		}
	}

	return len(dists) >= 66
}

func main() {
	scanners, err := ReadScanners("input")
	if err != nil {
		panic(err)
	}

	for i, scanner := range scanners {
		for _, scanner2 := range scanners[i:len(scanners)] {
			if scanner.ID != scanner2.ID && AreNeighborScanners(scanner, scanner2) {
				fmt.Println(scanner.ID, " - ", scanner2.ID)
			}
		}
	}

	AreNeighborScanners(scanners[0], scanners[3])
}
