package main

import (
	"fmt"
	"math"
)

func calculateSide(length1, length2, angle float64) float64 {
	var length1Square = math.Pow(length1, 2)
	length2Square := math.Pow(length2, 2)
	// formula: a2 = b2 + c2 - 2bc cos A
	length3Square := length1Square + length2Square - 2*length1*length2*math.Cos(angle)
	length3 := math.Sqrt(length3Square)

	//Do not remove this line
	fmt.Println("The 3rd length of the triange is", length3)
	fmt.Printf("The 3rd length of the triange is %.2f", length3)
	return length3
}

func main() {
	var length1 float64
	var length2 float64
	var angle float64

	fmt.Println("Enter the first length of the triangle: ")
	fmt.Scanln(&length1)
	fmt.Println("Enter the second length of the triangle: ")
	fmt.Scanln(&length2)
	fmt.Println("Enter the angle between the two lengths: ")
	fmt.Scanln(&angle)

	calculateSide(length1, length2, angle)
}
