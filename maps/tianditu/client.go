package tianditu

import (
	"net"
	"net/http"
	"time"
)

const BASE_URL = `https://api.tianditu.gov.cn`

type Client struct {
	apiKey     string
	HttpClient *http.Client
}

func NewClient(key string) *Client {
	return &Client{apiKey: key, HttpClient: DefaultHttpClient}
}

var (
	timeout           = 30 * time.Second
	DefaultHttpClient = &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   timeout,
				KeepAlive: 90 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 30 * time.Second,
			MaxIdleConnsPerHost:   100,
		},
	}
)
