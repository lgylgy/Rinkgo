package parsers

import (
	"testing"

	"github.com/lgylgy/rinkgo/pkg/api"
	"github.com/stretchr/testify/require"
)

func TestParseResults(t *testing.T) {
	data := readFile(t, "results.txt")
	results, err := ParseResults(data)
	require.NoError(t, err, "could not read test file results.txt")

	require.Len(t, results, 3)
	require.Equal(t, results,
		[]*api.Result{
			{
				Date:     "24/12/2021",
				HomeTeam: "PSG",
				Score:    "4-0",
				OutTeam:  "OM",
			},
			{
				Date:     "24/12/2021",
				HomeTeam: "EAG",
				Score:    "3-2",
				OutTeam:  "SB29",
			},
			{
				Date:     "23/01/2022",
				HomeTeam: "MONACO",
				Score:    "2-4",
				OutTeam:  "RENNES",
			},
		})
}
