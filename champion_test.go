package riotapi

import (
	"io"
	"testing"
)

func TestGetAllChampions(t *testing.T) {
	service := new(ChampionService)
	service.Service = NewFakeService()

	service.Parser.(*FakeParser).ParseFunc = func(response io.Reader, ret interface{}) {
		pointer := ret.(*map[string][]Champion)
		*pointer = map[string][]Champion{
			"champions": {
				Champion{},
				Champion{},
				Champion{},
				Champion{},
			},
		}
	}

	champions := service.All()

	if len(champions) != 4 {
		t.Error("Expected length 4 got: ", len(champions))
	}
}

func TestGetOneChampion(t *testing.T) {
	service := new(ChampionService)
	service.Service = NewFakeService()

	service.Parser.(*FakeParser).ParseFunc = func(response io.Reader, ret interface{}) {
		pointer := ret.(*Champion)
		*pointer = Champion{
			Id: 42,
		}
	}

	champion := service.One(42)

	if champion.Id != 42 {
		t.Error("Expected champion id 42 got: ", champion.Id)
	}
}
