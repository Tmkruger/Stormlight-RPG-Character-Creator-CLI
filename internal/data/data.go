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

type Specialties struct {
	Id          string
	Name        string
	Description string
	Talents     []Talent
}
type Talent struct {
	Id           string
	Name         string
	Prereq_Stat  []Prereq
	Prereq_Skill []Prereq
	Activation   string
	Description  string
	Summary      string
}
type Prereq struct {
	Skill      string
	SkillLevel int
}

type Paths struct {
	Heroic  []HeroicPath
	Radiant []RadiantPath
}
type HeroicPath struct {
	Id            string
	Name          string
	Description   string
	KeyTalent     Talent
	StartingSkill string
	Specialties   Specialties
	Attributes    []string
	Skills        []string
	MultiPaths    []string
	Rewards       []string
}

type RadiantPath struct {
	Id          string
	Name        string
	Description string
	Philosophy  string
	Spren       string
	Surges      []string
	KeyTalent   Talent
	Specialties Specialties
	MultiPaths  []string
}

func ReadPaths() (Paths, error) {
	data, err := os.ReadFile("data/paths.json")
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
