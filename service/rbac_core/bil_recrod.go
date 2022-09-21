/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-09-20 11:39:05
 * @LastEditTime: 2022-09-21 19:02:53
 * @LastEditors: Gavin
 */
package rbac_core

import (
	"split-server/model/RBAC/request"
	"split-server/model/global"
	"split-server/utils"
)

type BilRecrod struct {
}

func (r BilRecrod) GetRecrodList(body *global.PrimaryUUID) ([]request.BilRecord, error) {

	db := utils.GAA_SQL.GetDB()
	var result []request.BilRecord
	u := body.ID
	err := db.Model(&request.BilRecord{}).Where("table_id = ?", u).Preload("Creator").Find(&result).Error

	return result, err
	// return

}
