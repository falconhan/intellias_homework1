package main

import "fmt"

func main() {

	var (
		n int
		a int
		b int
	)
	// считываем числа пока не будет введен 0
	for fmt.Scan(&a); a != 0; fmt.Scan(&a) {
		fmt.Println(a)
	}
	for fmt.Scan(&b); b != 0; fmt.Scan(&b) {
		fmt.Println(b)
	}
	for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
		fmt.Println(n)
	}
}

