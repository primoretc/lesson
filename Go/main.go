package main

import (
	"fmt"
)

func summa(x int, y int) {
	var z = x + y
	fmt.Println(z)
}

func main() {
	var x int
	var y int
	fmt.Println("Введите первое число ")
	fmt.Scanln(&x)
	fmt.Println("Введите второе число ")
	fmt.Scanln(&y)
	summa(x, y)

}
