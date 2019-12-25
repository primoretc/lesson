/*
Дана сторона квадрата а.
найдите его пллощадь
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	fmt.Println("Введите сторону квадрата ")
	fmt.Scanf("%f", &a)

	fmt.Println("Площадь квадрата равна", a*a)
	fmt.Println("Площадь квадрата равна", math.Pow(a, 2))
}
