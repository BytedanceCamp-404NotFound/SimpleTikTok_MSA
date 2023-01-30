package commonerror

type CommonErr int64

const (
	CommonErr_UNKNOW_ERROR      CommonErr = 0   // 未知错误
	CommonErr_DB_ERROR          CommonErr = 13  // 通用的数据库错误
	CommonErr_STATUS_OK         CommonErr = 200 // 正常, 目前只是为了展示警告信息
	CommonErr_PARAMETER_FAILED  CommonErr = 400 // 参数错误
	CommonErr_PAGE_NOT_EXIT     CommonErr = 404 // 请求网页不存在
	CommonErr_INTERNAL_ERROR    CommonErr = 500 // 服务内部错误
	CommonErr_TIMEOUT           CommonErr = 504 // 超时
	CommonErr_PARSE_TOKEN_ERROR CommonErr = 700 //解析token失败
)

var (
	CommonErr_name = map[int32]string{
		0:   "UNKNOW_ERROR",
		13:  "DB_ERROR",
		200: "STATUS_OK",
		400: "PARAMETER_FAILED",
		404: "PAGE_NOT_EXIT",
		500: "INTERNAL_ERROR",
		504: "TIMEOUT",
		700: "PARSE_TOKEN_ERROR",
	}
	CommonErr_value = map[string]int32{
		"UNKNOW_ERROR":      0,
		"DB_ERROR":          13,
		"STATUS_OK":         200,
		"PARAMETER_FAILED":  400,
		"PAGE_NOT_EXIT":     404,
		"INTERNAL_ERROR":    500,
		"TIMEOUT":           504,
		"PARSE_TOKEN_ERROR": 700,
	}
)
