package functions

import (
	"fmt"
	"strings"

	"github.com/tmkruger/sl_character_gen/internal/data"
)

// Core App Functions
func CustomIntro(c *data.Character, mainMenu *bool) {
	//Remove Main Menu and Flavor Text for Introduction :)
	RmLineAbove(7)
	fmt.Println("”Life before death, Strength before weakness, Journey before destination”\n	- Oath of the Knights Radiant\n")
	fmt.Println("”Ah, there you are. You look like someone standing on the edge of a story.\nTell me...do you intend to leap in, blades flashing and heart pounding?\nOr will you sit back and watch others risk their lives while you sip wine and critique the prose?\nYes or no, friend: are you joining this tale?\n")
	fmt.Println("Would you like to make a new character? (y/n)\n")
	var answer string
	for {
		answer = GetUserInput("Enter your choice: ")
		//Determine if we want to make a character
		switch answer {
		case "n", "no":
			RmLineAbove(2)
			fmt.Println("In another tale than...\n")
			return
		case "y", "yes":
			//Wants to make character so get name
			*mainMenu = false
			RmLineAbove(3)
			fmt.Println("\n”Every good story needs a name for its hero. Or villain. Or tragic fool, depending on how things go.\nWhat shall I call you when I whisper your deeds to the winds?” (Character Name)\n")
			for {
				name := GetUserInput("Enter your character name: ")
				RmLineAbove(1)
				fmt.Printf("\nAre you certain %v is your name? (y/n)\n", name)
				for {
					answer = GetUserInput("\nEnter your choice: ")
					switch answer {
					case "y":
						RmLineAbove(2)
						c.Name = name
						return
					case "n":
						RmLineAbove(1)
						break
					default:
						RmLineAbove(1)
						fmt.Println("I don't understand, Can you repeat yourself? (y/n only please)")
						continue
					}
					break
				}
			}
		default:
			RmLineAbove(1)
			fmt.Println("I don't understand, Can you repeat yourself? (y/n only please)")
		}
	}
}

