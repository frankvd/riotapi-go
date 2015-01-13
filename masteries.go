package riotapi

// Mastery page
type MasteryPage struct {
	Masteries []map[string]int
	Id        int
	Name      string
	Current   bool
}

// Mastery
type Mastery struct {
	MasteryId int
	Rank      int
}

// Masteries
type Masteries struct {
	SummonerId int
	Pages      []MasteryPage
}
