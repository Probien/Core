package response

type PaginatedResponse struct {
	Status       int         `json:"status"`
	ItemsPerPage int         `json:"items_per_page"`
	TotalPages   int         `json:"total_pages"`
	CurrentPage  int         `json:"current_page"`
	Data         interface{} `json:"data"`
	Previous     string      `json:"previous"`
	Next         string      `json:"next"`
}
