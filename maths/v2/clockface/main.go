package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(secondsinradians())
}

func zero() float64 {
	return 0.0
}

func secondsinradians() float64 {
	return (math.Pi / (30 / (float64(zero()))))
}

// func main() {
// 	fmt.Println(10.0 / 0.0) // fails to compile
// }

// func main() {
// 	fmt.Println(10.0 / zero())
// }
