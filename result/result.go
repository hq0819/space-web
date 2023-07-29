package result

type Message struct {
	Code int8   `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Success(t any) *Message {
	return &Message{Code: 0, Msg: "成功", Data: t}
}

func Fail(msg string) *Message {
	return &Message{Code: -1, Msg: msg, Data: nil}
}
