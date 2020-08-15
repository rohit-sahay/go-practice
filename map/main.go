package main

import "fmt"

type mymap map[string]string

func main() {
	colors := mymap{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff",
	}

	colors.printMap()
}

func (m mymap) printMap() {
	for color, hex := range m {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
