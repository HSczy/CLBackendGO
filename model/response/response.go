package response

type Msg struct {
	Msg    string         `json:"msg"`
}

type Data struct {
	Msg    string         `json:"msg"`
	Data   interface{} `json:"data"`
}

func OK() Msg {
	return Msg{Msg: "操作成功"}
}

func Fail() Msg {
	return Msg{Msg: "操作失败"}
}

func WithMessage(message string) Msg {
	return Msg{
		Msg: message,
	}
}

func WithData(message string,data interface{}) Data {
	return Data{
		Msg: message,
		Data: data,
	}
}

