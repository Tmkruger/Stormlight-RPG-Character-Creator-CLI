package main

import (
	"fmt"
	"os"
	"strings"
)

type origin struct {
	ancestry    string
	nationality string
	forms       string
}

type character struct {
	name   string
	origin origin
	level  int
	paths  []string
	stats  map[string]int
}

func intro() (name string) {
	// Flavor Text for Introduction :)
	fmt.Println("\"Life before death, Strength before weakness, Journey before destination\"\n	- Oath of the Knights Radiant\n")
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
				name = answer
				fmt.Print("\033[1A\033[2K")
				fmt.Printf("Are you certain %v is your name? (y/n)\n", name)
				for {
					fmt.Scan(&answer)
					switch answer {
					case "y":
						return name
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
func originPicker() origin {
	o := origin{}
	println("We have many peoples among our rank's, where do you hail from? (Character Origin)")
	println("1.Human\n2.Singer")
	//ADD 1. OR 2. OPTIONS AS WELL
	var answer string
	for {
		fmt.Scan(&answer)
		answer = strings.ToLower(answer)
		fmt.Print("\033[1A\033[2K")
		switch answer {
		case "human", "1":
			o.ancestry = "Human"
			println("Which kingdom did you grow up in? (Nationality)")
			fmt.Println("1.Alethi	5.Reshi		9.Thaylen\n2.Azish 	6.Shin		10.Unkalaki\n3.Herdazian	7.Iriali	11.Veden\n4.Natan		8.Kharbranthian")
			for {
				fmt.Scan(&answer)
				answer = strings.ToLower(answer)
				fmt.Print("\033[1A\033[2K")
				switch answer {
				case "1", "alethi":
					o.nationality = "Alethi"
				case "2", "azish":
					o.nationality = "Azish"
				case "3", "herdazian":
					o.nationality = "Herdazian"
				case "4", "natan":
					o.nationality = "Natan"
				case "5", "reshi":
					o.nationality = "Reshi"
				case "6", "shin":
					o.nationality = "Shin"
				case "7", "iriali":
					o.nationality = "Iriali"
				case "8", "kharbranthian":
					o.nationality = "Kharbranthian"
				case "9", "thaylen":
					o.nationality = "Thaylen"
				case "10", "unkalaki":
					o.nationality = "Unkalaki"
				case "11", "veden":
					o.nationality = "Veden"
				default:
					fmt.Println("I don't understand, Can you repeat yourself? (only items from the list please)")
				}
				if o.nationality != "" {
					return o
				}
			}
		case "singer", "2":
			o.ancestry = "Singer"
			println("What forms do you know?")
			println("1.Finesse (Art Form | Nimble Form)\n2.Wisdom (Mediation Form | Scholar Form)\n3.Resolve (War Form | Work Form)")
			for {
				fmt.Scan(&answer)
				answer = strings.ToLower(answer)
				fmt.Print("\033[1A\033[2K")
				switch answer {
				case "1", "finesse":
					o.forms = "Finesse"
				case "2", "wisdom":
					o.forms = "Wisdom"
				case "3", "resolve":
					o.forms = "Resolve"
				default:
					fmt.Println("I don't understand, Can you repeat yourself? (Finesse/Wisdom/Resolve only please)")
				}
				if o.forms != "" {
					return o
				}
			}
		default:
			fmt.Println("I don't understand, Can you repeat yourself? (Items from the list only please)")
		}
	}
}

//func classPicker()

func main() {
	c := character{}
	c.name = intro()
	fmt.Print("\033[1A\033[2K")
	fmt.Printf("\nHello, %v\n", c.name)
	c.origin = originPicker()
	println(c.origin.ancestry)
	println(c.origin.nationality)
	println(c.origin.forms)
	//Start finding out what class and stats and such
	//Choose what level to create to. (max 3 for our sake)
	//Maybe Use AI or Premade Backstories or someform of backstory creator
	//Roll up all the data into a full blown character
	//Output to CLI or Official PDF
}
