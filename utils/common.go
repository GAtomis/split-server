/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-09-05 16:20:28
 * @LastEditTime: 2022-09-05 16:31:41
 * @LastEditors: Gavin
 */
package utils

import (
	"encoding/hex"
	"math/rand"
	"time"
)

type Common struct {
}

// 长度为62
// var bytes []byte = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890")

func init() {
	// 保证每次生成的随机数不一样
	rand.Seed(time.Now().UnixNano())
}

func RandStr2(n int) string {
	result := make([]byte, n/2)
	rand.Read(result)
	return hex.EncodeToString(result)
}
