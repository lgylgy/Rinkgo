package parsers

import (
	"testing"

	"github.com/lgylgy/rinkgo/pkg/api"
	"github.com/stretchr/testify/require"
)

func TestParsePlayerGoalsPerGames(t *testing.T) {
	data := readFile(t, "player_games.txt")
	player, err := ParsePlayerGoalsPerGames(data)
	require.NoError(t, err, "could not read test file player_games.txt")

	require.Equal(t, player.History,
		[]*api.HistoryGames{
			{
				Event: "Ligue 1",
				Game:  "PSG 3 - 8 EAG",
			},
			{
				Event: "Ligue 1",
				Game:  "SB29 7 - 3 PSG",
			},
			{
				Event: "Coupe de france",
				Game:  "PSG 6 - 5 LORIENT",
			},
		})
}
