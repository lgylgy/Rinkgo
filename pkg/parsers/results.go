package parsers

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/lgylgy/rinkgo/pkg/api"
)

func ParseResults(data string) ([]*api.Result, error) {
	results := []*api.Result{}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		return nil, err
	}

	doc.Find(".results-table").Each(func(_ int, div *goquery.Selection) {
		var date string

		div.Find("tbody tr").Each(func(_ int, tr *goquery.Selection) {
			if tr.HasClass("tr-date") {
				date = tr.Find("td").Text()
			} else {
				result := &api.Result{}
				tr.Find("td").Each(func(_ int, td *goquery.Selection) {
					if td.HasClass("equipe-name") {
						if result.HomeTeam != "" {
							result.OutTeam = td.Find("a").Text()
						} else {
							result.HomeTeam = td.Find("a").Text()
						}
					} else if td.HasClass("match-score") {
						td.Find("span").Each(func(_ int, span *goquery.Selection) {
							if result.Score != "" {
								result.Score += "-"
							}
							result.Score += strings.TrimSpace(span.Text())
						})
					}
				})
				result.Date = date
				results = append(results, result)
			}
		})
	})

	return results, nil
}
