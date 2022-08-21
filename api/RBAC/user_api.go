/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-22 16:00:03
 * @LastEditTime: 2022-08-17 18:26:09
 * @LastEditors: Gavin
 */
package RBAC

import (
	"fmt"
	"split-server/model/RBAC/request"
	"split-server/service/rbac_core"
	"split-server/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CreateUser(ctx *gin.Context) {

	var newUser request.SysUser

	err := ctx.ShouldBindBodyWith(&newUser, binding.JSON)
	if err != nil {
		utils.Fail(ctx)
		return
	}
	fmt.Printf("newUser: %v\n", newUser)
	api := new(rbac_core.User)
	res, err2 := api.CreateItem(newUser)
	if err2 != nil {
		utils.Fail(ctx)
		return
	}
	utils.OkDM(res.ID, "操作成功", ctx)

}

func GetUserInfo(ctx *gin.Context) {
	jwt := new(utils.JWT)
	tokenInfo := jwt.GetUserInfo(ctx)
	if tokenInfo != nil {
		utils.OkDM(tokenInfo.UserInfo, "操作成功", ctx)
	} else {
		utils.Fail(ctx)
	}

}
func UpdateUser(ctx *gin.Context) {
	jwt := new(utils.JWT)
	tokenInfo := jwt.GetUserInfo(ctx)
	var user request.SysUser
	ctx.ShouldBindJSON(&user)

	if tokenInfo != nil {
		api := new(rbac_core.User)
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
