package restclient

import (
	"bytes"
	"net/http"
	"time"
)

// package defaults - change as needed for production - or use Enviroment Variables
const default_host string = "http://localhost:8080/v1"
const default_timeout time.Duration = 15 * time.Second

type RestClient struct {
	Host       string
	HTTPClient *http.Client
}

//Create client
func NewClient(host string, timeout time.Duration) *RestClient {

	return &RestClient{
		Host: host,
		HTTPClient: &http.Client{
			Timeout: default_timeout,
		},
	}
}

// Return defaults for
func Defaults() (string, time.Duration) {
	return default_host, default_timeout
}

// func (c *RestClient) GetRequest(req *http.Request, v interface{}) error {

// }

// func (c *RestClient) DeleteRequest(req *http.Request, v interface{}) error {

// }

func (c *RestClient) PostRequest(method string, resource string, d interface{}) error {

	buf := new(bytes.Buffer)
	url := c.Host + resource

	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return nil
}

type errResponse struct {
	Error string `json:"error_msg"`
}

// func handleResponse(res *http.Response) error {

// 	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
// 		return nil
// 	} else {

// 	}

// }
