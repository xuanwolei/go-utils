/*
 * @Author: ybc
 * @Date: 2020-07-13 15:47:43
 * @LastEditors: ybc
 * @LastEditTime: 2020-08-20 16:35:29
 * @Description: file content
 */

package goutils

import (
	"strings"
	"time"
)

const TIME_LAOUT string = "2006-01-02 15:04:05"

func FormatIso8601(value string) string {
	return strings.Replace(strings.Replace(value, "T", " ", 1), "+08:00", "", -1)
}

func Ios8601ToUnix(value string) int64 {
	value = FormatIso8601(value)
	//修改为本地时区
	t, _ := time.ParseInLocation(TIME_LAOUT, value, time.Local)
	return t.Unix()
}
