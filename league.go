package riotapi

import (
	"sort"
	"strconv"
)

// League object
type League struct {
	Entries       []LeagueEntry
	Name          string
	ParticipantId string
	Queue         string
	Tier          string
}

// League entry
type LeagueEntry struct {
	Division         string
	IsFreshBlood     bool
	IsHotStreak      bool
	IsInactive       bool
	IsVeteran        bool
	LeaguePoints     int
	PlayerOrTeamId   string
	PlayerOrTeamName string
	Wins             int
	MiniSeries       struct {
		Losses   int
		Progress string
		Target   int
		Wins     int
	}
}

// Sort type
type ByLeaguePoints []LeagueEntry

func (entries ByLeaguePoints) Len() int      { return len(entries) }
func (entries ByLeaguePoints) Swap(i, j int) { entries[i], entries[j] = entries[j], entries[i] }
func (entries ByLeaguePoints) Less(i, j int) bool {
	return entries[i].LeaguePoints > entries[j].LeaguePoints
}

// League service
type LeagueService struct {
	*Service
}

// Returns the league of the summoner
func (service *LeagueService) ForSummoner(id int) League {
	stringid := strconv.Itoa(id)
	resp := service.Client.Call("league", stringid)

	leagues := map[string][]League{}
	service.Parser.Parse(resp.Body, &leagues)

	league := leagues[stringid][0]
	sort.Sort(ByLeaguePoints(league.Entries))

	return league
}
