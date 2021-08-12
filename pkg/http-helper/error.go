package http_helper

type Error interface {
	Response

	Error() string
}

type HttpError struct {
	response
	StatusCode int    `json:"status"`
	Code       int    `json:"code"`
	System     string `json:"system"`
	Message    string `json:"message"`
}

func (err HttpError) Error() string {
	return err.Message
}

func (err HttpError) Response() interface{} {
	return err
}
