package student

import "github.com/01-edu/z01"

func Raid1d(x, y int) {
	if x > 0 && y > 0 {
		for yc := 0; yc < y; yc++ {
			for xc := 0; xc < x; xc++ {
				if xc == 0 || xc == x-1 || yc == 0 || yc == y-1 {
					if xc == 0 && (yc == 0 || yc == y-1) {
						z01.PrintRune('A')
					} else if xc == x-1 && (yc == 0 || yc == y-1) {
						z01.PrintRune('C')
					} else {
						z01.PrintRune('B')
					}
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	}
}
