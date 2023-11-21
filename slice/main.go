package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// Initializing slice which contains slices
	image := make([][]uint8, dy)

	// Loops around map and make a new slice at every occurence
	for x := 0; x < dy; x++ {
		image[x] = make([]uint8, dx)

		// Initializing slice
		for y := 0; y < dx; y++ {
			image[x][y] = uint8(x + y/2)
		}
	}
	return image
}

func main() {
	pic.Show(Pic)
}
