/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-09-20 11:39:05
 * @LastEditTime: 2022-10-03 14:43:10
 * @LastEditors: Gavin
 */
package rbac_core

import (
	"fmt"
	"split-server/model/RBAC/request"
	"split-server/model/global"
	"split-server/utils"
)

type BilRecord struct {
}

func (r BilRecord) GetRecordList(body *request.BilRecord, info *global.PageInfo) (map[string]any, error) {
	db := utils.GAA_SQL.GetDB()
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	var result []request.BilRecord
	var total int64 = 0

	fmt.Printf("body.TableId: %v\n", body.TableId)
	fmt.Printf("body.CreatorId: %v\n", body.CreatorId)
	if err1 := db.Model(body).Where("creator_id = ? AND table_id = ?", body.CreatorId, body.TableId).Count(&total).Error; err1 != nil {
		return nil, err1
	}
	if err2 := db.Model(body).Where("creator_id = ? AND table_id = ?", body.CreatorId, body.TableId).Limit(limit).Offset(offset).Preload("Creator").Find(&result).Error; err2 != nil {
		return map[string]any{"item": nil, "total": total}, err2
	}
	return map[string]any{"item": result, "total": total}, nil
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
