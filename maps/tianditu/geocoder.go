package tianditu

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkgplus/location/maps"
	"net/http"
	"net/url"
)

type GeocoderResponse struct {
	Msg           string `json:"msg"`
	SearchVersion string `json:"searchVersion"`
	Status        string `json:"status"`
	Location      struct {
		Score   int    `json:"score"`
		Level   string `json:"level"`
		Lon     string `json:"lon"`
		Lat     string `json:"lat"`
		KeyWord string `json:"keyWord"`
	} `json:"location"`
}

func (gr *GeocoderResponse) Geocode() *maps.Geocode {
	return &maps.Geocode{
		Address:  gr.Location.KeyWord,
		Province: "",
		City:     "",
		Level:    gr.Location.Level,
		Location: fmt.Sprintf("%s,%s", gr.Location.Lon, gr.Location.Lat),
	}
}

func (c *Client) Geocoder(ctx context.Context, keyword string) ([]*maps.Geocode, error) {
	// 构建请求 URL
	params := url.Values{}
	params.Set("ds", fmt.Sprintf(`{"keyWord":"%s"}`, keyword))
	params.Set("tk", c.apiKey)
	requestURL := fmt.Sprintf("%s%s?%s", BASE_URL, "/geocoder", params.Encode())

	// 发送 HTTP GET 请求
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("new request failed: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(ctx)
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// 解析响应
	var geocoderResponse GeocoderResponse
	if err := json.NewDecoder(resp.Body).Decode(&geocoderResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	// 检查状态
	if geocoderResponse.Status != "0" {
		return nil, fmt.Errorf("API returned error: %s", geocoderResponse.Msg)
	}

	return []*maps.Geocode{geocoderResponse.Geocode()}, nil
}
