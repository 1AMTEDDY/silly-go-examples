package main

import "fmt"

func main() {
	names := []string{"peter", "kim", "emma", "peter"}
	//for num := 2000; num >= 0; num -= 20 {  //this counts from 2000 to 0 with a difference of 20
	//fmt.Println(num)}
	for c := 0; c < len(names); c++ {
		if names[c] == "peter" {
			fmt.Println("hello peter")
		}
	}
}