func AncestryPicker(c *data.Character) {
	fmt.Printf("“Ah, well than %v. About your lineage, the tapestry from which you were cut.\nWere your ancestors proud Alethi, diligent Shin, wandering Horneaters, or something stranger still?\nDon’t be shy, I’ve heard worse answers than 'illegitimate goat spawn,' and that turned out quite the tale.“ (Character Ancestry)\n", c.Name)
	println("\n1.Human\n2.Singer")
	var answer string
	for {
		answer = GetUserInput("\nEnter your choice: ")
		RmLineAbove(3)
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
}

func GetHumanNationality(c *data.Character) {
	fmt.Println("\n“Ah, a child of humankind. But even you lot squabble over borders and banners.\nWere you raised beneath Alethkar’s war drums, the scholarly domes of Jah Keved, the stone valleys of Shinovar, or perhaps Thaylenah’s bustling markets?\nTell me, from which patch of soil did your story sprout?“\n")
	println("Which kingdom did you grow up in? (Character Nationality)\n")
	fmt.Println("1.Alethi	5.Reshi		9.Thaylen\n2.Azish 	6.Shin		10.Unkalaki\n3.Herdazian	7.Iriali	11.Veden\n4.Natan		8.Kharbranthian")
	var answer string
	for {
		answer = GetUserInput("Enter your choice: ")
		RmLineAbove(7)
		switch answer {
		case "1", "Alethi":
			c.Expertises = append(c.Expertises, "Alethi")
			return

		case "2", "Azish":
			c.Expertises = append(c.Expertises, "Azish")
			return

		case "3", "Herdazian":
			c.Expertises = append(c.Expertises, "Herdazian")
			return

		case "4", "Natan":
			c.Expertises = append(c.Expertises, "Natan")
			return

		case "5", "Reshi":
			c.Expertises = append(c.Expertises, "Reshi")
			return

		case "6", "Shin":
			c.Expertises = append(c.Expertises, "Shin")
			return

		case "7", "Iriali":
			c.Expertises = append(c.Expertises, "Iriali")
			return

		case "8", "Kharbranthian":
			c.Expertises = append(c.Expertises, "Kharbranthian")
			return

		case "9", "Thaylen":
			c.Expertises = append(c.Expertises, "Thaylen")
			return

		case "10", "Unkalaki":
			c.Expertises = append(c.Expertises, "Unkalaki")
			return

		case "11", "Veden":
			c.Expertises = append(c.Expertises, "Veden")
			return

		default:
			fmt.Println("I don't understand, Can you repeat yourself? (Items from the list only please)")
			break
		}
	}
}

func GetSingerNationality(c *data.Character) {
	println("“And where, pray tell, do you skulk these days? Perhaps hiding beneath Alethkar’s banners, sweeping floors in a Thaylen market, or serving wine to some pompous lighteyes in Jah Keved?\nTell me which kingdom’s cobblestones you’re haunting, it helps me place you in the right chapter of this little drama.“ (Cultural Expertise)\n")
	println("1.Alethi	5.Reshi		9.Thaylen\n2.Azish 	6.Shin		10.Unkalaki\n3.Herdazian	7.Iriali	11.Veden\n4.Natan		8.Kharbranthian")
	var answer string
	for {
		answer = GetUserInput("Enter your choice: ")
		switch answer {
		case "1", "alethi":
			c.Expertises = append(c.Expertises, "Alethi")
			return
		case "2", "azish":
			c.Expertises = append(c.Expertises, "Azish")
			return
		case "3", "herdazian":
			c.Expertises = append(c.Expertises, "Herdazian")
			return
		case "4", "natan":
			c.Expertises = append(c.Expertises, "Natan")
			return
		case "5", "reshi":
			c.Expertises = append(c.Expertises, "Reshi")
			return
		case "6", "shin":
			c.Expertises = append(c.Expertises, "Shin")
			return
		case "7", "iriali":
			c.Expertises = append(c.Expertises, "Iriali")
			return
		case "8", "kharbranthian":
			c.Expertises = append(c.Expertises, "Kharbranthian")
			return
		case "9", "thaylen":
			c.Expertises = append(c.Expertises, "Thaylen")
			return
		case "10", "unkalaki":
			c.Expertises = append(c.Expertises, "Unkalaki")
			return
		case "11", "veden":
			c.Expertises = append(c.Expertises, "Veden")
			return
		default:
			fmt.Println("I don't understand, Can you repeat yourself? (Items from the list only please)")
		}
	}
}

func GetForms(c *data.Character) {
	println("“Ah, a Singer! How delightful, you folk have a knack for turning a melody into muscle and sinew.\nTell me, which forms have you learned to dance with? Warform for strength? Workform for clever hands? Perhaps even something… more secret? Don’t be shy;\nI’m a collector of stories, not a tattler… usually.“ (Singer Forms)\n")
	println("1.Finesse (Art Form | Nimble Form)\n2.Wisdom (Mediation Form | Scholar Form)\n3.Resolve (War Form | Work Form)\n")
	var answer string
	for {
		answer = GetUserInput("Enter your choice: ")
		switch answer {
		case "1", "finesse":
			c.Talents = append(c.Talents, "Finesse")
			return
		case "2", "wisdom":
			c.Talents = append(c.Talents, "Wisdom")
			return
		case "3", "resolve":
			c.Talents = append(c.Talents, "Resolve")
			return
		default:
			fmt.Println("I don't understand, Can you repeat yourself? (Items from the list only please)")
		}
	}
}

func GetSingerStatus(c *data.Character) {
	RmLineAbove(4)
	println("“Mmm, one more thing, friend of rhythms. Are you quietly watching for the Listeners, gathering secrets behind an innocent smile?\nOr have you worn the slave’s brands, bending knee while plotting freedom—or perhaps just survival?\nDon’t worry, your answer will stay between us… probably.“ (Status)\n")
	println("1.Spy\n2.Slave\n")
	var answer string
	for {
		answer = GetUserInput("Enter your choice: ")
		switch answer {
		case "1", "spy":
			c.Status = "Spy"
			c.Expertises = append(c.Expertises, "Listener")
			return
		case "2", "slave":
			c.Status = "Slave"
			return
		default:
			fmt.Println("I don't understand, Can you repeat yourself? (Items from the list only please)")
		}
	}
}

func GetLevel(c *data.Character) {
	println("\n“Let’s speak of experience. Are you but a wide-eyed squire fumbling with a wooden blade,\nor a seasoned veteran whose scars could each tell a book’s worth of stories?\nIn other words, what level of legend do you claim?“ (Character Level 1-3)\n")
	for {
		fmt.Print("Enter your choice: ")
		var choice string
		fmt.Scanln(&choice)
		RmLineAbove(1)
		switch choice {
		case "1":
			c.Level = 1
			return
		case "2":
			c.Level = 2
			return
		case "3":
			c.Level = 3
			return
		default:
			fmt.Println("I don't understand, Can you repeat yourself? (Items from the list only please)")
		}
	}
}

func GetPaths(c *data.Character) {
	println("“The world offers many roads, and few of them are straight. Which path started your tale?\nDo you follow the ideals of Radiance? Tread the blood soaked road of Odium’s servants?\nOr perhaps you wander, untethered, carving your own crooked trail.\nWhich path will you stumble down, oh eager traveler?“ (Starting Path)")
}

func GetAttributes(c *data.Character) {
	println("“Ah, the bones of your tale—strength, wit, charm, and all the little numbers that will later decide whether you climb triumphantly or plummet hilariously. Tell me, in what qualities do you shine, and in which do you stumble like a drunk chull?“")
}

func GetTalents(c *data.Character) {
	println("“Every protagonist needs a trick. Shall we say you’re handy with a blade? Or perhaps your tongue spins lies smoother than satin? Maybe you juggle knives while reciting dirty limericks—that would be memorable. What talents do you bring to the tale?“")
}

func GetEquipment(c *data.Character) {
	println("“Heroes and fools alike are only as good as the tools they carry. Sword, spear, boots without holes, or a lucky rock named Kevin—what shall I place in your hands as you stride into destiny?“")
}

func GetBackstory(c *data.Character) {
	println("“At last, we come to the marrow of the matter: your past. Every soul drags behind them a cart of stories, some glorious, some… less so. What tales haunt you, drive you, or whisper in your ear as you walk into tomorrow?“")
}

// Helper Functions
func RmLineAbove(num int) {
	for i := 0; i < num; i++ {
		fmt.Printf("\033[1A\033[2K")
	}
}

func GetUserInput(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	input = strings.ToLower(input)
	RmLineAbove(1)
	return input
}
