package main

import "fmt"

func main() {
	var c, d, e float64
	fmt.Print("Masukan harga: ")
	fmt.Scanln(&c)
	fmt.Print("Masukan diskon: ")
	fmt.Scanln(&d)
	d = d / 100
	e = c - (c * d)

	fmt.Println("Harga Diskon  ", int(e))
}
