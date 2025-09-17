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

type Specialties struct {
	id      string
	Name    string
	Talents []string
}
type Talent struct {
	Name        string
	Req         Prereq
	Description string
}
type Prereq struct {
	Skill      string
	SkillLevel int
}

type Path struct {
	id            string
	Name          string
	Type          string
	Description   string
	KeyTalent     string
	StartingSkill string
	Specialties   Specialties
	Attributes    []string
	Skills        []string
	MultiPaths    []string
}
