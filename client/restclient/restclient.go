package restclient

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// API client defaults
// change as needed for production
// or refactor to use Enviroment Variables
const default_host string = "http://localhost:8080/v1"

//possibly move this to models
type postRequestBody struct {
	Data interface{} `json:"data"`
}

type errResponse struct {
	Error string `json:"error_message"`
}

//Rest API Client struct
//Handles HTTP(s) requests made to API
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

func (c RestClient) doRequest(req *http.Request) ([]byte, error) {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	//handle response
	return c.handleResponse(res)
}

//handle API response
//for now used by Post and Get, Delete does not have a response body
func (c RestClient) handleResponse(res *http.Response) ([]byte, error) {
	b, err := io.ReadAll(res.Body)

	if err != nil {
		return b, err
	}

	// looking for status codes 2xx
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
