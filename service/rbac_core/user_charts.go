/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-09-26 19:41:33
 * @LastEditTime: 2022-10-15 19:29:04
 * @LastEditors: Gavin
 */
package rbac_core

import (
	"split-server/model/RBAC/request"
	"split-server/utils"
)

type UserCharts struct {
}

func (uc *UserCharts) GetRecordsBySeasons(UUID string) (any, error) {

	db := utils.GAA_SQL.GetDB()
	userInfo := request.SysUserInfo{
		ID: UUID,
	}
	var records []request.BilRecord
	lastWeek := "2022-09-01 00:00:00"
	today := "2022-10-30 23:59:59"

	db.Model(&userInfo).Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&records)

	for i := 0; i < len(records); i++ {

	}
	return nil, err

}
