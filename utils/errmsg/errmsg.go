package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// code = 1000 用户模块错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NULL        = 1008
	ERROR_USER_NOT_RIGHT = 1009

	// code = 2000 骗子信息模块错误
	ERROR_SWINDLER_NOT_EXIST = 2001
)

var codeMsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已存在",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_EXIST:      "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期",
	ERROR_TOKEN_WRONG:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",
	ERROR_USER_NULL:        "用户名和密码不能为空",
	ERROR_USER_NOT_RIGHT: "该用户无权限",

	ERROR_SWINDLER_NOT_EXIST: "骗子信息不存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
