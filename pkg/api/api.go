package api

type Entry struct {
	Season string
	Team   string
	Event  string
	Matchs uint32
	Goals  uint32
}

type Player struct {
	Name    string
	History []*Entry
}

type Scorer struct {
	Name   string
	Team   string
	Matchs uint32
	Goals  uint32
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
