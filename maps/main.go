package main

import (
	"fmt"
)

func main() {
	//One way of declare Maps
	colors_iterator := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#0FFFFFFF",
	} // This map used for map iterator example at bottom

	fmt.Println(colors_iterator)

	// Another possible ways
	//var colors map[string]string  Declare empty maps

	color_code := make(map[string]string) //Declare empty maps, This map used for add and delete functionality of map below

	fmt.Println(color_code)

	//To add data to map
	color_code["White"] = "#0FFFFFF"
	fmt.Println(color_code)
	//Delete data in map
	delete(color_code, "White")
	fmt.Println("After Delete of white color", color_code)

	mapIterator(colors_iterator)

}

func mapIterator(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
