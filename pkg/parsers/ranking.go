package parsers

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/lgylgy/rinkgo/pkg/api"
)

func convertToInteger(td *goquery.Selection) uint32 {
	i, err := strconv.Atoi(td.Text())
	if err != nil {
		return 0
	}
	return uint32(i)
}

func ParseRanking(data string) ([]*api.Rank, error) {
	ranking := []*api.Rank{}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		return nil, err
	}

	doc.Find(".ranking-table").Each(func(_ int, div *goquery.Selection) {
		div.Find("tbody tr").Each(func(_ int, tr *goquery.Selection) {
			rank := &api.Rank{}
			tr.Find("td").Each(func(ix int, td *goquery.Selection) {
				switch ix {
				case 1:
					rank.Team = td.Find("a").Text()
				case 2:
					rank.Points = convertToInteger(td)
				case 3:
					rank.Played = convertToInteger(td)
				case 4:
					rank.Won = convertToInteger(td)
				case 7:
					rank.Drawn = convertToInteger(td)
				case 8:
					rank.Lost = convertToInteger(td)
				case 12:
					rank.GoalsFor = convertToInteger(td)
				case 13:
					rank.GoalsAgainst = convertToInteger(td)
				}
			})
			ranking = append(ranking, rank)
		})
	})

	return ranking, nil
}
