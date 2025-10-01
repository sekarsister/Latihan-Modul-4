package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b, c, e, g, f float64
	fmt.Scan(&a, &b)
	fmt.Scan(&c, &e)
	fmt.Scan(&g, &f)

	ab := math.Sqrt(math.Pow(c-a, 2) + math.Pow(e-b, 2))

	bc := math.Sqrt(math.Pow(g-c, 2) + math.Pow(f-e, 2))

	ca := math.Sqrt(math.Pow(a-g, 2) + math.Pow(b-f, 2))

	ST := ab
	if bc > ST {
		ST = bc
	}
	if ca > ST {
		ST = ca
	}

	fmt.Printf("%f\n", ST)
}
