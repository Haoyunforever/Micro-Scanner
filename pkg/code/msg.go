package code

var MsgFlags = map[int]string{
	SUCCESS:              "ok",
	ERROR:                "fail",
	INVALID_PARAMS:       "请求参数错误",
	ERROR_EXIST_USER:     "用户名已存在",
	ERROR_NOT_EXIST_USER: "用户名不存在",
	ERROR_NOT_COMPARE:    "用户名或密码不正确",
	ERROR_GEN_TOKEN:      "Token 生成失败",
}

func GetMsgFlag(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
