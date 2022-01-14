package restclient

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

func (c RestClient) PostRequest(ctx context.Context, resource string, d interface{}) ([]byte, error) {
	//create request
	req, err := c.buildPostRequest(ctx, resource, d)
	if err != nil {
		return []byte{}, err
	}

	return c.doRequest(req)
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
