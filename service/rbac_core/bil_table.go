/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-08-21 22:26:04
 * @LastEditTime: 2022-09-14 21:41:48
 * @LastEditors: Gavin
 */
package rbac_core

import (
	"split-server/model/RBAC/request"
	"split-server/model/global"
	"split-server/utils"
)

type BilTable struct {
}

// //查询列表
func (b *BilTable) GetList(body *request.BilTable, info *global.PageInfo) (map[string]any, error) {

	db := utils.GAA_SQL.GetDB()
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	var result []request.BilTable
	var total int64 = 0
	if err1 := db.Model(body).Count(&total).Error; err1 != nil {
		return nil, err1
	}
	if err2 := db.Model(body).Limit(limit).Offset(offset).Find(&result).Error; err2 != nil {
		return map[string]any{"item": nil, "total": total}, err2
	}
	return map[string]any{"item": result, "total": total}, nil

}

func (b *BilTable) CreateItem(body *request.BilTable) (*request.BilTable, error) {

	db := utils.GAA_SQL.GetDB()
	err := db.Create(body).Error
	return body, err

}

// func (p *Permission) UpdateItem(body *request.SysPermission) (*request.SysPermission, error) {
// 	db := utils.GAA_SQL.GetDB()
// 	err := db.Model(body).Updates(body).Error
// 	return body, err
// }

// func (p *Permission) DelItem(body *request.SysPermission) (*request.SysPermission, error) {
// 	db := utils.GAA_SQL.GetDB()
// 	err := db.Delete(body).Error
// 	return body, err
// }
func (b *BilTable) GetItem(body *global.PrimaryUUID) (*request.BilTable, error) {
	db := utils.GAA_SQL.GetDB()
	result := request.BilTable{}
	err := db.Model(&result).Preload("BilRecords").First(&result, "id = ?", body.ID).Error
	return &result, err
}
