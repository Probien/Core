package common

type QueryResponse struct {
	TotalPages int
	Previous   string
	Next       string
	Data       interface{}
}
