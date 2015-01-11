package riotapi

type MasteryPage struct {
	Masteries []map[string]int
	Id        int
	Name      string
	Current   bool
}

type Mastery struct {
	MasteryId int
	Rank      int
}

type Masteries struct {
	SummonerId int
	Pages      []MasteryPage
}
