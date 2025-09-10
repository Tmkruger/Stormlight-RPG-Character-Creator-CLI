package data

type Character struct {
	Name   		string
	Ancestry   	string
	Level  		int
	Paths  		[]string
	Stats  		map[string]int
	Expertises 	[]string
}