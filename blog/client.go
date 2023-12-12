package blog

import "net/http"

type Client struct {
	http *http.Client
}

func NewClient() *Client {
	return &Client{
		http: &http.Client{},
	}
}

func (c *Client) do(r *http.Request) (*http.Response, error) {
	r.Header.Set("User-Agent", "ucblogmonitor/0.1 (https://github.com/nsigel/ucblogmonitor)")

	return c.http.Do(r)
}
