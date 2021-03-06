package restclient

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
)

// API client defaults
// change as needed for production
// or refactor to use Enviroment Variables
const default_host string = "http://localhost:8080"

//possibly move this to models
type postRequestBody struct {
	Data interface{} `json:"data"`
}

type errResponse struct {
	Error   string `json:"error_message"`
	Message string `json:"message"`
}

//Rest API Client struct
//Handles HTTP(s) requests made to Form3 API
type RestClient struct {
	Host          string
	HTTPClient    *http.Client
	Authorization string //Not used in exercise but present in production
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

	//get API_HOST env variable if available
	host := os.Getenv("API_HOST")

	if host != "" {
		return host
	}

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

		return b, errors.New(e.Error + e.Message) //errors provided in different formats

	} else {
		return b, nil
	}
}
