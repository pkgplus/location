package amap

import (
	"fmt"
	"testing"
)

func TestInputtips(t *testing.T) {
	keyword := "公交车站"
	loc := "116.498764,39.923069"
	req := NewInputtipsRequest(client, keyword).
		SetCityLimit(true).
		SetLocation(loc)
	resp, err := req.Do()
	if err != nil {
		t.Fatal(err)
	} else {
		for _, tip := range resp.Tips {
			fmt.Printf("%v\n", tip)
		}
	}
}
