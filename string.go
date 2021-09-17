/*
 * @Author: ybc
 * @Date: 2020-08-20 20:25:41
 * @LastEditors: ybc
 * @LastEditTime: 2020-08-20 20:32:22
 * @Description: file content
 */
package goutils

import (
	"strings"
)

func StrFirstToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArry := []rune(str)
	if strArry[0] >= 97 && strArry[0] <= 122 {
		strArry[0] -= 32
	}
	return string(strArry)
}

//转换为小写
func StrToLower(str string) string {
	if len(str) < 1 {
		return ""
	}
	var newStr []rune
	strArray := []rune(str)
	for _, v := range strArray {
		if v >= 65 && v <= 90 {
			v += 32
		}
		newStr = append(newStr, v)
	}
	return string(newStr)
}

//
func StrUnderlineToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArry := strings.Split(str, "_")
	var newStr string
	for _, v := range strArry {
		newStr += StrFirstToUpper(v)
	}

	return newStr
}

func StrToUnderlineWithLower(str string) string {
	if len(str) < 1 {
		return ""
	}
	var newStr []rune
	strArry := []rune(str)
	for _, v := range strArry {
		if v < 91 {
			newStr = append(newStr, 95)
			v += 32
		}
		newStr = append(newStr, v)
	}

	return strings.TrimLeft(string(newStr), "_")
}
