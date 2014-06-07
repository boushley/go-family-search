package models

type DisplayProperties struct {
	AscendancyNumber  string `json:"ascendancyNumber"`
	BirthDate         string `json:"birthDate"`
	BirthPlace        string `json:"birthPlace"`
	DeathDate         string `json:"deathDate"`
	DeathPlace        string `json:"deathPlace"`
	DescendancyNumber string `json:"descendancyNumber"`
	Gender            string `json:"gender"`
	Lifespan          string `json:"lifespan"`
	MarriageDate      string `json:"marriageDate"`
	MarriagePlace     string `json:"marriagePlace"`
	Name              string `json:"name"`
	Id                ID     `json:"id"`
}
