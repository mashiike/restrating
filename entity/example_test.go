package entity_test

import (
	"fmt"
	"time"

	"github.com/mashiike/restrating/entity"
)

func ExamplePlayer() {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	baseTime := time.Date(2019, 05, 01, 0, 0, 0, 0, loc)
	config := entity.NewConfig()
	players := make(entity.RatingResourceCollection, 3)
	players = players.Add(
		entity.NewPlayer("sheep", baseTime, config),  //rrn:player:sheep
		entity.NewPlayer("goat", baseTime, config),   //rrn:player:goat
		entity.NewPlayer("donkey", baseTime, config), //rrn:player:donkey
	)
	scores := entity.ScoreCollection{
		entity.MustParseRNN("rrn:player:sheep"):  0.0,
		entity.MustParseRNN("rrn:player:goat"):   1.0,
		entity.MustParseRNN("rrn:player:donkey"): 2.0,
	}

	for k := 0; k < 5; k++ {
		players.Updates(scores, baseTime.AddDate(0, 0, k*8), config)
	}

	for _, player := range players.Ranking() {
		fmt.Println(player)
	}
	//Output:
	//rrn:player:donkey 1822.6p-273.5(10/0/0)
	//rrn:player:goat 1500.0p-299.6(5/5/0)
	//rrn:player:sheep 1177.4p-273.5(0/10/0)
}
