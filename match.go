package riotapi

import (
	"fmt"
	"net/url"
	"strconv"
)

// Match history service
type MatchHistoryService struct {
	*Service
}

// Match service
type MatchService struct {
	*Service
}

// Returns the match history of the summoner
func (service *MatchHistoryService) ForSummoner(id int) MatchHistory {
	textId := strconv.Itoa(id)
	resp := service.Client.Call("match-history", []string{textId}, nil)
	history := MatchHistory{}
	service.Parser.Parse(resp.Body, &history)

	return history
}

func (service *MatchService) Get(id int, timeline bool) Match {
	textId := strconv.Itoa(id)
	fmt.Printf("%s", textId)
	query := url.Values{}
	if timeline {
		query.Add("includeTimeline", "true")
	}
	resp := service.Client.Call("match", []string{textId}, query)
	match := Match{}
	service.Parser.Parse(resp.Body, &match)

	return match
}

// MatchHistory
type MatchHistory struct {
	Matches []Match
}

// Match
type Match struct {
	MapId                 int
	MatchCreation         int
	MatchDuration         int
	MatchId               int
	MatchMode             string
	MatchType             string
	MatchVersion          string
	ParticipantIdentities []ParticipantIdentity
	Participants          []Participant
	PlatformId            string
	QueueType             string
	Region                string
	Season                string
	Teams                 []Team
	Timeline              MatchTimeline
}

type Team struct {
	Bans                 []BannedChampion
	BaronKills           int
	DominionVictoryScore int
	DragonKills          int
	FirstBaron           bool
	FirstBlood           bool
	FirstDragon          bool
	FirstInhibitor       bool
	FirstTower           bool
	InhibitorKills       int
	TeamId               int
	TowerKills           int
	VilemawKills         int
	Winner               bool
}

type BannedChampion struct {
	ChampionId int
	PickTurn   int
}

// Participant
type Participant struct {
	ChampionId                int
	HighestAchievedSeasonTier string
	Masteries                 []Mastery
	ParticipantId             int
	Runes                     []Rune
	Spell1Id                  int
	Spell2Id                  int
	Stats                     ParticipantStats
	TeamId                    int
	Timeline                  ParticipantTimeline
}

// Participant identity
type ParticipantIdentity struct {
	ParticipantId int
	Player        Player
}

// Participant stats
type ParticipantStats struct {
	Assists                         int
	ChampLevel                      int
	CombatPlayerScore               int
	Deaths                          int
	DoubleKills                     int
	FirstBloodAssist                bool
	FirstBloodKill                  bool
	FirstInhibitorAssist            bool
	FirstInhibitorKill              bool
	FirstTowerAssist                bool
	FirstTowerKill                  bool
	GoldEarned                      int
	GoldSpent                       int
	InhibitorKills                  int
	Item0                           int
	Item1                           int
	Item2                           int
	Item3                           int
	Item4                           int
	Item5                           int
	Item6                           int
	KillingSprees                   int
	Kills                           int
	LargestCriticalStrike           int
	LargestKillingSpree             int
	LargestMultiKill                int
	MagicDamageDealt                int
	MagicDamageDealtToChampions     int
	MagicDamageTaken                int
	MinionsKilled                   int
	NeutralMinionsKilled            int
	NeutralMinionsKilledEnemyJungle int
	NeutralMinionsKilledTeamJungle  int
	NodeCapture                     int
	NodeCaptureAssist               int
	NodeNeutralize                  int
	NodeNeutralizeAssist            int
	ObjectivePlayerScore            int
	PentaKills                      int
	PhysicalDamageDealt             int
	PhysicalDamageDealtToChampions  int
	PhysicalDamageTaken             int
	QuadraKills                     int
	SightWardsBoughtInGame          int
	TeamObjective                   int
	TotalDamageDealt                int
	TotalDamageDealtToChampions     int
	TotalDamageTaken                int
	TotalHeal                       int
	TotalPlayerScore                int
	TotalScoreRank                  int
	TotalTimeCrowdControlDealt      int
	TotalUnitsHealed                int
	TowerKills                      int
	TripleKills                     int
	TrueDamageDealt                 int
	TrueDamageDealtToChampions      int
	TrueDamageTaken                 int
	UnrealKills                     int
	VisionWardsBoughtInGame         int
	WardsKilled                     int
	WardsPlaced                     int
	Winner                          bool
}

// Participant timeline
type ParticipantTimeline struct {
	AncientGolemAssistsPerMinCounts ParticipantTimelineData
	AncientGolemKillsPerMinCounts   ParticipantTimelineData
	AssistedLaneDeathsPerMinDeltas  ParticipantTimelineData
	AssistedLaneKillsPerMinDeltas   ParticipantTimelineData
	BaronAssistsPerMinCounts        ParticipantTimelineData
	BaronKillsPerMinCounts          ParticipantTimelineData
	CreepsPerMinDeltas              ParticipantTimelineData
	CsDiffPerMinDeltas              ParticipantTimelineData
	DamageTakenDiffPerMinDeltas     ParticipantTimelineData
	DamageTakenPerMinDeltas         ParticipantTimelineData
	DragonAssistsPerMinCounts       ParticipantTimelineData
	DragonKillsPerMinCounts         ParticipantTimelineData
	ElderLizardAssistsPerMinCounts  ParticipantTimelineData
	ElderLizardKillsPerMinCounts    ParticipantTimelineData
	GoldPerMinDeltas                ParticipantTimelineData
	InhibitorAssistsPerMinCounts    ParticipantTimelineData
	InhibitorKillsPerMinCounts      ParticipantTimelineData
	Lane                            string
	Role                            string
	TowerAssistsPerMinCounts        ParticipantTimelineData
	TowerKillsPerMinCounts          ParticipantTimelineData
	TowerKillsPerMinDeltas          ParticipantTimelineData
	VilemawAssistsPerMinCounts      ParticipantTimelineData
	VilemawKillsPerMinCounts        ParticipantTimelineData
	WardsPerMinDeltas               ParticipantTimelineData
	XpDiffPerMinDeltas              ParticipantTimelineData
	XpPerMinDeltas                  ParticipantTimelineData
}

// Participant timeline data
type ParticipantTimelineData struct {
	ZeroToTen      float32
	TenToTwenty    float32
	TwentyToThirty float32
	ThirtyToEnd    float32
}

// Player
type Player struct {
	MatchHistoryUri string
	ProfileIcon     int
	SummonerId      int
	SummonerName    string
}

type MatchTimeline struct {
	FrameInterval int
	Frames        []Frame
}

type Frame struct {
	Events            []Event
	ParticipantFrames map[string]ParticipantFrame
	Timestamp         int
}

type Event struct {
	AscendedType            string
	AssistingParticipantIds []int
	BuildingType            string
	CreatorId               int
	EventType               string
	ItemAfter               int
	ItemBefore              int
	ItemId                  int
	KillerId                int
	LaneType                string
	LevelUpType             string
	MonsterType             string
	ParticipantId           int
	PointCaptured           string
	Position                Position
	SkillSlot               int
	TeamId                  int
	Timestamp               int
	TowerType               string
	VictimId                int
	WardType                string
}

type Position struct {
	X int
	Y int
}

type ParticipantFrame struct {
	CurrentGold         int
	DominionScore       int
	JungleMinionsKilled int
	Level               int
	MinionsKilled       int
	ParticipantId       int
	Position            Position
	TeamScore           int
	TotalGold           int
	Xp                  int
}
