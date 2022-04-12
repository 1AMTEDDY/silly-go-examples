package main

import "fmt"

// print numbers 0-100 but if
// number is divisible by 3, print "fizz"
// number is divisible by 5, print "buzz"
// number is divisible by both 3 and 5 print "FizzBuzz"

func main() {
	for num := 0; num <= 100; num++ {
		if num%15 == 0 {
			fmt.Println("Fizz")
		} else if num%5 == 0 {
			fmt.Println("Buzz")

		} else if num%3 == 0 {
			fmt.Println("FizzBuzz")
		} else {
			fmt.Println(num)
		}
	}
}
