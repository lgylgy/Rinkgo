package parsers

import (
	"golang.org/x/net/html"
	"strings"
)

const (
	home  = uint32(1)
	score = uint32(2)
	out   = uint32(3)
)

type Fixture struct {
	Date     string
	HomeTeam string
	Score    string
	OutTeam  string
}

func ExtractFixtures(value string) ([]Fixture, error) {
	content, err := html.Parse(strings.NewReader(value))
	if err != nil {
		return nil, err
	}

	fixtures := []Fixture{}
	counter := uint32(0)
	date := ""

	// Extract date
	var extractDate func(*html.Node)
	extractDate = func(n *html.Node) {
		data := strings.TrimSpace(n.Data)
		if n.Type == html.TextNode && len(data) > 0 {
			date = data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractDate(c)
		}
	}

	// Extract fixture informations
	var extractData func(*html.Node, *Fixture)
	extractData = func(n *html.Node, fixture *Fixture) {
		data := strings.TrimSpace(n.Data)
		if n.Type == html.TextNode && len(data) > 0 {
			switch counter {
			case home:
				fixture.HomeTeam = data
			case score:
				fixture.Score = data
			case out:
				fixture.OutTeam = data
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractData(c, fixture)
		}
	}

	// Extract fixture
	var extractfixture func(*html.Node, *Fixture)
	extractfixture = func(n *html.Node, fixture *Fixture) {
		if n.Type == html.ElementNode && n.Data == "td" {
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "right" {
					counter++
					extractData(n, fixture)
				}
				if a.Key == "class" && a.Val == "center score" {
					counter++
					extractData(n, fixture)
				}
				if a.Key == "class" && a.Val == "left" {
					counter++
					extractData(n, fixture)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractfixture(c, fixture)
		}
	}

	// Extract fixtures
	var extractFixtures func(*html.Node)
	extractFixtures = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "caption" {
			extractDate(n)
		}
		if n.Type == html.ElementNode && n.Data == "tr" {
			fixture := Fixture{}
			fixture.Date = date
			counter = uint32(0)
			extractfixture(n, &fixture)
			if len(fixture.HomeTeam) > 0 && len(fixture.Score) > 0 && len(fixture.OutTeam) > 0 {
				fixtures = append(fixtures, fixture)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractFixtures(c)
		}
	}
	extractFixtures(content)
	return fixtures, nil
}
