package riotapi

type Summoner struct {
	Name          string `json:name`
	ID            int
	SummonerLevel int
	ProfileIconID int
	RevisionDate  int
}
