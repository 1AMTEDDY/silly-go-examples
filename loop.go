package main

import (
	"fmt"
)

func main() {
	num := 13
	fmt.Println(5 > num || num == 13) //this is an "OR"(||) statement only one statement needs to be true for all statement to be true
	fmt.Println(5 > num && num == 13) // this is an "AND" (&&)statement only one statement needs to be false for all statement to be

}
