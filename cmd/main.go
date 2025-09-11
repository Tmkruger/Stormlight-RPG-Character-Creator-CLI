package main

import (
	"fmt"
	"os"

	"github.com/tmkruger/sl_character_gen/internal/data"
	"github.com/tmkruger/sl_character_gen/internal/functions"
)

func intro(c *data.Character, mainMenu *bool) {
	for *mainMenu == true {
		fmt.Println("\nWelcome to the Stormlight RPG Character Creator!")
		fmt.Println("1. Create a new custom character")
		fmt.Println("2. Load an existing character")
		fmt.Println("3. Go through ”The First Step” Adventure to make a character")
		fmt.Println("4. Exit\n")
		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scanln(&choice)
		fmt.Print("\033[1A\033[2K")
		switch choice {
		case 1:
			functions.CustomIntro(c, mainMenu)
		case 2:
			println("Loading Character...")
			// functions.LoadExistingCharacter(&c)
		case 3:
			println("Starting The First Step Adventure...")
			//functions.FirstStepAdventure(&c)
		case 4:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice!")
		}
	}
}

func main() {
	c := data.Character{}
	mainMenu := true
	for mainMenu == true {
		intro(&c, &mainMenu)
	}
	functions.AncestryPicker(&c)
	functions.GetLevel(&c)

	println(c.Name)
	println(c.Ancestry)
	println(c.Level)
	//Start finding out what class and stats and such
	//Choose what level to create to. (max 3 for our sake)
	//Maybe Use AI or Premade Backstories or someform of backstory creator
	//Roll up all the data into a full blown character
	//Output to CLI or Official PDF
}
