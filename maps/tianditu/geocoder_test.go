package tianditu

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pkgplus/location/maps"
	"github.com/stretchr/testify/assert"
)

func TestClient_Geocoder(t *testing.T) {
	// 模拟成功的 API 响应
	successResponse := GeocoderResponse{
		Msg:           "ok",
		SearchVersion: "6.4.9V",
		Status:        "0",
		Location: struct {
			Score   int    `json:"score"`
			Level   string `json:"level"`
			Lon     string `json:"lon"`
			Lat     string `json:"lat"`
			KeyWord string `json:"keyWord"`
		}{
			Score:   100,
			Level:   "门址",
			Lon:     "116.290158",
			Lat:     "39.894696",
			KeyWord: "北京市海淀区莲花池西路28号",
		},
	}

	// 模拟失败的 API 响应
	errorResponse := GeocoderResponse{
		Msg:           "invalid key",
		SearchVersion: "6.4.9V",
		Status:        "1",
	}

	// 测试用例
	tests := []struct {
		name           string
		mockResponse   GeocoderResponse
		expectedResult []*maps.Geocode
		expectedError  string
	}{
		{
			name:           "Success",
			mockResponse:   successResponse,
			expectedResult: []*maps.Geocode{successResponse.Geocode()},
			expectedError:  "",
		},
		{
			name:           "Error",
			mockResponse:   errorResponse,
			expectedResult: nil,
			expectedError:  "API returned error: invalid key",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建模拟的 HTTP 服务器
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				if err := json.NewEncoder(w).Encode(tt.mockResponse); err != nil {
					t.Fatalf("failed to encode mock response: %v", err)
				}
			}))
			defer server.Close()

			// 创建 Client 实例
			client := &Client{
				apiKey:     "test-key",
				HttpClient: server.Client(),
			}

			// 调用 Geocoder 方法
			result, err := client.Geocoder(context.Background(), "北京市海淀区莲花池西路28号")

			// 检查结果
			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, result)
			}
		})
	}
}
