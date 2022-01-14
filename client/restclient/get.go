package restclient

import (
	"context"
	"net/http"
)

func (c RestClient) GetRequest(ctx context.Context, resource string) ([]byte, error) {
	req, err := c.buildGetRequest(ctx, resource)
	if err != nil {
		return []byte{}, err
	}

	return c.doRequest(req)
}

func (c RestClient) buildGetRequest(ctx context.Context, resource string) (*http.Request, error) {
	url := c.Host + resource
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return req, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/vnd.api+json")

	/*
		headers - not used but present in production
		'Authorization: {{authorization}}'
		'Digest: {{request_signing_digest}}'
	*/

	return req, nil
}
