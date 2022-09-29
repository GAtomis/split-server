/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-09-26 19:41:33
 * @LastEditTime: 2022-09-27 11:31:52
 * @LastEditors: Gavin
 */
package rbac_core

import (
	"split-server/model/RBAC/request"
	"split-server/utils"

	"gorm.io/gorm"
)

type UserCharts struct {
}

func (uc *UserCharts) GetRecordsBySeasons(UUID string) (map[any]any, error) {

	db := utils.GAA_SQL.GetDB()

	err := db.Transaction(func(tx *gorm.DB) error {
		userInfo := request.SysUserInfo{
			ID: UUID,
		}
		var tables []request.BilTable
		lastWeek := "2022-09-01 00:00:00"
		today := "2022-09-30 23:59:59"
		//查询所有归属用户的table
		if err := tx.Model(&userInfo).Where("created_at BETWEEN ? AND ?", lastWeek, today).Association("BilTables").Find(&tables); err != nil {

			return err
		}

		var tableIds []string

		for i := 0; i < len(tables); i++ {
			tableIds[i] = tables[i].ID
		}
		
		var records []request.BilRecord
		if err:= tx.Model(&request.BilRecord{}).Where("id IN ?" ,tableIds).Find(&records)

		return nil

	})

	return nil, err

}
