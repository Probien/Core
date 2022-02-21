package common

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Token   string      `json:"token,omitempty"`
	Help    string      `json:"help,omitempty"`
}
