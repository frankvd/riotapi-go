package riotapi

import (
	"io"
	"testing"
)

func TestGetMatchHistory(t *testing.T) {
	service := new(MatchHistoryService)
	service.Service = NewFakeService()

	service.Parser.(*FakeParser).ParseFunc = func(response io.Reader, ret interface{}) {
		pointer := ret.(*MatchHistory)
		*pointer = MatchHistory{
			Matches: []Match{
				Match{},
				Match{},
				Match{},
				Match{},
			},
		}
	}

	history := service.ForSummoner(1)

	if len(history.Matches) != 4 {
		t.Error("Expected length 4 got: ", len(history.Matches))
	}
}
