package util

const (
	ResultCode_Success = 200
	ResultCode_Error   = 300
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success() *Result {
	rs := new(Result)
	rs.Code = ResultCode_Success
	rs.Msg = "success"
	rs.Data = nil
	return rs
}

func SuccessWithData(data interface{}) *Result {
	rs := new(Result)
	rs.Code = ResultCode_Success
	rs.Msg = "success"
	rs.Data = data
	return rs
}

func Error() *Result {
	rs := new(Result)
	rs.Code = ResultCode_Error
	rs.Msg = "error"
	rs.Data = nil
	return rs
}
func ErrorWithMsg(msg string) *Result {
	rs := new(Result)
	rs.Code = ResultCode_Error
	rs.Msg = msg
	rs.Data = nil
	return rs
}
