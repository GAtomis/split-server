/*
 * @Description: 初始化
 * @Author: Gavin
 * @Date: 2022-07-18 15:11:28
 * @LastEditTime: 2022-09-14 10:30:42
 * @LastEditors: Gavin
 */
package main

import (
	"split-server/router"
	"split-server/utils"
)

/**
 * @description: 入口函数....
 * @return {*}
 * @Date: 2022-07-18 15:12:06
 */
//声明一个wg类型

func main() {

	db, _ := utils.GAA_SQL.StartSQL()
	if db != nil {
		sqlDB, _ := db.DB()
		defer sqlDB.Close()
	}

	router.Start()

}
