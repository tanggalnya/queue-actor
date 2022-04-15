package apiModels

type Response struct {
	Errors []Err `json:"errors"`
}

type errCode string

type Err struct {
	Code    errCode `json:"code"`
	Message string  `json:"message"`
}

var ErrCodes = newErrCodes()

var (
	internalError errCode = "InternalError"
	payloadError  errCode = "PayloadError"
)

type ErrCode struct {
	InternalError errCode
	PayloadError  errCode
}

func newErrCodes() ErrCode {
	return ErrCode{
		InternalError: internalError,
		PayloadError:  payloadError,
	}
}
