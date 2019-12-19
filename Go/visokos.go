package main

import (
	"fmt"
)

func main() {
	var y int
	fmt.Println("Введите год:")
	fmt.Scanln(&y)
	fmt.Println(visokos(y))
}
func visokos(y int) bool {
	var god bool
	if y%4 != 0 {
		fmt.Println("1")
		god = false
	} else if y%100 == 0 {
		if y%400 == 0 {
			fmt.Println("2")
			god = true
		} else {
			fmt.Println("3")
			god = false
		}
	} else if y%400 == 0 {
		fmt.Println("4")
		god = true
	} else {
		fmt.Println("5")
		god = true
	}
	return god
}
