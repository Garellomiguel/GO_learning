package main

import "fmt"

func main() {
	// var colors map[string]string
	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff000",
		"green": "#14126",
	}

	colors["withe"] = "#fffff"

	delete(colors, "white")
	fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for k, v := range c {
		fmt.Printf("Key is %v and value %v \n", k, v)
	}
}
