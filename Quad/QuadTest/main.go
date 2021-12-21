package main

import (
	"Quad"
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) == 4 {
		x, _ := strconv.Atoi(os.Args[1])
		y, _ := strconv.Atoi(os.Args[2])

		question := os.Args[3]

		if question == "A" || question == "B" || question == "C" || question == "D" || question == "E" {

			switch question {
			case "A":
				Quad.QuadA(x, y)
			case "B":
				Quad.QuadB(x, y)
			case "C":
				Quad.QuadC(x, y)
			case "D":
				Quad.QuadD(x, y)
			case "E":
				Quad.QuadE(x, y)
			}
		} else {

			fmt.Println("Usage: go run main.go [x] [y] [q]")
			fmt.Println("q options are A,B,C,D,E")
			fmt.Println("Example: go run main.go [2] [3] [A]")
		}

	} else {
		fmt.Println("Must input 3 arguments")
		fmt.Println("Usage: go run main.go [x] [y] [q]")

	}
}
