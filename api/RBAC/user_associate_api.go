/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-09-15 11:43:43
 * @LastEditTime: 2022-09-15 11:52:19
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

type USER_ASSOCIATE_API struct {
}

func (u *USER_ASSOCIATE_API) CreateUser(ctx *gin.Context) {

	var newUser request.SysUserAssociate

	err := ctx.ShouldBindBodyWith(&newUser, binding.JSON)
	if err != nil {
		utils.Fail(ctx)
		return
	}
	fmt.Printf("newUser: %v\n", newUser)
	api := new(rbac_core.UserAssociate)
	res, err2 := api.CreateItem(newUser)
	if err2 != nil {
		utils.Fail(ctx)
		return
	}
	utils.OkDM(res.ID, "操作成功", ctx)

}

func (u *USER_ASSOCIATE_API) GetUserInfo(ctx *gin.Context) {
	jwt := new(utils.JWT)
	tokenInfo := jwt.GetUserInfo(ctx)
	if tokenInfo != nil {
		utils.OkDM(tokenInfo.UserInfo, "操作成功", ctx)
	} else {
		utils.Fail(ctx)
	}

}
func (u *USER_ASSOCIATE_API) UpdateUser(ctx *gin.Context) {
	jwt := new(utils.JWT)
	tokenInfo := jwt.GetUserInfo(ctx)
	var user request.SysUserAssociate
	ctx.ShouldBindJSON(&user)

	if tokenInfo != nil {
		api := new(rbac_core.UserAssociate)
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
