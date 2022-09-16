/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-22 16:05:43
 * @LastEditTime: 2022-09-15 10:41:22
 * @LastEditors: Gavin
 */
package rbac_core

import (
	"split-server/model/RBAC/request"
	"split-server/utils"
)

type UserLogin struct {
}

func (c UserLogin) CreateItem(body request.SysUserLogin) (*request.SysUserLogin, error) {
	db := utils.GAA_SQL.GetDB()
	err := db.Create(&body).Error
	return &body, err
}
func (c UserLogin) UpdateItem(body request.SysUserLogin) (*request.SysUserLogin, error) {
	db := utils.GAA_SQL.GetDB()
	err := db.Model(&body).Updates(body).Error
	return &body, err
}
