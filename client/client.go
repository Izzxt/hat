package client

import (
	"log"
	"net"
	"net/http"
	"os"
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
			KeepAlive: 24 * time.Second,
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

	req, err := http.NewRequest("GET", fetchUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	response, err := c.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	return response
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
