package rest

type RestError struct {
	StatusCode int
	ErrorCode  string
}

func (re *RestError) Error() string {
	return re.ErrorCode
}

var NotFoundError = &RestError{StatusCode: 404, ErrorCode: "not_found"}
