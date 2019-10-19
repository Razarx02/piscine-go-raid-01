package student

import "github.com/01-edu/z01"

func Raid1a(x, y int) {
	if x < 0 || y < 0 {
		return
	}
	y2 := y - 1
	x2 := x - 1
	for y1 := 0; y1 < y; y1++ {
		if y1 != 0 && y1 != y2 {
			for x1 := 0; x1 < x; x1++ {
				if x1 == 0 || x1 == x2 {
					z01.PrintRune('*')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune(10)
		} else {
			for x1 := 0; x1 < x; x1++ {
				if x1 == 0 && y1 == 0 || x1 == x2 && y1 == y2 && x != 1 {
					z01.PrintRune('/')
				} else if x1 == x2 && y1 == 0 || x1 == 0 && y1 == y2 {
					z01.PrintRune(92)
				} else {
					z01.PrintRune('*')
				}
			}
			z01.PrintRune(10)
		}
	}
}
