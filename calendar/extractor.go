package main

import (
	"fmt"
	"github.com/lgylgy/rinkgo/parsers"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func extractDom(address string, count int) ([]string, error) {
	result := make([]string, count)
	client := http.Client{}

	for i := 1; i <= count; i++ {
		param := url.Values{}
		param.Set("numero", strconv.Itoa(i))
		fmt.Printf("%v/%v\n", i, count)

		req, err := http.NewRequest("POST", address, strings.NewReader(param.Encode()))
		if err != nil {
			return result, err
		}

		rpy, err := client.Do(req)
		if err != nil {
			return result, err
		}
		defer rpy.Body.Close()

		if rpy.StatusCode != http.StatusOK {
			return result, fmt.Errorf("Unable to retrieve fixture '%v': %v", i, rpy.Status)
		}

		data, err := ioutil.ReadAll(rpy.Body)
		if err != nil {
			return result, err
		}
		result[i-1] = string(data)
	}
	return result, nil
}

func extractFixtures(team string, dom []string) ([]parsers.Fixture, error) {
	result := []parsers.Fixture{}

	for _, v := range dom {
		fixtures, err := parsers.ExtractFixtures(v)
		if err != nil {
			return result, err
		}

		for _, t := range fixtures {
			if t.HomeTeam == team || t.OutTeam == team {
				result = append(result, t)
			}
		}
	}
	return result, nil
}
