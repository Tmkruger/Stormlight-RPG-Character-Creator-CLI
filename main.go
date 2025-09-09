package main

import (
	"fmt"
	"os"
)

func main() {
	name := intro()
	fmt.Printf("Hello, %v\n", name)
}

func intro() (name string) {
	fmt.Println("\"Life before death, Strength before weakness, Journey before destination\" - Oath of the Knights Radiant")
	fmt.Println("Odium's forces are growing stronger, we need all the help we can get")
	fmt.Println("Are you willing to fight with us? (Would you like to make a new character?) y/n")
	var answer string
	for {
		fmt.Scan(&answer)
		// Determine if we want to make a character
		switch answer {
		case "n":
			fmt.Println("We understand, hopefully we will be enough...")
			os.Exit(0)
		case "y":
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
							fmt.Println("y or n only please")
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