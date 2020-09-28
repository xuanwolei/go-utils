/*
 * @Author: ybc
 * @Date: 2020-08-19 15:25:30
 * @LastEditors: ybc
 * @LastEditTime: 2020-08-19 15:35:10
 * @Description: file content
 */
package utils

import (
	"strconv"
)

func InterfaceToInt(data interface{}) int {
	var res int
	switch v := data.(type) {
	case string:
		res, _ = strconv.Atoi(v)
	case int64:
		strInt64 := strconv.FormatInt(v, 10)
		res, _ = strconv.Atoi(strInt64)
	case int:
		res = v
	default:
		res = 0
	}

	return res
}

func ParseInt(value int, err error) int {
	if err != nil {
		return 0
	}
	return value
}

func ParseBool(value bool, err error) bool {
	if err != nil {
		return false
	}
	return value
}
