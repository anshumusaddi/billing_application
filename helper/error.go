package helper

import "fmt"

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
	ErrDBInsert                    = NewAPIError(500, "BA1002", "E_DB_INSERT", "error while inserting the document in the database")
	ErrDuplicateKey                = NewAPIError(400, "BA1003", "E_DB_OPERATION", "key already exists, check if key is already present")
	ErrInvalidQueryParams          = NewAPIError(400, "BA1004", "E_BAD_REQUEST", "invalid query params")
	ErrDBOperation                 = NewAPIError(500, "BA1005", "E_DB_OPERATION", "failed to execute the query in database")
	ErrKafkaInsert                 = NewAPIError(500, "BA1006", "E_KAFKA_INSERT", "error while inserting the document in kafka")
)

func ApiErrorWithCustomMessage(apiErr *APIError, message string) *APIError {
	apiErr.Message = message
	return apiErr
}

func (err APIError) Error() string {
	return fmt.Sprintf("[%s]: %s", err.Title, err.Message)
}
