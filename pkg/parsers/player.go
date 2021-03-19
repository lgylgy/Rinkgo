package parsers

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/lgylgy/rinkgo/pkg/api"
)

func readFile(t *testing.T, filename string) string {
	path := filepath.Join("datatest", filename)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		t.Error("Failed to read file: " + filename + " (" + err.Error() + ")")
	}
	return string(data)
}

func ParsePlayer(data string) (*api.Player, error) {
	player := &api.Player{
		History: []*api.Entry{},
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
			entry := &api.Entry{}
			tr.Find("td").Each(func(ix int, td *goquery.Selection) {
				switch ix {
				case 0:
					entry.Season = td.Text()
				case 1:
					entry.Team = td.Find("a").Text()
				case 2:
					entry.Event = td.Text()
				case 3:
					entry.Matchs = convertToInteger(td)
				case 4:
					entry.Goals = convertToInteger(td)
				}
			})
			player.History = append(player.History, entry)
		})
	})
	return player, nil
}
