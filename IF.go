package main

import "fmt"

func main() {
	age := 24
	// check if age is old enough to drive
	if age < 12 && age > 24 {
		fmt.Println("you shall play")
	} else if age < 24 || age > 20 {
		fmt.Println("you can't play")
	} else {
		fmt.Println("you shall not play")
	}
}
