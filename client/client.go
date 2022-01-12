package client

import (
	"net/http"
	"time"
)

// package defaults - change as needed for production
const default_host string = "http://localhost:8080/v1"
const default_timeout time.Duration = 15 * time.Second

type Client struct {
	Host       string
	HTTPClient *http.Client
}

//Create client
func NewClient(host string) *Client {

	return &Client{
		Host: host,
		HTTPClient: &http.Client{
			Timeout: default_timeout,
		},
	}
}

func (c *Client) GetRequest(req *http.Request, v interface{}) error {

}

func (c *Client) DeleteRequest(req *http.Request, v interface{}) error {

}

func (c *Client) PostRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {

	}

	return nil
}

type response struct {
}
