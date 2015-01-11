package riotapi

type RunePage struct {
	Id    int
	Slots []map[string]int
}

type Runes struct {
	SummonerId int
	Pages      []MasteryPage
}

type Rune struct {
	Rank   int
	RuneId int
}
