/*
 * @Description: sql工具
 * @Author: Gavin
 * @Date: 2022-07-20 12:48:34
 * @LastEditTime: 2022-09-15 15:26:28
 * @LastEditors: Gavin
 */
package utils

import (
	"fmt"
	"split-server/config"
	"split-server/model/RBAC/request"
	bzReq "split-server/model/business/request"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
 * @description: 方法说明....
 * @param {*} cb 内容cb
 * @return {*}
 * @Date: 2022-07-20 13:53:33
 */

type SqlType struct {
	db *gorm.DB
}

var GAA_SQL = new(SqlType)

//初始化cb调用方法
func InitSQL(cb func(*gorm.DB) (any, error)) (any, error) {

	dsn := config.GetDsn()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return "", err
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	res, err2 := cb(db)
	return res, err2

}

func (sql *SqlType) StartSQL() (db *gorm.DB, err error) {

	dsn := config.GetDsn()

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("数据库连接异常,%s", err.Error())
		return
	}

	sp := request.SysPermission{}
	sr := request.SysRole{}
	sul := request.SysUserLogin{}
	sui := request.SysUserInfo{}
	// sua := request.SysUserAssociate{}
	bz := bzReq.BizComment{}
	bt := request.BilTable{}
	br := request.BilRecord{}
	db.AutoMigrate(&sr, &sul, &sui, &sp, &bz, &bt, &br)
	// db.AutoMigrate(&bt, &br)
	sql.db = db
	return db, err
}

func (sql *SqlType) GetDB() (db *gorm.DB) {

	return sql.db

}
