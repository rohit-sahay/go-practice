package main

import "fmt"

func main() {
	a := []int{}

	for i := 0; i <= 10; i++ {
		a = append(a, i)
	}

	for _, i := range a {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}
}
