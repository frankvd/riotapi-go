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
	t.Log(champions)
	if len(champions) != 4 {
		t.Error("Expected length 4 got: ", len(champions))
	}
}
