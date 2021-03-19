package parsers

import (
	"strings"

	"golang.org/x/net/html"
)

type Scorer struct {
	Name   string
	Club   string
	Matchs string
	Goals  string
}

const (
	nameField  = uint32(2)
	clubField  = uint32(3)
	matchField = uint32(4)
	goalField  = uint32(5)
)

func extractScorers(value string) ([]Scorer, error) {
	content, err := html.Parse(strings.NewReader(value))
	if err != nil {
		return nil, err
	}
	scorers := []Scorer{}
	counter := uint32(0)

	// Extract scorer informations
	var extractScorerFields func(*html.Node, *Scorer)
	extractScorerFields = func(n *html.Node, scorer *Scorer) {
		if n.Type == html.TextNode {
			switch counter {
			case nameField:
				scorer.Name = n.Data
			case clubField:
				scorer.Club = n.Data
			case matchField:
				scorer.Matchs = n.Data
			case goalField:
				scorer.Goals = n.Data
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractScorerFields(c, scorer)
		}
	}

	// Extract scorer struct
	var extractScorer func(*html.Node, *Scorer)
	extractScorer = func(n *html.Node, scorer *Scorer) {
		if n.Type == html.ElementNode && n.Data == "td" {
			counter++
			extractScorerFields(n, scorer)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractScorer(c, scorer)
		}
	}

	// Extract scorers struct
	var extractScorers func(*html.Node)
	extractScorers = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "class" && (a.Val == "separate" ||
					a.Val == "odd separate" ||
					a.Val == "odd " ||
					a.Val == "") {
					scorer := Scorer{}
					counter = uint32(0)
					extractScorer(n, &scorer)
					scorers = append(scorers, scorer)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractScorers(c)
		}
	}
	extractScorers(content)
	return scorers, nil
}
