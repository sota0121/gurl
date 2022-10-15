package proxy

import (
	"fmt"
	"net/http"
)

func NewHttpClient() *HttpClient {
	return &HttpClient{}
}

type HttpClient struct {
	connected bool
}

func (c *HttpClient) IsConnected() bool {
	return c.connected
}

func (c *HttpClient) Connect() error {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	c.connected = true
	return nil
}
