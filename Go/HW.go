package main

import (
	"fmt"
	"math"
	"os"
)

// comment
func main() {
	var nm string
	fmt.Println("Введите имя  ")
	fmt.Scanln(&nm)
	fmt.Println("HW " + nm)
	fmt.Println(os.Getgid())
	fmt.Println(math.Sqrt(625))
}
