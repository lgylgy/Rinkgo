package parsers

import (
	"testing"

	"github.com/lgylgy/rinkgo/pkg/api"
	"github.com/stretchr/testify/require"
)

func TestParsePlayer(t *testing.T) {
	data := readFile(t, "player.txt")
	player, err := ParsePlayer(data)
	if err != nil {
		t.Errorf("could not read test file player.txt: %s", err)
	}
	require.Equal(t, player.Name, "ZIDANE")
	require.Equal(t, player.History,
		[]*api.Entry{
			{
				Season: "2020/2021",
				Team:   "PSG",
				Event:  "Ligue 1",
				Matchs: 1,
				Goals:  1,
			},
			{
				Season: "2019/2020",
				Team:   "PSG",
				Event:  "Ligue 1",
				Matchs: 1,
				Goals:  0,
			},
			{
				Season: "2018/2019",
				Team:   "LORIENT",
				Event:  "Ligue 2",
				Matchs: 8,
				Goals:  5,
			},
			{
				Season: "2017/2018",
				Team:   "LORIENT",
				Event:  "Ligue 2",
				Matchs: 1,
				Goals:  1,
			},
			{
				Season: "2017/2018",
				Team:   "LORIENT",
				Event:  "Coupe de France",
				Matchs: 8,
				Goals:  2,
			},
		})
}
