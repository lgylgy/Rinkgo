package parsers

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/lgylgy/rinkgo/pkg/api"
)

func ParseScorers(data string) ([]*api.Scorer, error) {
	scorers := []*api.Scorer{}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		return nil, err
	}

	doc.Find(".scorers-table").Each(func(_ int, div *goquery.Selection) {
		div.Find("tbody tr").Each(func(_ int, tr *goquery.Selection) {
			scorer := &api.Scorer{}
			tr.Find("td").Each(func(ix int, td *goquery.Selection) {
				switch ix {
				case 1:
					scorer.Name = td.Find("a").Text()
				case 2:
					scorer.Team = td.Find("a").Text()
				case 4:
					scorer.Goals = convertToInteger(td)
				case 6:
					scorer.Matchs = convertToInteger(td)
				}
			})
			scorers = append(scorers, scorer)
		})
	})

	return scorers, nil
}
