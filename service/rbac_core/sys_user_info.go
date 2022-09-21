/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-22 16:05:43
 * @LastEditTime: 2022-09-19 15:05:18
 * @LastEditors: Gavin
 */
package rbac_core

import (
	"split-server/model/RBAC/request"
	"split-server/model/global"
	"split-server/utils"
)

type UserInfo struct {
}

func (c UserInfo) CreateItem(body request.SysUserInfo) (*request.SysUserInfo, error) {
	db := utils.GAA_SQL.GetDB()
	err := db.Create(&body).Error
	return &body, err
}
func (c UserInfo) UpdateItem(body request.SysUserInfo) (*request.SysUserInfo, error) {
	db := utils.GAA_SQL.GetDB()
	err := db.Model(&body).Updates(body).Error
	return &body, err
}

func (c UserInfo) GetItemById(body *global.PrimaryUUID) (*request.SysUserInfo, error) {
	db := utils.GAA_SQL.GetDB()
	result := request.SysUserInfo{}
	err := db.Model(&result).First(&result, "id = ?", body.ID).Error
	return &result, err
}

//通过like 去检索 最多20条
func (c UserInfo) GetInfoListByName(body *global.SearchParams) ([]request.SysUserInfo, error) {
	db := utils.GAA_SQL.GetDB()
	var result []request.SysUserInfo
	limit := 20

	err := db.Model(&request.SysUserInfo{}).Where("name LIKE ?", "%"+body.Searchkey+"%").Limit(limit).Find(&result).Error

	return result, err

}
