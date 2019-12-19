package main

import (
	"fmt"
<<<<<<< HEAD
)

type vehicle interface {
	move()
}

func main() {
	fmt.Println("Hello Go")
=======
	"strconv"
)

func summa(x int, y int) (z int) {
	z = x + y
	//fmt.Println(z)
	return z
}

func main() {
	var x int
	var y int
	fmt.Println("Введите первое число ")
	fmt.Scanln(&x)
	fmt.Println("Введите второе число ")
	fmt.Scanln(&y)
	//s := summa(x, y)
	st1 := strconv.Itoa(x)
	st2 := strconv.Itoa(y)
	fmt.Println("Сумма двух чисел "+st1+" и "+st2+" равна", summa(x, y))

>>>>>>> 1a7efb558f63211742d5a53a1e6bd0d418ffe56d
}
