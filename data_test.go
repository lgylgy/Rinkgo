package data

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestExtractScorers(t *testing.T) {

	filename := "datatest/scorers.txt"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal("Failed to read file: " + filename + " (" + err.Error() + ")")
	}

	scorers, err := extractScorers(string(data))
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

func TestExtractMatch(t *testing.T) {

	filename := "datatest/match.txt"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal("Failed to read file: " + filename + " (" + err.Error() + ")")
	}

	scorers, err := extractMatchs(string(data))
	if err != nil {
		t.Fatalf("could not read test file match.txt: %s", err)
	}

	expected := []Match{
		Match{
			"02/12/2017",
			"PSG",
			"6 - 1",
			"MARSEILLE",
		},
		Match{
			"02/12/2017",
			"EAG",
			"5 - 4",
			"BREST",
		},
		Match{
			"02/12/2017",
			"NANTES",
			"0 - 11",
			"LILLE",
		},
	}
	if !reflect.DeepEqual(scorers, expected) {
		t.Fatalf("scorers informations differ: %+v\n!=\n%+v", expected, scorers)
	}
}
