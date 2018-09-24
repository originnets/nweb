package models

const (
	RECODE_OK         = "0"
	RECODE_DBERR      = "2001"
	RECODE_NODATA     = "2002"
	RECODE_DATAEXIST  = "2003"
	RECODE_DATAERR    = "2004"
	RECODE_SESSIONERR = "2101"
	RECODE_LOGINERR   = "2102"
	RECODE_PARAMERR   = "2103"
	RECODE_USERERR    = "2104"
	RECODE_ROLEERR    = "2105"
	RECODE_PWDERR     = "2106"
	RECODE_REQERR     = "2201"
	RECODE_IPERR      = "2202"
	RECODE_THIRDERR   = "2301"
	RECODE_IOERR      = "2302"
	RECODE_SERVERERR  = "2500"
	RECODE_UNKNOWERR  = "2501"
)

var ReText = map[string]string{
	RECODE_OK:         "成功",
	RECODE_DBERR:      "数据库查询错误",
	RECODE_NODATA:     "无数据",
	RECODE_DATAEXIST:  "数据已存在",
	RECODE_DATAERR:    "数据错误",
	RECODE_SESSIONERR: "用户未登录",
	RECODE_LOGINERR:   "用户登录失败",
	RECODE_PARAMERR:   "参数错误",
	RECODE_USERERR:    "用户不存在或未激活",
	RECODE_ROLEERR:    "用户身份错误",
	RECODE_PWDERR:     "用户名或密码错误",
	RECODE_REQERR:     "非法请求或请求次数受限",
	RECODE_IPERR:      "IP受限",
	RECODE_THIRDERR:   "第三方系统错误",
	RECODE_IOERR:      "文件读写错误",
	RECODE_SERVERERR:  "内部错误",
	RECODE_UNKNOWERR:  "未知错误",
}

func ReCodeText(code string) string {
	str, ok := ReText[code]
	if ok {
		return str
	}
	return ReText[RECODE_UNKNOWERR]
}