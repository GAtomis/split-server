/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-20 13:59:03
 * @LastEditTime: 2022-08-16 23:28:04
 * @LastEditors: Gavin
 */
package config

import "fmt"

const (
	SQL_URL = "ec2-3-112-56-234.ap-northeast-1.compute.amazonaws.com"
	// SQL_URL    = "127.0.0.1"
	SQL_NAME   = "split"
	SQL_CONFIG = "?charset=utf8mb4&parseTime=True&loc=Local"
	USERNAME   = "root"
	PASSWORD   = "asdADSDC!!!UJNk"
	SQL_PORT   = "3306"
	PROTOCOL   = "tcp"
	TIME_OUT   = "30s"
)

/**
 * @description: 获取sql连接
 * @return {*}
 * @Date: 2022-07-20 14:06:48
 */
func GetDsn() (dsn string) {
	dsn = fmt.Sprintf("%s:%s@%s(%s:%s)/%s%s&%s", USERNAME, PASSWORD, PROTOCOL, SQL_URL, SQL_PORT, SQL_NAME, SQL_CONFIG, TIME_OUT)
	fmt.Println(dsn)
	return
}
