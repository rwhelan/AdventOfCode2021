package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

type LookupKey map[int]bool

type Pixel struct {
	X, Y  int
	IsLit bool
}

type row map[int]Pixel
type Image map[int]row

// func (img Image) Process(key LookupKey) Image {
// 	minX, maxX, minY, maxY := img.dimensions()

// }

func (img Image) PixelValue(X, Y int) int {

}

func (img *Image) AddPixel(X, Y int, IsLit bool) {
	_, ok := (*img)[Y]
	if !ok {
		(*img)[Y] = make(row)
	}

	(*img)[Y][X] = Pixel{X: X, Y: Y, IsLit: IsLit}
}

func (img *Image) Print() string {
	resp := strings.Builder{}

	minX, maxX, minY, maxY := img.dimensions()
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if img.GetPixel(x, y).IsLit {
				resp.WriteByte('#')
			} else {
				resp.WriteByte('.')
			}
		}
		resp.WriteByte('\n')
	}

	return resp.String()
}

func (img *Image) GetPixel(X, Y int) Pixel {
	row, ok := (*img)[Y]
	if !ok {
		return Pixel{
			X:     X,
			Y:     Y,
			IsLit: false,
		}
	}

	p, ok := row[X]
	if !ok {
		return Pixel{
			X:     X,
			Y:     Y,
			IsLit: false,
		}
	}

	return p
}

func (img *Image) dimensions() (minX, maxX, minY, maxY int) {
	for _, row := range *img {
		for _, p := range row {
			if p.X < minX {
				minX = p.X
			}

			if p.X > maxX {
				maxX = p.X
			}

			if p.Y < minY {
				minY = p.Y
			}

			if p.Y > maxY {
				maxY = p.Y
			}
		}
	}

	return
}

func ParseLookupKey(b []byte) LookupKey {
	resp := make(map[int]bool)

	for i, c := range b {
		if c == '.' {
			resp[i] = false
		} else {
			resp[i] = true
		}
	}

	return resp
}

func ParseImageData(b []byte) Image {
	resp := make(Image)
	for y, r := range bytes.Split(b, []byte("\n")) {
		resp[y] = make(row)
		for x, v := range r {
			if v == '.' {
				resp[y][x] = Pixel{X: x, Y: y, IsLit: false}
			} else {
				resp[y][x] = Pixel{X: x, Y: y, IsLit: true}
			}
		}
	}

	return resp
}

func ParseInput(b []byte) (Image, LookupKey) {
	input := bytes.Split(b, []byte("\n\n"))
	if len(input) != 2 {
		panic("Bad input file format")
	}

	return ParseImageData(input[1]), ParseLookupKey(input[0])
}

func main() {
	rawInput, err := ioutil.ReadFile("test")
	if err != nil {
		panic(err)
	}

	img, key := ParseInput(rawInput)

	fmt.Println(img, key)

	fmt.Println(img.Print())
}
