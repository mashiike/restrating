package entity

import (
	"strconv"

	"github.com/mashiike/rating"
)

//GameCount is (win/lose/draw) count
type GameCount struct {
	Win  uint64 `json:"win"`
	Lose uint64 `json:"lose"`
	Draw uint64 `json:"draw"`
}

func (g GameCount) byteLength() int {
	return len("(win/lose/draw)")
}

func (g GameCount) appendByte(b []byte) []byte {
	b = append(b, '(')
	b = append(b, strconv.FormatUint(g.Win, 10)...)
	b = append(b, '/')
	b = append(b, strconv.FormatUint(g.Lose, 10)...)
	b = append(b, '/')
	b = append(b, strconv.FormatUint(g.Draw, 10)...)
	b = append(b, ')')
	return b
}

func (g GameCount) Total() uint64 {
	return g.Win + g.Lose + g.Draw
}

//String is implemets fmt.Stringer
func (g GameCount) String() string {
	b := make([]byte, 0, g.byteLength())
	b = g.appendByte(b)
	return string(b)
}

func (g GameCount) AddScore(score float64) GameCount {
	switch {
	case score == rating.ScoreDraw:
		g.Draw++
	case score > rating.ScoreDraw:
		g.Win++
	case score < rating.ScoreDraw:
		g.Lose++
	}
	return g
}
