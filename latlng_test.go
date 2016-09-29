package location

import (
	"testing"
)

func TestGetLatLng(t *testing.T) {
	province := "香港"
	city := "香港"

	loc := GetLatLng(province, city)
	if loc != "22.2,114.1" {
		t.Fatal(`expect "22.2,114.1", but get "` + loc + `"`)
	}
}
