package parsers

import (
	"testing"

	"github.com/lgylgy/rinkgo/pkg/api"
	"github.com/stretchr/testify/require"
)

func TestParsePlayerGoals(t *testing.T) {
	data := readFile(t, "player.txt")
	player, err := ParsePlayerGoals(data)
	require.NoError(t, err, "could not read test file player.txt")

	require.Equal(t, player.Name, "ZIDANE")
	require.Equal(t, player.History,
		[]*api.HistoryGoals{
			{
				Season: "2020/2021",
				Team:   "PSG",
				Event:  "Ligue 1",
				Games:  1,
				Goals:  1,
			},
			{
				Season: "2019/2020",
				Team:   "PSG",
				Event:  "Ligue 1",
				Games:  1,
				Goals:  0,
			},
			{
				Season: "2018/2019",
				Team:   "LORIENT",
				Event:  "Ligue 2",
				Games:  8,
				Goals:  5,
			},
			{
				Season: "2017/2018",
				Team:   "LORIENT",
				Event:  "Ligue 2",
				Games:  1,
				Goals:  1,
			},
			{
				Season: "2017/2018",
				Team:   "LORIENT",
				Event:  "Coupe de France",
				Games:  8,
				Goals:  2,
			},
		})
}
