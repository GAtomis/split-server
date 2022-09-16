/*
 * @Description: 用户信息表
 * @Author: Gavin
 * @Date: 2022-09-15 10:37:31
 * @LastEditTime: 2022-09-16 14:54:25
 * @LastEditors: Gavin
 */
package RBAC

import (
	"errors"
	"fmt"
	"split-server/model/RBAC/request"
	"split-server/model/global"
	"split-server/service/rbac_core"
	"split-server/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

type USER_INFO_API struct {
}

func (u *USER_INFO_API) CreateUser(ctx *gin.Context) {

	var newUser request.SysUserInfo

	err := ctx.ShouldBindBodyWith(&newUser, binding.JSON)
	if err != nil {
		utils.Fail(ctx)
		return
	}
	fmt.Printf("newUser: %v\n", newUser)
	api := new(rbac_core.UserInfo)
	res, err2 := api.CreateItem(newUser)
	if err2 != nil {
		utils.Fail(ctx)
		return
	}
	utils.OkDM(res.ID, "操作成功", ctx)

}

func (u *USER_INFO_API) GetUserInfo(ctx *gin.Context) {
	jwt := new(utils.JWT)
	tokenInfo := jwt.GetUserInfo(ctx)
	id := tokenInfo.UserInfo.ID
	api := new(rbac_core.UserInfo)
	if res, err2 := api.GetItemById(&global.PrimaryUUID{ID: id.String()}); err2 != nil {

		utils.FailDM(err2.Error(), "操作异常", ctx)
	} else {
		utils.OkDM(res, "操作成功", ctx)
	}

	// if tokenInfo != nil {
	// 	utils.OkDM(tokenInfo.UserInfo, "操作成功", ctx)
	// } else {
	// 	utils.Fail(ctx)
	// }

}
func (u *USER_INFO_API) UpdateUser(ctx *gin.Context) {
	jwt := new(utils.JWT)
	tokenInfo := jwt.GetUserInfo(ctx)
	var user request.SysUserInfo
	ctx.ShouldBindJSON(&user)

	if tokenInfo != nil {
		api := new(rbac_core.UserInfo)
		res, err2 := api.UpdateItem(user)
		if err2 != nil {
			utils.Fail(ctx)
			return
		}
		utils.OkDM(res.ID, "操作成功", ctx)
	} else {
		utils.Fail(ctx)
	}

}

func (u *USER_INFO_API) CreateInfoByToken(ctx *gin.Context) {

	jwt := new(utils.JWT)
	tokenInfo := jwt.GetUserInfo(ctx)
	id := tokenInfo.UserInfo.ID
	api := new(rbac_core.UserInfo)
	api.GetItemById(&global.PrimaryUUID{ID: id.String()})

	if _, err := api.GetItemById(&global.PrimaryUUID{ID: id.String()}); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			var user request.SysUserInfo
			ctx.ShouldBindJSON(&user)
			user.ID = tokenInfo.UserInfo.ID
			if res, err2 := api.CreateItem(user); err2 != nil {

				utils.FailM(err2.Error(), ctx)

			} else {
				utils.OkDM(res.ID, "创建成功", ctx)
			}

		} else {
			utils.Fail(ctx)

		}
	} else {
		utils.FailM("当前用户非新用户", ctx)
	}

}
