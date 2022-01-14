package restclient

import (
	"context"
	"errors"
	"net/http"
)

//Delete API Request
//Accepts Context and Form3 API resource link
//Returns request status code , error
//API Does not return body for delete requests
func (c RestClient) DeleteRequest(ctx context.Context, resource string) (int, error) {
	url := c.Host + resource

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return 0, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/vnd.api+json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	// looking for status codes 2xx
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return res.StatusCode, errors.New("delete failed with status code: " + res.Status)
	} else {
		return res.StatusCode, nil
	}
}
