package riotapi

// Runepage
type RunePage struct {
	Id    int
	Slots []map[string]int
}

// Runes
type Runes struct {
	SummonerId int
	Pages      []MasteryPage
}

// Rune
type Rune struct {
	Rank   int
	RuneId int
}
