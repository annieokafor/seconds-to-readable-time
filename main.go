package main

import "fmt"

func main() {

	var t int = 86400
	var d int = 0
	var h int = 0
	var m int = 0

	d = t / 86400
	fmt.Println(d)

	r := t - (d * 86400)
	fmt.Println(r)

	h = r / 3600
	fmt.Println(h)

	r = r - (h * 3600)
	fmt.Println(r)

	m = r / 60
	fmt.Println(m)

	r = r - (m * 60)
	fmt.Println(r)

	fmt.Printf("%d days %d hours %d minutes %d seconds", d, h, m, r)

}
