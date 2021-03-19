package pstat

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lgylgy/rinkgo/pkg/api"
	"github.com/lgylgy/rinkgo/pkg/parsers"
	pb "github.com/lgylgy/rinkgo/pkg/services/pstat/proto"
)

type pStatServer struct {
	pb.UnimplementedPStatServiceServer
	players map[int32]*api.Player
	urlDB   string
}

func NewPStatServer(urlDB string) *pStatServer {
	return &pStatServer{
		players: make(map[int32]*api.Player),
		urlDB:   urlDB,
	}
}

func convertToProto(id int32, player *api.Player) *pb.History {
	entries := []*pb.Entry{}
	for _, v := range player.History {
		entries = append(entries, &pb.Entry{
			Season: v.Season,
			Team:   v.Team,
			Event:  v.Event,
			Matchs: int32(v.Matchs),
			Goals:  int32(v.Goals),
		})
	}
	return &pb.History{
		PlayerID: id,
		Name:     player.Name,
		Entries:  entries,
	}
}

func (ps *pStatServer) downloadHistory(id int32) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%v/%v", ps.urlDB, id))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (ps *pStatServer) GetHistory(ctx context.Context, request *pb.Request) (*pb.History, error) {
	id := request.GetPlayerID()
	p, ok := ps.players[id]
	if ok {
		return convertToProto(id, p), nil
	}
	data, err := ps.downloadHistory(id)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve history information")
	}
	p, err = parsers.ParsePlayer(data)
	if err != nil {
		return nil, fmt.Errorf("unable to parse history information")
	}
	ps.players[id] = p
	return convertToProto(id, p), nil
}
