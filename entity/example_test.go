package entity_test

import (
	"fmt"
	"sort"
	"time"

	"github.com/mashiike/rating"
	"github.com/mashiike/restrating/entity"
)

func ExamplePlayer() {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	baseTime := time.Date(2019, 05, 01, 0, 0, 0, 0, loc)
	config := entity.NewConfig()
	players := []*entity.Player{
		entity.NewPlayer(1, baseTime, config),
		entity.NewPlayer(2, baseTime, config),
		entity.NewPlayer(3, baseTime, config),
	}

	for k := 0; k < 5; k++ {
		for i := 0; i < len(players); i++ {
			for j := 0; j < len(players); j++ {
				players[j].Prepare(baseTime.AddDate(0, 0, k*8), config)
			}
			for j := i + 1; j < len(players); j++ {
				players[i].Update(&entity.MatchResult{
					Opponent:  players[j].Rating(),
					Score:     rating.ScoreLose,
					OutcomeAt: baseTime.AddDate(0, 0, k*8+i),
				})
				players[j].Update(&entity.MatchResult{
					Opponent:  players[i].Rating(),
					Score:     rating.ScoreWin,
					OutcomeAt: baseTime.AddDate(0, 0, k*8+i),
				})
			}
		}
	}

	sort.Slice(players, func(i, j int) bool { return players[i].Rating().Strength() > players[j].Rating().Strength() })
	for _, player := range players {
		fmt.Println(player)
	}
	//Output:
	//id:3 1866.4p-365.9(10/0/0)
	//id:2 1463.4p-333.5(5/5/0)
	//id:1 1060.6p-370.9(0/10/0)
}
