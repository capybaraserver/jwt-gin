package ret

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	SuccessCode = 200
	FailCode    = 400
	ErrorCode   = 500
)

const (
	SuccessMsg = "success"
	FailMsg    = "fail"
	ErrorMsg   = "error"
)

func RetSuccess() Result {
	return Result{
		Code: SuccessCode,
		Msg:  SuccessMsg,
		Data: nil,
	}
}

func RetSuccessData(data interface{}) Result {
	return Result{
		Code: SuccessCode,
		Msg:  SuccessMsg,
		Data: data,
	}
}

func Ret(code int, msg string, data interface{}) Result {
	return Result{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func RetFail() Result {
	return Result{
		Code: FailCode,
		Msg:  FailMsg,
		Data: nil,
	}
}

func RetFailMsg(msg string) Result {
	return Result{
		Code: FailCode,
		Msg:  msg,
		Data: nil,
	}
}
