package helper

type APIError struct {
	Status  int    `json:"status"`
	Code    string `json:"err_code"`
	Title   string `json:"err_str"`
	Message string `json:"err_msg"`
}

func NewAPIError(status int, code string, title, message string) *APIError {
	return &APIError{
		Status:  status,
		Code:    code,
		Title:   title,
		Message: message,
	}
}

var (
	ErrInvalidRequestPayloadParams = NewAPIError(400, "BA1001", "E_BAD_REQUEST", "invalid payload in request, check if payload data types are correct")
)
