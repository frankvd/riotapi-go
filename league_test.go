package riotapi

import (
	"io"
	"sort"
	"testing"
)

func TestGetLeague(t *testing.T) {
	service := new(LeagueService)
	service.Service = NewFakeService()

	service.Parser.(*FakeParser).ParseFunc = func(response io.Reader, ret interface{}) {
		pointer := ret.(*map[string][]League)
		*pointer = map[string][]League{
			"1": {
				League{Entries: []LeagueEntry{
					LeagueEntry{LeaguePoints: 100},
					LeagueEntry{LeaguePoints: 50},
					LeagueEntry{LeaguePoints: 70},
					LeagueEntry{LeaguePoints: 60},
				}},
			},
		}
	}

	league := service.ForSummoner(1)

	if len(league.Entries) != 4 {
		t.Error("Expected length 4 got: ", len(league.Entries))
	}

	if !sort.IsSorted(ByLeaguePoints(league.Entries)) {
		t.Error("League entries not sorted")
	}
}
