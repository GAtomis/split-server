/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-21 18:01:53
 * @LastEditTime: 2022-08-16 23:15:57
 * @LastEditors: Gavin
 */
package base_core

import (
	"fmt"
	reqRBAC "split-server/model/RBAC/request"
	reqBase "split-server/model/base/request"
	respBase "split-server/model/base/response"
	"split-server/utils"

	"gorm.io/gorm"
)

type Base struct {
}

func (b Base) OldLogin(body reqBase.Login) (*respBase.Token, error) {
	res, err := utils.InitSQL(func(db *gorm.DB) (any, error) {
		var sr reqRBAC.SysUserLogin
		if err := db.AutoMigrate(&sr); err != nil {
			return &body, err
		}
		err := db.Where("username = ? AND password = ?", body.Username, body.Password).First(&sr).Error
		fmt.Printf("sr: %v\n", sr)
		return &sr, err

	})
	if err != nil {
		return &respBase.Token{Token: "获取失败"}, err
	}
	if isMatch(&body, res.(*reqRBAC.SysUserLogin)) {
		jwt := new(utils.JWT)
		s, err2 := jwt.InitJWT(*res.(*reqRBAC.SysUserLogin))
		if err2 == nil {
			return &respBase.Token{
				Token: s,
			}, err
		}
	}
	return nil, err
}
func (b Base) Login(body reqBase.Login) (*respBase.Token, error) {

	db := utils.GAA_SQL.GetDB()

	var sr reqRBAC.SysUserLogin
	// if err := db.AutoMigrate(&sr); err != nil {
	// 	return nil, err
	// }
	err := db.Where("username = ? AND password = ?", body.Username, body.Password).First(&sr).Error
	fmt.Printf("sr: %v\n", sr)

	if err != nil {
		return &respBase.Token{Token: "获取失败"}, err
	}
	if isMatch(&body, &sr) {
		jwt := new(utils.JWT)
		s, err2 := jwt.InitJWT(sr)
		if err2 == nil {
			return &respBase.Token{
				Token: s,
			}, err
		}
	}
	return nil, err
}

func isMatch(target *reqBase.Login, source *reqRBAC.SysUserLogin) bool {
	return target.Password == source.Password
}
