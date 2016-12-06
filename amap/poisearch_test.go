package amap

import (
	"fmt"
	"testing"
)

const (
	AMAP_KEY = "b3abf03fa1e83992727f0625a918fe73"
)

var (
	client = NewClient(AMAP_KEY)
)

func TestPoiSearch(t *testing.T) {
	keyword := "乐视大厦"
	city := "北京市"

	req := NewPoiSearchRequest(client, keyword).SetCity(city)
	resp, err := req.Do()
	if err != nil {
		t.Fatal(err)
	} else {
		for _, poi := range resp.Pois {
			fmt.Printf("%v\n", poi)
		}
	}
}
