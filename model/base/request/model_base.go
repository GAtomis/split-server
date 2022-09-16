/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-21 16:16:23
 * @LastEditTime: 2022-09-14 21:41:25
 * @LastEditors: Gavin
 */
package request

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

type Register struct {
	Code string `json:"code"`
}
