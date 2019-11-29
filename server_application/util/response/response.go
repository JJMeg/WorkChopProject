package response

import "github.com/jjmeg/WorkChopProject/util/errors"

type ErrorResponse struct {
	RequestId string       `json:"request_id"`
	Err       errors.Error `json:"error"`
}

func NewErrorResponse(reId string, err errors.Error) *ErrorResponse {
	return &ErrorResponse{
		RequestId: reId,
		Err:       err,
	}
}

func (e *ErrorResponse) StatusCode() int {
	return e.Err.Code
}

func (e *ErrorResponse) Error() string {
	return e.Err.Error()
}
