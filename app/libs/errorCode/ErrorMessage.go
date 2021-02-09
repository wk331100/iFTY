package errorCode

import errors "github.com/wk331100/iFTY/system/error"

const(
	//响应状态码
	SUCCESS		= 200
	FAILED		= 201

	//通用状态码
	ERR_UNKNOWN = 1000
	ERR_DB		= 1001
	ERR_CACHE	= 1002

	//业务状态码
	ERR_PARAMS	= 2000
	ERR_NOT_EXIST = 2001

)

var ErrorCode = map[int]string{
	SUCCESS 		: "成功",
	FAILED 			: "失败",

	ERR_UNKNOWN 	: "未知错误",
	ERR_DB 			: "数据库错误",
	ERR_CACHE 		: "缓存错误",

	ERR_PARAMS 		: "参数错误",
	ERR_NOT_EXIST 	: "记录不存在",
}

/**
 * 获取错误信息： 优先从自定义中取，如果取不到，从系统内置错误码中取
 */
func GetErrorMessage(code int) string {
	if ErrorCode[code] != ""{
		return ErrorCode[code]
	}
	return errors.GetMessage(code, "")
}