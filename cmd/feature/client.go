package feature

import (
	"fmt"
	"net/http"
)

var (
	exampleDotCom = "http://example.com"
)

func NewGurlClient() *GurlClient {
	return &GurlClient{}
}

type GurlClient struct {
	connected bool
}

func (c GurlClient) IsConnected() bool {
	return c.connected
}

func (c *GurlClient) Connect() error {
	resp, err := http.Get(exampleDotCom)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	c.connected = true
	return nil
}

func (c GurlClient) Request(reqctx *ReqContext, url string) (*http.Response, error) {
	switch reqctx.Method {
	case http.MethodGet:
		return http.Get(url)
	// case http.MethodPost:
	// 	return http.Post(url, reqctx.ContentType, reqctx.Body)
	// case http.MethodPut:
	// 	return http.Put(url, reqctx.ContentType, reqctx.Body)
	default:
		return nil, fmt.Errorf("not implemented")
	}
}
