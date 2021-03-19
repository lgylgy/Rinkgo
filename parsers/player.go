package parsers

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

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

func ParsePlayer(data string) (*Player, error) {
	player := &Player{
		History: []*Entry{},
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		return nil, err
	}

	doc.Find(".info-competiteur").Each(func(_ int, div *goquery.Selection) {
		div.Find("span").Each(func(i int, span *goquery.Selection) {
			switch i {
			case 0:
				player.Name = span.Text()
			}
		})
	})

	doc.Find(".competitor-stats-tab").Each(func(_ int, div *goquery.Selection) {
		div.Find("tbody tr").Each(func(_ int, tr *goquery.Selection) {
			entry := &Entry{}
			tr.Find("td").Each(func(ix int, td *goquery.Selection) {
				switch ix {
				case 0:
					entry.Season = td.Text()
				case 1:
					entry.Team = td.Find("a").Text()
				case 2:
					entry.Event = td.Text()
				case 3:
					i, err := strconv.Atoi(td.Text())
					if err == nil {
						entry.Matchs = uint32(i)
					}
				case 4:
					i, err := strconv.Atoi(td.Text())
					if err == nil {
						entry.Goals = uint32(i)
					}
				}
			})
			player.History = append(player.History, entry)
		})
	})
	return player, nil
}