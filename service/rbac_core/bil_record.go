/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-09-20 11:39:05
 * @LastEditTime: 2022-09-24 22:05:25
 * @LastEditors: Gavin
 */
package rbac_core

import (
	"split-server/model/RBAC/request"
	"split-server/model/global"
	"split-server/utils"
)

type BilRecord struct {
}

func (r BilRecord) GetRecordList(body *global.PrimaryUUID) ([]request.BilRecord, error) {

	db := utils.GAA_SQL.GetDB()
	var result []request.BilRecord
	u := body.ID
	err := db.Model(&request.BilRecord{}).Where("table_id = ?", u).Preload("Creator").Find(&result).Error

	return result, err
	// return

}

func (r *BilRecord) CreateItem(body *request.BilRecord) (*request.BilRecord, error) {
	db := utils.GAA_SQL.GetDB()
	err := db.Model(body).Omit("Creator").Create(body).Error
	return body, err
}
func (r *BilRecord) UpdateItem(body *request.BilRecord) (*request.BilRecord, error) {
	db := utils.GAA_SQL.GetDB()
	err := db.Model(body).Omit("Creator", "ID").Updates(body).Error
	return body, err
}
func (r *BilRecord) DeleteItem(body *request.BilRecord) (*request.BilRecord, error) {
	db := utils.GAA_SQL.GetDB()
	err := db.Delete(body).Error
	return body, err
}
