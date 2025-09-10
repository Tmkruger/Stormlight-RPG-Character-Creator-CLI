package functions

import (
	"fmt"
	"os"
	"strings"

	"github.com/tmkruger/sl_character_gen/internal/data"
)


func Intro(c *data.Character) {
	// Flavor Text for Introduction :)
	fmt.Println("”Life before death, Strength before weakness, Journey before destination”\n	- Oath of the Knights Radiant\n")
	fmt.Println("Odium's forces are growing stronger, we need all the help we can get")
	fmt.Println("Are you willing to fight with us? (Would you like to make a new character?) y/n\n")
	var answer string
	for {
		fmt.Scan(&answer)
		//Determine if we want to make a character
		switch answer {
		case "n":
			fmt.Print("\033[1A\033[2K")
			fmt.Println("We understand, hopefully we will be enough...")
			os.Exit(0)
		case "y":
			//Wants to make character so get name
			for {
				fmt.Print("\033[1A\033[2K")
				fmt.Println("What can we call you? (Character Name)")
				fmt.Scan(&answer)
				name := answer
				fmt.Print("\033[1A\033[2K")
				fmt.Printf("Are you certain %v is your name? (y/n)\n", name)
				for {
					fmt.Scan(&answer)
					switch answer {
					case "y":
						c.Name = name
						return
					case "n":
						break
					default:
						fmt.Print("\033[1A\033[2K")
						fmt.Println("I don't understand, Can you repeat yourself? (y/n only please)")
						continue
					}
					break
				}
			}
		default:
			fmt.Print("\033[1A\033[2K")
			fmt.Println("I don't understand, Can you repeat yourself? (y/n only please)")
		}
	}
}

func AncestryPicker(c *data.Character) {
	println("“Before we continue, I must know... \nDo you walk the path of humankind, or do you sing with the rhythms of Roshar’s singers?” (Character Ancestry)")
	println("1.Human\n2.Singer")
	var answer string
	for {
		fmt.Scan(&answer)
		answer = strings.ToLower(answer)
		//Gets rid of what user typed as to not break immersion
		fmt.Print("\033[1A\033[2K")
		switch answer {
		case "human", "1":
			c.Ancestry = "Human"
			return

		case "singer", "2":
			c.Ancestry = "Singer"
			return

		default:
			fmt.Println("I don't understand, Can you repeat yourself? (Items from the list only please)")
		}
	}
	return
}

//func GetLevel() 


			//println("Which kingdom did you grow up in? (Nationality)")
			//fmt.Println("1.Alethi	5.Reshi		9.Thaylen\n2.Azish 	6.Shin		10.Unkalaki\n3.Herdazian	7.Iriali	11.Veden\n4.Natan		8.Kharbranthian")

			//println("What forms do you know?")
			//println("1.Finesse (Art Form | Nimble Form)\n2.Wisdom (Mediation Form | Scholar Form)\n3.Resolve (War Form | Work Form)")