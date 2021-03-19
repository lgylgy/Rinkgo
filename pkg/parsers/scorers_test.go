package parsers

import (
	"testing"

	"github.com/lgylgy/rinkgo/pkg/api"
	"github.com/stretchr/testify/require"
)

func TestParseScorers(t *testing.T) {
	data := readFile(t, "scorers.txt")
	scorers, err := ParseScorers(data)
	if err != nil {
		t.Errorf("could not read test file scorers.txt: %s", err)
	}

	require.Len(t, scorers, 3)
	require.Equal(t, scorers,
		[]*api.Scorer{
			{
				Name:   "NEYMAR",
				Team:   "PSG",
				Matchs: 3,
				Goals:  7,
			},
			{
				Name:   "DEPAY",
				Team:   "OL",
				Matchs: 2,
				Goals:  5,
			},
			{
				Name:   "MILIK",
				Team:   "OM",
				Matchs: 2,
				Goals:  5,
			},
		})
}
