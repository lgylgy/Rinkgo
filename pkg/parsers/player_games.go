package parsers

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/lgylgy/rinkgo/pkg/api"
)

func ParsePlayerGoalsPerGames(data string) (*api.PlayerGoalsPerGames, error) {
	player := &api.PlayerGoalsPerGames{
		History: []*api.HistoryGames{},
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		return nil, err
	}

	doc.Find(".competitor-matchs-tab").Each(func(_ int, div *goquery.Selection) {
		div.Find("tbody tr").Each(func(_ int, tr *goquery.Selection) {
			entry := &api.HistoryGames{}
			tr.Find("td").Each(func(ix int, td *goquery.Selection) {
				switch ix {
				case 2:
					entry.Event = td.Text()
				case 3:
					entry.Game = td.Text()
				}
			})
			player.History = append(player.History, entry)
		})
	})
	return player, nil
}
