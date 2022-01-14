package accounts

import "context"

/*
	stub rest client to inject into Accounts for testing
	simulates responses from restclient
	allows for decoupled testing
	could be moved to separate package for reuse in other API resources
*/

type stubClient struct {
	PostResponse   stubResponse
	GetResponse    stubResponse
	DeleteResponse stubResponse
}

type stubResponse struct {
	ResponseCode int
	Body         []byte
	Error        error
}

func (s stubClient) PostRequest(ctx context.Context, resource string, d interface{}) ([]byte, error) {
	return s.PostResponse.Body, s.PostResponse.Error
}

func (s stubClient) GetRequest(ctx context.Context, resource string) ([]byte, error) {
	return s.GetResponse.Body, s.GetResponse.Error
}

func (s stubClient) DeleteRequest(ctx context.Context, resource string) (int, error) {
	return s.DeleteResponse.ResponseCode, s.DeleteResponse.Error
}
