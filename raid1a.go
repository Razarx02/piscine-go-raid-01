package student

import "github.com/01-edu/z01"

func Raid1a(x, y int) {
	if x < 0 || y < 0 {
	} else {
		for yc := 0; yc < y; yc++ {
			for xc := 0; xc < x; xc++ {
				if xc == 0 || xc == x-1 {
					if yc == 0 || yc == y-1 {
						z01.PrintRune('o')
					} else {
						z01.PrintRune('|')
					}
				} else if yc == 0 || yc == y-1 {
					z01.PrintRune('-')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	}
}
