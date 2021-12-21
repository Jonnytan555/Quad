package Quad

import "github.com/01-edu/z01"

func QuadA(x, y int) {

	for c := 0; c <= y-1; c++ {
		if c == 0 || c == y-1 {
			for i := 0; i <= x-1; i++ {
				if i > 0 && i < x-1 {
					z01.PrintRune('-')
				} else {
					z01.PrintRune('o')
				}
			}
		} else {
			for i := 0; i <= x-1; i++ {
				if i > 0 && i < x-1 {
					z01.PrintRune(' ')
				} else {
					z01.PrintRune('|')
				}
			}
		}
		z01.PrintRune('\n')
	}
}

func QuadB(x, y int) {
	for c := 0; c <= y-1; c++ {
		if c == 0 {
			for i := 0; i <= x-1; i++ {
				if i == 0 {
					z01.PrintRune(47)
				} else if i == x-1 {
					z01.PrintRune(92)
				} else {
					z01.PrintRune('*')
				}
			}
		} else if c == y-1 {
			for i := 0; i <= x-1; i++ {
				if i == 0 {
					z01.PrintRune(92)
				} else if i == x-1 {
					z01.PrintRune(47)
				} else {
					z01.PrintRune('*')
				}
			}
		} else {
			for i := 0; i <= x-1; i++ {
				if i == 0 {
					z01.PrintRune('*')
				} else if i == x-1 {
					z01.PrintRune('*')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}

func QuadC(x, y int) {
	for c := 0; c <= y-1; c++ {
		if c == 0 {
			for i := 0; i <= x-1; i++ {
				if i == 0 {
					z01.PrintRune('A')
				} else if i == x-1 {
					z01.PrintRune('A')
				} else {
					z01.PrintRune('B')
				}
			}
		} else if c == y-1 {
			for i := 0; i <= x-1; i++ {
				if i == 0 {
					z01.PrintRune('C')
				} else if i == x-1 {
					z01.PrintRune('C')
				} else {
					z01.PrintRune('B')
				}
			}
		} else {
			for i := 0; i <= x-1; i++ {
				if i == 0 || i == x-1 {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}

func QuadD(x, y int) {
	for c := 0; c <= y-1; c++ {
		if c == 0 || c == x-1 {
			for i := 0; i <= x-1; i++ {
				if i := 0; i <= x-1 {
					z01.PrintRune('A')
				} else if i == x-1 {
					z01.PrintRune('C')
				} else {
					z01.PrintRune('B')
				}
			}
		} else {
			for i := 0; i <= x-1; i++ {
				if i == 0 {
					z01.PrintRune('B')
				} else if i == x-1 {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}

func QuadE(x, y int) {
	for c := 0; c <= y-1; c++ {
		if c == 0 {
			for i := 0; i <= x-1; i++ {
				if i == 0 {
					z01.PrintRune('A')
				} else if i == x-1 {
					z01.PrintRune('A')
				} else {
					z01.PrintRune('B')
				}
			}
		} else if c == y-1 {
			for i := 0; i <= x-1; i++ {
				if i == 0 {
					z01.PrintRune('C')
				} else if i == x-1 {
					z01.PrintRune('C')
				} else {
					z01.PrintRune('B')
				}
			}
		} else {
			for i := 0; i <= x-1; i++ {
				if i == 0 {
					z01.PrintRune('B')
				} else if i == x-1 {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}
