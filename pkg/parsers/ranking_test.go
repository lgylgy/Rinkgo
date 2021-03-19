package parsers

import (
	"testing"

	"github.com/lgylgy/rinkgo/pkg/api"
	"github.com/stretchr/testify/require"
)

func TestParseRanking(t *testing.T) {
	data := readFile(t, "ranking.txt")
	ranking, err := ParseRanking(data)
	if err != nil {
		t.Errorf("could not read test file ranking.txt: %s", err)
	}

	require.Len(t, ranking, 4)
	require.Equal(t, ranking,
		[]*api.Rank{
			{
				Team:         "PSG",
				Points:       9,
				Played:       3,
				Won:          3,
				Drawn:        0,
				Lost:         0,
				GoalsFor:     22,
				GoalsAgainst: 8,
			},
			{
				Team:         "OL",
				Points:       7,
				Played:       3,
				Won:          2,
				Drawn:        1,
				Lost:         0,
				GoalsFor:     14,
				GoalsAgainst: 8,
			},
			{
				Team:         "OM",
				Points:       6,
				Played:       3,
				Won:          2,
				Drawn:        0,
				Lost:         1,
				GoalsFor:     22,
				GoalsAgainst: 8,
			},
			{
				Team:         "LOSC",
				Points:       6,
				Played:       2,
				Won:          2,
				Drawn:        0,
				Lost:         0,
				GoalsFor:     17,
				GoalsAgainst: 5,
			},
		})
}
