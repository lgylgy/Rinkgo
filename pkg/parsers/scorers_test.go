package parsers

import (
	"testing"

	"github.com/lgylgy/rinkgo/pkg/api"
	"github.com/stretchr/testify/require"
)

func TestParseScorers(t *testing.T) {
	data := readFile(t, "scorers.txt")
	scorers, err := ParseScorers(data)
	require.NoError(t, err, "could not read test file scorers.txt")

	require.Len(t, scorers, 3)
	require.Equal(t, scorers,
		[]*api.Scorer{
			{
				Name:  "NEYMAR",
				Team:  "PSG",
				Games: 3,
				Goals: 7,
			},
			{
				Name:  "DEPAY",
				Team:  "OL",
				Games: 2,
				Goals: 5,
			},
			{
				Name:  "MILIK",
				Team:  "OM",
				Games: 2,
				Goals: 5,
			},
		})
}
