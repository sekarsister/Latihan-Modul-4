package main

import (
	"fmt"
	"math"
)

func main() {
	var c, d, e float64
	fmt.Print("MASUKAN BMI: ")
	fmt.Scanln(&c)
	fmt.Print("MASUKAN TINGGI: ")
	fmt.Scanln(&d)
	e = c * (math.Pow(d, 2))

	fmt.Println("Berat Badan  ", int(math.Round(e)))
}
