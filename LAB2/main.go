package main

import (
	"fmt"

	"LAB2/mathutil"
	"LAB2/strop"
)

func main() {
	text := "Hello Go"
	fmt.Println("Reverse:", strop.Reverse(text))
	fmt.Println("Vowels:", strop.CountVowels(text))
	fmt.Println("Factorial:", mathutil.Factorial(5))
	fmt.Println("Power:", mathutil.Power(2, 4))
}
