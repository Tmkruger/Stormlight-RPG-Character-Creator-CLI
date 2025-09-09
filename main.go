package main

import (
	"fmt"
	"os"
)

func main() {
	name := intro()
	fmt.Printf("Hello, %v\n.", name)
	//Start finding out what class and stats and such
	//Choose what level to create to. (max 3 for our sake)
	//Maybe Use AI or Premade Backstories or someform of backstory creator
	//Roll up all the data into a full blown character
	//Output to CLI or Official PDF
}

func intro() (name string) {
	// Flavor Text for Introduction :)
	fmt.Println("\"Life before death, Strength before weakness, Journey before destination\" - Oath of the Knights Radiant")
	fmt.Println("Odium's forces are growing stronger, we need all the help we can get")
	fmt.Println("Are you willing to fight with us? (Would you like to make a new character?) y/n")
	var answer string
	for {
		fmt.Scan(&answer)
		//Determine if we want to make a character
		switch answer {
		case "n":
			fmt.Println("We understand, hopefully we will be enough...")
			os.Exit(0)
		case "y":
			//Wants to make character so get name
			for {
				fmt.Println("What can we call you? (Character Name)")
				fmt.Scan(&answer)
				name = answer
				fmt.Printf("Are you certain %v is your name? (y/n)\n", name)
				for {
					fmt.Scan(&answer)
					switch answer {
						case "y":
							return name
						case "n":
							break
						default:
							fmt.Println("I don't understand, Can you repeat yourself? (y/n only please)")
							continue
					}
					break
				}
			}

		default:
			fmt.Println("I don't understand, Can you repeat yourself? (y/n only please)")
		}
	}
}