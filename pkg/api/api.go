package api

type HistoryGoals struct {
	Season string
	Team   string
	Event  string
	Games  uint32
	Goals  uint32
}

type PlayerGoals struct {
	Name    string
	History []*HistoryGoals
}

type Scorer struct {
	Name  string
	Team  string
	Games uint32
	Goals uint32
}

type Rank struct {
	Team         string
	Points       uint32
	Played       uint32
	Won          uint32
	Drawn        uint32
	Lost         uint32
	GoalsFor     uint32
	GoalsAgainst uint32
}

type Result struct {
	Date     string
	HomeTeam string
	Score    string
	OutTeam  string
}

type HistoryGames struct {
	Event string
	Game  string
}

type PlayerGoalsPerGames struct {
	Name    string
	History []*HistoryGames
}
