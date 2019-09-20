package fake

type fakeError struct {
	msg        string
	statusCode int
	reqId      string
}

func New(msg string, statusCode int, reqId string) *fakeError {
	return &fakeError{
		msg:        msg,
		statusCode: statusCode,
		reqId:      reqId,
	}
}

func (fe fakeError) Error() string {
	return fe.msg
}

func (fe *fakeError) StatusCode() int {
	return fe.statusCode
}

func (fe *fakeError) RequestID() string {
	return fe.reqId
}

func (fe fakeError)  Code() string {
	return "NotImplemented"
}

func (fe fakeError)  Message() string {
	return fe.msg
}

func (fe fakeError)  OrigErr() error {
	return nil
}
