package errors

// Types of Razorpay Errors
const (
	BAD_REQUEST_ERROR = "BAD_REQUEST_ERROR"
	SERVER_ERROR      = "SERVER_ERROR"
	GATEWAY_ERROR     = "GATEWAY_ERROR"
)

type Error struct {
	InternalErrorCode string `json:"internal_error_code"`
	Field             string `json:"field"`
	Description       string `json:"description"`
	Code              string `json:"code"`
}

func (e *Error) GetInternalErrorCode() string {
	return e.InternalErrorCode
}

func (e *Error) GetDescription() string {
	return e.Description
}

type ErrorJson struct {
	ErrorData Error `json:"error"`
}

type BadRequestError struct {
	Message string
	Err     error
}

type ServerError struct {
	Message string
	Err     error
}

type GatewayError struct {
	Message string
	Err     error
}

type SignatureVerificationError struct {
	Message string
	Err     error
}

func (s *BadRequestError) Error() string {
	var error_message string

	if s == nil {
		return "<nil>"
	}

	error_message = ""
	if s.Message != "" {
		error_message += s.Message
	}

	error_message += s.Err.Error()

	return error_message
}

func (s *GatewayError) Error() string {
	var error_message string

	if s == nil {
		return "<nil>"
	}

	error_message = ""
	if s.Message != "" {
		error_message += s.Message
	}

	error_message += s.Err.Error()

	return error_message
}

func (s *ServerError) Error() string {
	var error_message string

	if s == nil {
		return "<nil>"
	}

	error_message = ""
	if s.Message != "" {
		error_message += s.Message
	}

	error_message += s.Err.Error()

	return error_message
}

func (s *SignatureVerificationError) Error() string {
	var error_message string

	if s == nil {
		return "<nil>"
	}

	error_message = ""
	if s.Message != "" {
		error_message += s.Message
	}

	error_message += s.Err.Error()

	return error_message
}
