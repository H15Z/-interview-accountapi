package restclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// API client defaults
// change as needed for production
// or refactor to use Enviroment Variables
const default_host string = "http://localhost:8080/v1"

type postRequestBody struct {
	Data interface{} `json:"data"`
}

type errResponse struct {
	Error string `json:"error_message"`
}

type RestClient struct {
	Host          string
	HTTPClient    *http.Client
	Authorization string //Not used in exercise
}

//Create client
func NewClient(host string) *RestClient {

	return &RestClient{
		Host:       host,
		HTTPClient: &http.Client{},
	}
}

//Returns client defaults
func Defaults() string {
	return default_host
}

// func (c *RestClient) DeleteRequest(req *http.Request, v interface{}) error {

// }

func (c RestClient) GetRequest(ctx context.Context, resource string) ([]byte, error) {

	return []byte{}, nil
}

func (c RestClient) PostRequest(ctx context.Context, resource string, d interface{}) ([]byte, error) {

	//create request
	req, err := c.buildPostRequest(ctx, resource, d)
	if err != nil {
		return []byte{}, err
	}

	// TODO move to sparate function to reuse by other request types
	//execute request
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	//handle response
	return c.handleResponse(res)
}

func (c RestClient) buildPostRequest(ctx context.Context, resource string, d interface{}) (*http.Request, error) {

	var req *http.Request
	var err error

	//create payload
	url := c.Host + resource
	buff := new(bytes.Buffer)
	body := postRequestBody{
		Data: d,
	}

	err = json.NewEncoder(buff).Encode(body)
	if err != nil {
		return req, err
	}

	//create request
	req, err = http.NewRequest("POST", url, buff)
	if err != nil {
		return req, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json")

	/*
		headers - not used but present in production
		'Authorization: {{authorization}}'
		'Digest: {{request_signing_digest}}'
	*/

	return req, nil

}

func (c RestClient) handleResponse(res *http.Response) ([]byte, error) {
	b, err := io.ReadAll(res.Body)

	if err != nil {
		return b, err
	}

	// looking for status 2xx
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {

		// parse "error_msg" from response into new error
		var e errResponse
		err = json.Unmarshal(b, &e)

		if err != nil {
			return b, err
		}

		return b, errors.New(e.Error)

	} else {
		return b, nil
	}
}
