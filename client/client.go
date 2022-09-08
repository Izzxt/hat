package client

import (
	"log"
	"net"
	"net/http"
	"time"
)

type ClientConfig struct {
}

type Client struct {
	client http.Client
}

func NewClient() Client {
	var DefaultTransport http.RoundTripper = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:    100,
		IdleConnTimeout: 90 * time.Second,
	}

	client := http.Client{
		Transport: DefaultTransport,
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	return Client{client}
}

func (c *Client) Get(fetchUrl string) *http.Response {

	response, err := c.client.Get(fetchUrl)
	if err != nil {
		log.Fatal(err)
	}

	return response
}
