package location

import (
	"testing"
)

func TestCode(t *testing.T) {

	testdatas := [][3]string{
		[3]string{"山东省", "370000", ""},
		[3]string{"临沂市", "371300", "0539"},
		[3]string{"沂水县", "371323", ""},
		[3]string{"北京市", "110000", "010"},
	}

	for _, testdate := range testdatas {
		adcode := GetAdcode(testdate[0])
		if adcode != testdate[1] {
			t.Fatal(`[GetAdcode] expect "` + testdate[1] + `", but get "` + adcode + `"`)
		}

		name := GetNameByAdcode(testdate[1])
		if name != testdate[0] {
			t.Fatal(`[GetNameByAdcode] expect "` + testdate[0] + `", but get "` + name + `"`)
		}

		//county
		if testdate[2] != "" {
			citycode := GetCitycode(testdate[0])
			if citycode != testdate[2] {
				t.Fatal(`[GetCitycode] expect "` + testdate[2] + `", but get "` + citycode + `"`)
			}

			name := GetNameByCitycode(testdate[2])
			if name != testdate[0] {
				t.Fatal(`[GetNameByCitycode] expect "` + testdate[0] + `", but get "` + name + `"`)
			}
		}
	}
}
