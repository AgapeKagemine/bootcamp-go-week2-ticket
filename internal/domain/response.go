package domain

type ResponseBody struct {
	StatusCode uint   `json:"status_code"`
	Message    string `json:"message"`
	Payload    any    `json:"payload,omitempty"`
}
