package data

import (
	"golang.org/x/net/html"
	"strings"
)

const (
	home  = uint32(1)
	score = uint32(2)
	out   = uint32(3)
)

type Match struct {
	Date     string
	HomeTeam string
	Score    string
	OutTeam  string
}

func extractMatchs(value string) ([]Match, error) {
	content, err := html.Parse(strings.NewReader(value))
	if err != nil {
		return nil, err
	}

	matchs := []Match{}
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

	// Extract match informations
	var extractData func(*html.Node, *Match)
	extractData = func(n *html.Node, match *Match) {
		data := strings.TrimSpace(n.Data)
		if n.Type == html.TextNode && len(data) > 0 {
			switch counter {
			case home:
				match.HomeTeam = data
			case score:
				match.Score = data
			case out:
				match.OutTeam = data
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractData(c, match)
		}
	}

	// Extract match
	var extractMatch func(*html.Node, *Match)
	extractMatch = func(n *html.Node, match *Match) {
		if n.Type == html.ElementNode && n.Data == "td" {
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "right" {
					counter++
					extractData(n, match)
				}
				if a.Key == "class" && a.Val == "center score" {
					counter++
					extractData(n, match)
				}
				if a.Key == "class" && a.Val == "left" {
					counter++
					extractData(n, match)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractMatch(c, match)
		}
	}

	// Extract matchs
	var extractMatchs func(*html.Node)
	extractMatchs = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "caption" {
			extractDate(n)
		}
		if n.Type == html.ElementNode && n.Data == "tr" {
			match := Match{}
			match.Date = date
			counter = uint32(0)
			extractMatch(n, &match)
			if len(match.HomeTeam) > 0 && len(match.Score) > 0 && len(match.OutTeam) > 0 {
				matchs = append(matchs, match)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractMatchs(c)
		}
	}
	extractMatchs(content)
	return matchs, nil
}
