package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

type Reader struct {
	b   []byte
	idx int
}

func (r *Reader) Read(i int) []byte {
	resp := r.b[r.idx : r.idx+i]
	r.idx += i

	return resp
}

func (r *Reader) Remaining() int {
	return len(r.b[r.idx:])
}

type Packet struct {
	version, typeId int
	subPackets      []*Packet
	literalValue    int
	isLiteral       bool
}

func (p *Packet) values() int {
	subValues := make([]int, 0)
	for _, v := range p.subPackets {
		subValues = append(subValues, v.values())
	}

	switch p.typeId {
	case 0:
		return sum(subValues)
	case 1:
		return product(subValues)
	case 2:
		return min(subValues)
	case 3:
		return max(subValues)
	case 4:
		return p.literalValue
	case 5:
		return grt(subValues)
	case 6:
		return lst(subValues)
	case 7:
		return eql(subValues)
	default:
		panic("Unknown Packet Type")
	}
}

func versionSum(p *Packet) int {
	sub := 0
	for _, v := range p.subPackets {
		sub += versionSum(v)
	}

	return sub + p.version

}

// ToDo: Make Not Stupid
func hexToBinStr(b []byte) []byte {
	lookup := map[byte][]byte{
		'0': []byte("0000"),
		'1': []byte("0001"),
		'2': []byte("0010"),
		'3': []byte("0011"),
		'4': []byte("0100"),
		'5': []byte("0101"),
		'6': []byte("0110"),
		'7': []byte("0111"),
		'8': []byte("1000"),
		'9': []byte("1001"),
		'A': []byte("1010"),
		'B': []byte("1011"),
		'C': []byte("1100"),
		'D': []byte("1101"),
		'E': []byte("1110"),
		'F': []byte("1111"),
	}

	resp := make([]byte, 0, len(b)*4)
	for _, v := range b {
		resp = append(resp, lookup[v]...)
	}

	return resp
}

func binStrToInt(b []byte) int {
	v, err := strconv.ParseUint(string(b), 2, 64)
	if err != nil {
		panic(err)
	}

	return int(v)
}

func NewReader(b []byte) *Reader {
	return &Reader{
		b: b,
	}
}

func readLiteral(r *Reader) int {
	resp := 0
	for {
		cont := binStrToInt(r.Read(1))
		resp = (resp << 4) + binStrToInt(r.Read(4))
		if cont == 0 {
			return resp
		}
	}
}

func readPacket(r *Reader) *Packet {
	resp := &Packet{
		version: binStrToInt(r.Read(3)),
		typeId:  binStrToInt(r.Read(3)),
	}

	if resp.typeId == 4 {
		resp.isLiteral = true
		resp.literalValue = readLiteral(r)
		return resp
	}

	// type ID
	if binStrToInt(r.Read(1)) == 0 {
		bitSize := binStrToInt(r.Read(15))
		nr := NewReader(r.Read(bitSize))

		for nr.Remaining() != 0 {
			resp.subPackets = append(resp.subPackets, readPacket(nr))
		}
	} else {
		c := binStrToInt(r.Read(11))
		for i := 0; i < c; i++ {
			resp.subPackets = append(resp.subPackets, readPacket(r))
		}
	}

	return resp
}

func main() {
	rawInput, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	// r := NewReader(hexToBinStr([]byte("D8005AC2A8F0")))
	r := NewReader(hexToBinStr(rawInput))

	p := readPacket(r)

	fmt.Println("Problem One:", versionSum(p))
	fmt.Println("Problem Two:", p.values())
}

func sum(i []int) int {
	resp := 0
	for _, v := range i {
		resp += v
	}

	return resp
}

func max(i []int) int {
	resp := 0
	for _, v := range i {
		if v > resp {
			resp = v
		}
	}

	return resp
}

func min(i []int) int {
	resp := (1 << 62)
	for _, v := range i {
		if v < resp {
			resp = v
		}
	}

	return resp
}

func product(i []int) int {
	resp := 1
	for _, v := range i {
		resp *= v
	}

	return resp
}

func grt(i []int) int {
	if i[0] > i[1] {
		return 1
	}

	return 0
}

func lst(i []int) int {
	if i[0] > i[1] {
		return 0
	}

	return 1
}

func eql(i []int) int {
	if i[0] == i[1] {
		return 1
	}

	return 0
}
