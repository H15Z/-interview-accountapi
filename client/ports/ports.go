package ports

// use inferface for dependecy injection
// provided as separate package for reuse in additional API resources not just accounts
type RestClient interface {
	PostRequest(string, string, interface{}) error
}
