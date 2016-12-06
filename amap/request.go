package amap

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type ApiRequest struct {
	client *Client
}

func (req *ApiRequest) HttpGet(url string, v Response) error {
	//http client
	httpclient := req.client.HttpClient
	if httpclient == nil {
		httpclient = DefaultHttpClient
	}

	//http get
	resp, err := httpclient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//result
	resp_body, read_err := ioutil.ReadAll(resp.Body)
	if read_err != nil {
		return read_err
	}

	//json Unmarshal
	json_err := json.Unmarshal(resp_body, v)
	if v.GetStatus() != "1" {
		if json_err != nil {
			return json_err
		} else {
			return errors.New(v.GetInfo())
		}
	} else {
		return nil
	}
}
