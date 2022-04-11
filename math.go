package main

import (
	"fmt"
)

func main() {
	//addition
	fmt.Println(100 + 50)
	//multiplication
	fmt.Println(5 * 20)
	//division
	fmt.Println(10 / 5)
	//subtracting
	fmt.Println(10 - 5)
	//modulus returns the remainder
	fmt.Println(100 % 30)
	//increment
	number := 10

	number++ //(this adds 1 does the same thing 'number += 1 or 2 or 3....')
	fmt.Println(number)
	//decrement
	letter := 5

	letter-- //(this minuses 1 does the same thing 'letter -= 1 or 2 or 3....')
	fmt.Println(letter)

	//constant
	const p = 7 //(basically a variable but cant be changed)
	fmt.Println(p)

}
