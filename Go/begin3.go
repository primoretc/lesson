/*
Даны стороны прямоугольника a и b/
Найти его площадь S=a*b и периметр P=2*(a+b)
*/

package main

import "fmt"

func main() {
	var a float64
	var b float64
	fmt.Println("Введите стороны прямоугольника a и b")
	fmt.Scanf("%f", &a)
	fmt.Scanf("%f", &b)
	fmt.Println("Площадь прямоугольника равна ", a*b)
	fmt.Println("Периметр прямоугольника равен ", 2*(a+b))
}
