package riotapi

import (
	"io"
	"testing"
)

func TestGetSummoner(t *testing.T) {
	service := new(SummonerService)
	service.Service = NewFakeService()

	service.Parser.(*FakeParser).ParseFunc = func(response io.Reader, ret interface{}) {
		pointer := ret.(*map[string]Summoner)
		*pointer = map[string]Summoner{
			"henk": Summoner{Name: "PASSED"},
		}
	}

	summoner := service.ByName("henk")

	if summoner.Name != "PASSED" {
		t.Error("Expected name PASSED got: ", summoner.Name)
	}
}

func TestGetSummonerRunes(t *testing.T) {
	service := new(SummonerService)
	service.Service = NewFakeService()

	service.Parser.(*FakeParser).ParseFunc = func(response io.Reader, ret interface{}) {
		pointer := ret.(*map[string]Runes)
		*pointer = map[string]Runes{
			"42": Runes{SummonerId: 42},
		}
	}

	runes := service.Runes(42)

	if runes.SummonerId != 42 {
		t.Error("Expected id 42 got: ", runes.SummonerId)
	}
}

func TestGetSummonerMasteries(t *testing.T) {
	service := new(SummonerService)
	service.Service = NewFakeService()

	service.Parser.(*FakeParser).ParseFunc = func(response io.Reader, ret interface{}) {
		pointer := ret.(*map[string]Masteries)
		*pointer = map[string]Masteries{
			"42": Masteries{SummonerId: 42},
		}
	}

	masteries := service.Masteries(42)

	if masteries.SummonerId != 42 {
		t.Error("Expected id 42 got: ", masteries.SummonerId)
	}
}
