package rest

type RestError struct {
	StatusCode int
	ErrorCode  string
}

func (re *RestError) Error() string {
	return re.ErrorCode
}

var NotFoundError = RestError{StatusCode: 404, ErrorCode: "not_found"}
var InternalServerError = RestError{StatusCode: 500, ErrorCode: "internal_server_error"}
