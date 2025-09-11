package data

type Character struct {
	Name       string
	Ancestry   string
	Status     string
	Level      int
	Paths      []string
	Stats      map[string]int
	Expertises []string
	Talents    []string
}
