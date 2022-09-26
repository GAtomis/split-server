/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-20 16:13:20
 * @LastEditTime: 2022-09-25 12:06:31
 * @LastEditors: Gavin
 */
package config

const (
	SERVER_PORT = "9527"
	ROOT        = "rapi"
)

func GetPort() string {
	return SERVER_PORT
}
