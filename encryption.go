/*
 * @Author: ybc
 * @Date: 2020-07-14 16:24:34
 * @LastEditors: ybc
 * @LastEditTime: 2020-08-20 20:01:35
 * @Description: file content
 */
package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
