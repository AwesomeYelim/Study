package main

import "fmt"

func main() {
	age := 17

	if age >= 20 {
		fmt.Println("성인")
	} else {
		fmt.Println("급식")
		fmt.Printf("%d 년후 성인임", 20-age)
	}
}