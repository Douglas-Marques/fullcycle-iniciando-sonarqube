package main

import "fmt"

func main() {
	fmt.Print(Soma(10, 10))
}

func Soma(a int, b int) int {
	return a + b
}

func Sub(a int, b int) int {
	return a - b
}

func Times(a int, b int) int {
	return a * b
}

func Divide(a int, b int) int {
	return a / b
}

func Power(a int, b int) int {
	return a ^ b
}
