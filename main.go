package main

import (
	"fmt"

	"newage/LAB2/mathutil"
	"newage/LAB2/strop"
)

func main() {
	text := "Hello Go"
	fmt.Println("Reverse:", strop.Reverse(text))
	fmt.Println("Vowels:", strop.CountVowels(text))
	fmt.Println("Factorial:", mathutil.Factorial(5))
	fmt.Println("Power:", mathutil.Power(2, 4))
}
