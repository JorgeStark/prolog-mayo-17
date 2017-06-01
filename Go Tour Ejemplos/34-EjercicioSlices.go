package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {

    // Allocate two-dimensioanl array.
	a := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		a[i] = make([]uint8, dx)
	}

	// Do something.
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			//Solo para cambiar el color
			switch {
			case j % 2 == 0:
				a[i][j] = 240
			default:
				a[i][j] = 100
			}
		}
	}
	return a
}

func main() {
    pic.Show(Pic)
}
