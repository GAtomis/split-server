/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-08-21 22:26:04
 * @LastEditTime: 2022-09-25 14:09:13
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
	err := db.Model(body).Omit("Creator", "BilRecords").Create(body).Error
	return body, err
}

func (b *BilTable) GetItem(body *global.PrimaryUUID) (*request.BilTable, error) {
	db := utils.GAA_SQL.GetDB()
	result := request.BilTable{}
	err := db.Model(&result).Preload("BilRecords.Creator").Preload("SysUsers").Preload("Creator").First(&result, "id = ?", body.ID).Error
	return &result, err
}
func (b *BilTable) GetBillTableListByUserId(body *global.PrimaryUUID, pageInfo *global.PageInfo) (map[string]any, error) {
	db := utils.GAA_SQL.GetDB()

	userInfo := request.SysUserInfo{
		ID: body.ID,
	}
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	var result []request.BilTable

	total := db.Model(&userInfo).Association("BilTables").Count()

	if err := db.Model(&userInfo).Limit(limit).Offset(offset).Preload("Creator").Preload("SysUsers").Preload("BilRecords").Association("BilTables").Find(&result); err != nil {

		return map[string]any{"item": nil, "total": total}, err
	}
	return map[string]any{"item": result, "total": total}, nil
	// var result []request.BilTable

	// userInfo := request.SysUserInfo{
	// 	ID: body.ID,
	// }

}

// func (b *BilTable) UpdateSysUsers(body *request.BilTable) error {
// 	db := utils.GAA_SQL.GetDB()
// 	sui := body.SysUsers
// 	err := db.Model(body).Association("SysUsers").Replace(&sui)
// 	return err
// }
// func (b *BilTable) UpdateBilRecords(body *request.BilTable) error {
// 	db := utils.GAA_SQL.GetDB()
// 	br := body.BilRecords
// 	err := db.Model(body).Association("BilRecords").Replace(&br)
// 	return err
// }
func (b *BilTable) UpdateBilTableInfo(body *request.BilTable) (*request.BilTable, error) {
	db := utils.GAA_SQL.GetDB()
	err := db.Model(body).Omit("Creator", "SysUsers", "BilRecords", "ID").Updates(body).Error
	// err := db.Transaction(func(tx *gorm.DB) error {
	// 	if err := tx.Model(body).Omit("Creator", "SysUsers", "BilRecords", "ID").Updates(body).Error; err != nil {
	// 		return err
	// 	}
	// 	if err := tx.Model(body).Association("SysUsers").Replace(&body.SysUsers); err != nil {
	// 		return err
	// 	}
	// 	if err := tx.Model(body).Association("BilRecords").Replace(&body.BilRecords); err != nil {
	// 		return err
	// 	}

	// 	return nil
	// })

	return body, err
}
