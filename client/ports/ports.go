package ports

import "context"

// use inferface for dependecy injection
// provided as separate package for reuse in additional API resources (useful for future development)
type RestClient interface {
	PostRequest(ctx context.Context, resource string, d interface{}) ([]byte, error)
	GetRequest(ctx context.Context, resource string) ([]byte, error)
	DeleteRequest(ctx context.Context, resource string) (int, error)
}
