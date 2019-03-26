package parsers

import (
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"
)

func readFile(t *testing.T, filename string) string {
	path := filepath.Join("datatest", filename)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal("Failed to read file: " + filename + " (" + err.Error() + ")")
	}
	return string(data)
}

func TestExtractScorers(t *testing.T) {

	data := readFile(t, "scorers.txt")
	scorers, err := extractScorers(data)
	if err != nil {
		t.Fatalf("could not read test file scorers.txt: %s", err)
	}

	expected := []Scorer{
		Scorer{
			"PLAYER A",
			"CLUB 1",
			"22",
			"49",
		},
		Scorer{
			"PLAYER B",
			"CLUB 2",
			"22",
			"30",
		},
		Scorer{
			"PLAYER C",
			"CLUB 2",
			"21",
			"30",
		},
		Scorer{
			"PLAYER D",
			"CLUB 4",
			"22",
			"27",
		},
	}
	if !reflect.DeepEqual(scorers, expected) {
		t.Fatalf("scorers informations differ: %+v\n!=\n%+v", expected, scorers)
	}
}

func TestExtractFixtures(t *testing.T) {

	data := readFile(t, "fixtures.txt")
	fixtures, err := ExtractFixtures(data)
	if err != nil {
		t.Fatalf("could not read test file fixtures.txt: %s", err)
	}

	expected := []Fixture{
		Fixture{
			"02/12/2017",
			"PSG",
			"6 - 1",
			"MARSEILLE",
		},
		Fixture{
			"02/12/2017",
			"EAG",
			"5 - 4",
			"BREST",
		},
		Fixture{
			"02/12/2017",
			"NANTES",
			"0 - 11",
			"LILLE",
		},
	}
	if !reflect.DeepEqual(fixtures, expected) {
		t.Fatalf("fixtures informations differ: %+v\n!=\n%+v", expected, fixtures)
	}
}

func TestExtractRanking(t *testing.T) {

	data := readFile(t, "ranking.txt")
	ranking, err := extractRanking(data)
	if err != nil {
		t.Fatalf("could not read test file ranking.txt: %s", err)
	}

	expected := []Rank{
		Rank{
			"PSG",
			"55",
			"22",
			"17",
			"4",
			"1",
			"120",
			"44",
		},
		Rank{
			"EAG",
			"53",
			"22",
			"16",
			"5",
			"1",
			"107",
			"36",
		},
		Rank{
			"NANTES",
			"49",
			"22",
			"16",
			"1",
			"5",
			"103",
			"62",
		},
	}
	if !reflect.DeepEqual(ranking, expected) {
		t.Fatalf("scorers informations differ: %+v\n!=\n%+v", expected, ranking)
	}
}
