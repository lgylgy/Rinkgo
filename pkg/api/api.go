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
