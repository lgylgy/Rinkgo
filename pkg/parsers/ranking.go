package parsers

import (
	"golang.org/x/net/html"
	"strings"
)

type Rank struct {
	Club         string
	Points       string
	Played       string
	Won          string
	Drawn        string
	Lost         string
	GoalsFor     string
	GoalsAgainst string
}

const (
	club         = uint32(2)
	points       = uint32(3)
	played       = uint32(4)
	won          = uint32(5)
	drawn        = uint32(6)
	lost         = uint32(7)
	goalsFor     = uint32(9)
	goalsAgainst = uint32(10)
)

func extractRanking(value string) ([]Rank, error) {
	content, err := html.Parse(strings.NewReader(value))
	if err != nil {
		return nil, err
	}
	table := []Rank{}
	counter := uint32(0)

	// Extract rank informations
	var extractRankFields func(*html.Node, *Rank)
	extractRankFields = func(n *html.Node, row *Rank) {
		if n.Type == html.TextNode {
			switch counter {
			case club:
				row.Club = n.Data
			case points:
				row.Points = n.Data
			case played:
				row.Played = n.Data
			case won:
				row.Won = n.Data
			case drawn:
				row.Drawn = n.Data
			case lost:
				row.Lost = n.Data
			case goalsFor:
				row.GoalsFor = n.Data
			case goalsAgainst:
				row.GoalsAgainst = n.Data
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractRankFields(c, row)
		}
	}

	// Extract rank
	var extractRank func(*html.Node, *Rank)
	extractRank = func(n *html.Node, row *Rank) {
		if n.Type == html.ElementNode && n.Data == "td" {
			counter++
			extractRankFields(n, row)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractRank(c, row)
		}
	}

	// Retrieve table
	var extractTable func(*html.Node)
	extractTable = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "class" && (a.Val == "separate" ||
					a.Val == "odd separate" ||
					a.Val == "odd " ||
					a.Val == "") {
					row := Rank{}
					counter = uint32(0)
					extractRank(n, &row)
					table = append(table, row)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractTable(c)
		}
	}
	extractTable(content)
	return table, nil
}
