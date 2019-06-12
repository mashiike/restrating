package entity

import (
	"errors"
	"strings"
)

const (
	rrnPrefix    = "rrn:"
	rrnDelimiter = ":"
)

//RRN is rating resource name
type RRN struct {
	Type string
	Name string
}

//MustParseRNN はexampleをシンプルに書くよう。errorが起きたらpanicする
func MustParseRNN(rrnStr string) RRN {
	rrn, err := ParseRRN(rrnStr)
	if err != nil {
		panic(err)
	}
	return rrn
}

//ParseRRN はRRNを読み取ります
func ParseRRN(rrn string) (RRN, error) {
	if !strings.HasPrefix(rrn, rrnPrefix) {
		return RRN{}, errors.New("rrn: invalid prefix")
	}
	sections := strings.SplitN(rrn, rrnDelimiter, 3)
	if len(sections) != 3 {
		return RRN{}, errors.New("rrn: not enough sections")
	}
	return RRN{
		Type: sections[1],
		Name: sections[2],
	}, nil
}

//String はRRNの形式フォーマットを返します。
//
//examples:
//  rrn:player:abdea3dfa
//  rrn:team:dda32341
func (rrn RRN) String() string {
	return rrnPrefix + rrn.Type + rrnDelimiter + rrn.Name
}
