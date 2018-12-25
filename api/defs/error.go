package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSc int
	Error  Err
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{HttpSc: 400, Error: Err{Error: "Request failed", ErrorCode: "001"}}
	ErrorNotAuthUser            = ErrorResponse{HttpSc: 401, Error: Err{Error: "User authentication failed.", ErrorCode: "002"}}
	ErrorDBError                = ErrorResponse{HttpSc: 500, Error: Err{Error: "DB ops failed", ErrorCode: "003"}}
	ErrorInternalFail           = ErrorResponse{HttpSc: 500, Error: Err{Error: "Internal service error", ErrorCode: "004"}}
)
