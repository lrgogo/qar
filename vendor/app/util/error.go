package util

type Response struct {
	Code int            `json:"code"`
	Msg  string        `json:"msg"`
	Data interface{}    `json:"data,omitempty"`
}

func (e *Response) Error() string {
	return e.Msg
}

func Error(code int, msg string) error {
	return &Response{
		Code: code,
		Msg:  msg,
	}
}

func Error3(code int, msg string, data interface{}) error {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func Success(data interface{}) error {
	return &Response{
		Code: SUCCESS,
		Msg:  "",
		Data: data,
	}
}
