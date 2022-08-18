package main

import "golang.org/x/tour/pic"

// Pic returns a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers
func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		res[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			res[y][x] = uint8(x*x + y*y + 33)
		}
	}

	return res
}

func main() {
	// display the picture
	pic.Show(Pic)
}
