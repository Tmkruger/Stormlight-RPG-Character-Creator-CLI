package main

import (
	"fmt"

	"github.com/tmkruger/sl_character_gen/internal/data"
	"github.com/tmkruger/sl_character_gen/internal/functions"
)

func main() {
	c := data.Character{}
	functions.Intro(&c)
	fmt.Print("\033[1A\033[2K")
	fmt.Printf("\nHello, %v\n", c.Name)
	functions.AncestryPicker(&c)
	
	println(c.Name)
	println(c.Ancestry)
	//Start finding out what class and stats and such
	//Choose what level to create to. (max 3 for our sake)
	//Maybe Use AI or Premade Backstories or someform of backstory creator
	//Roll up all the data into a full blown character
	//Output to CLI or Official PDF
}
