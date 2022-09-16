/*
 * @Description: api层 RBAC
 * @Author: Gavin
 * @Date: 2022-07-19 17:56:36
 * @LastEditTime: 2022-09-15 15:05:57
 * @LastEditors: Gavin
 */
package RBAC

import (
	"split-server/model/RBAC/request"
	"split-server/model/global"
	"split-server/service/rbac_core"
	"split-server/utils"

	"github.com/gin-gonic/gin"
)

type ROLE_API struct {
}

/**
 * @description: 修改角色
 * @param {*gin.Context} ctx
 * @return {*}
 * @Date: 2022-08-04 18:02:51
 */
func (r *ROLE_API) UpdateRole(ctx *gin.Context) {
	//声明一个SysRole
	var newSysRole request.SysRole
	//成功JSON化
	err := ctx.ShouldBindJSON(&newSysRole)
	if err != nil {
		utils.Fail(ctx)
		return
	}
	// fmt.Printf("newSysRole.SysPermissions: %v\n", newSysRole.SysPermissions)
	//载入api
	r2 := new(rbac_core.Role)
	res, err2 := r2.UpdateItem(&newSysRole)
	if err2 != nil {
		utils.Fail(ctx)
		return
	}
	utils.OkDM(res.ID, "操作成功", ctx)

}

// @Tags AuthorityMenu
// @Summary 创建用户角色
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {object} response.Response{data=systemRes.SysMenusResponse,msg=string} "获取用户动态路由,返回包括系统菜单详情列表"
// @Router /menu/getMenu [post]
func (r *ROLE_API) CreateRole(ctx *gin.Context) {
	//声明一个SysRole
	var newSysRole request.SysRole
	//成功JSON化
	err := ctx.ShouldBindJSON(&newSysRole)
	if err != nil {
		utils.FailM(err.Error(), ctx)
		return
	}
	//载入api
	r2 := new(rbac_core.Role)
	res, err2 := r2.CreateItem(newSysRole)
	if err2 != nil {
		utils.FailM(err2.Error(), ctx)
		return
	}
	utils.OkDM(res.ID, "操作成功", ctx)
}

// @Tags AuthorityMenu
// @Summary 获得角色权限
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {object} response.Response{data=systemRes.SysMenusResponse,msg=string} "获取用户动态路由,返回包括系统菜单详情列表"
// @Router /menu/getMenu [post]
func (r *ROLE_API) GetRole(ctx *gin.Context) {

	var Primarykey global.Primarykey

	_ = ctx.ShouldBindQuery(&Primarykey)

	if err := utils.Verify(Primarykey, utils.Primarykey); err != nil {
		utils.FailM(err.Error(), ctx)
		return
	}
	//载入api
	r2 := new(rbac_core.Role)

	res, err2 := r2.GetItem(&Primarykey)
	if err2 != nil {
		utils.Fail(ctx)
		return
	}
	utils.OkDM(res, "操作成功", ctx)
}

/**
 * @description: 、获得所有角色列表
 * @param {*gin.Context} ctx
 * @return {*}
 * @Date: 2022-08-04 17:57:34
 */
func (r *ROLE_API) GetRoleList(ctx *gin.Context) {
	var pageInfo global.PageInfo
	_ = ctx.ShouldBindQuery(&pageInfo)

	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		utils.FailM(err.Error(), ctx)
		return
	}

	jwt := new(utils.JWT)
	userInfo := jwt.GetUserInfo(ctx).UserInfo

	Primarykey := global.Primarykey{
		ID: userInfo.RoleId,
	}

	r2 := new(rbac_core.Role)
	role, err2 := r2.GetItem(&Primarykey)
	if err2 != nil {
		utils.FailM("getRole fail"+err2.Error(), ctx)
		return
	}
	if res, err2 := r2.GetRoleList(&role, &pageInfo); err2 != nil {
		utils.FailDM(res, err2.Error(), ctx)
	} else {
		utils.OkDM(res, "success", ctx)
	}

}
