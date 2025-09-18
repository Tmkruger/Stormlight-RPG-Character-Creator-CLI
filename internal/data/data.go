package data

import (
	"encoding/json"
	"os"
)

type Character struct {
	Name       string
	Ancestry   string
	Status     string
	Level      int
	Paths      Paths
	Stats      map[string]int
	Expertises []string
	Talents    []string
}

type Specialty struct {
	Id          string		`json:"id"`
	Name        string		`json:"name"`
	Description string		`json:"description"`
	Talents     []Talent	`json:"talents"`
}
type Talent struct {
	Id           string			`json:"id"`
	Name         string			`json:"name"`
	Path         string			`json:"path"`
	Specialty	 string			`json:"specialty"`
	Prereq_Stat  map[string]int	`json:"prereq_stat"`
	Prereq_Skill string		`json:"prereq_skill"`
	Activation   string			`json:"activation"`
	Description  string			`json:"description"`
	Summary      string			`json:"summary"`
}
type Prereq_Stat struct {
	Stat      string
	Level 	  int	
}

type PreReq_Skill struct {
	Skill string
}

type Paths struct {
	Heroic  []HeroicPath	`json:"heroicPaths"`
	Radiant []RadiantPath	`json:"radiantPaths"`
}
type HeroicPath struct {
	Id            string		`json:"id"`
	Name          string		`json:"name"`
	Description   string		`json:"description"`
	KeyTalent     Talent		`json:"keyTalent"`
	StartingSkill string		`json:"startingSkill"`
	Specialties   []Specialty	`json:"specialties"`
	Attributes    []string		`json:"attributes"`
	Skills        []string		`json:"skills"`
	MultiPaths    []string		`json:"multipaths"`
	Rewards       []string		`json:"rewards"`
}

type RadiantPath struct {
	Id          string	`json:"id"`
	Name        string	`json:"name"`
	Description string	`json:"description"`
	Philosophy  string	`json:"philosophy"`
	Spren       string	`json:"spren"`
	Surges      []string	`json:"surges"`
	KeyTalent   Talent	`json:"keyTalent"`
	Specialties []Specialty	`json:"specialties"`
	MultiPaths  []string	`json:"multipaths"`
}

func ReadPaths() (Paths, error) {
	data, err := os.ReadFile("internal/data/paths.json")
	if err != nil {
		return Paths{}, err
	}
	return ParsePathData(data)
}

func ParsePathData(data []byte) (Paths, error) {
	var paths Paths
	err := json.Unmarshal(data, &paths)
	if err != nil {
		return Paths{}, err
	}
	return paths, nil
}
