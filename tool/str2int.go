package tool

import (
	"strconv"
)

//str 要转换的字符串
//defaultValue() 失败后的默认值
func Int64(str string, defaultValue func() int64) int64 {
	res, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return defaultValue()
	}
	return res
}
