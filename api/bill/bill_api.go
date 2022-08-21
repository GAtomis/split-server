/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-08-21 22:17:48
 * @LastEditTime: 2022-08-21 22:31:58
 * @LastEditors: Gavin
 */
package bill

import (
	"split-server/model/RBAC/request"
	"split-server/model/global"
	"split-server/service/rbac_core"
	"split-server/utils"

	"github.com/gin-gonic/gin"
)

type BILL_TABLE_API struct {
}

/**
 * @description: 查询列表
 * @param {*gin.Context} ctx
 * @return {*}
 * @Date: 2022-08-04 17:57:34
 */
func (api *BILL_TABLE_API) GetBillTableList(ctx *gin.Context) {
	var pageInfo global.PageInfo
	_ = ctx.ShouldBindQuery(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		utils.FailM(err.Error(), ctx)
		return
	}
	var permission request.SysPermission
	r2 := new(rbac_core.Permission)
	if res, err2 := r2.GetList(&permission, &pageInfo); err2 != nil {
		utils.FailDM(res, err2.Error(), ctx)
	} else {
		utils.OkDM(res, "success", ctx)
	}

}

/**
 * @description: 修改角色
 * @param {*gin.Context} ctx
 * @return {*}
 * @Date: 2022-08-04 18:02:51
 */
func (api *BILL_TABLE_API) UpdateBillTable(ctx *gin.Context) {
	// //声明一个SysPermission
	// var newBilTable request.BilTable
	// //成功JSON化
	// err := ctx.ShouldBindJSON(&newBilTable)
	// if err != nil {
	// 	utils.Fail(ctx)
	// 	return
	// }
	// // fmt.Printf("newSysPermission.SysPermissions: %v\n", newSysPermission.SysPermissions)
	// //载入api
	// fmt.Printf("%+v", newBilTable)
	// r2 := new(rbac_core.Permission)
	// res, err2 := r2.UpdateItem(&newBilTable)
	// if err2 != nil {
	// 	utils.Fail(ctx)
	// 	return
	// }
	// utils.OkDM(res.ID, "操作成功", ctx)

}

// @Tags AuthorityMenu
// @Summary 创建用户角色
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {object} response.Response{data=systemRes.SysMenusResponse,msg=string} "获取用户动态路由,返回包括系统菜单详情列表"
// @Router /menu/getMenu [post]
func (api *BILL_TABLE_API) CreateBillTable(ctx *gin.Context) {
	//声明一个SysPermission
	var newBilTable request.BilTable
	//成功JSON化
	err := ctx.ShouldBindJSON(&newBilTable)
	if err != nil {
		utils.Fail(ctx)
		return
	}
	//载入api
	r2 := new(rbac_core.BilTable)
	res, err2 := r2.CreateItem(&newBilTable)
	if err2 != nil {
		utils.Fail(ctx)
		return
	}
	utils.OkDM(res.ID, "操作成功", ctx)
}
func (api *BILL_TABLE_API) DeleteBillTable(ctx *gin.Context) {
	//声明一个SysPermission
	var newSysPermission request.SysPermission
	//成功JSON化
	err := ctx.ShouldBindJSON(&newSysPermission)
	if err != nil {
		utils.Fail(ctx)
		return
	}
	// fmt.Printf("newSysPermission.SysPermissions: %v\n", newSysPermission.SysPermissions)
	//载入api
	r2 := new(rbac_core.Permission)
	res, err2 := r2.DelItem(&newSysPermission)
	if err2 != nil {
		utils.Fail(ctx)
		return
	}
	utils.OkDM(res.ID, "操作成功", ctx)
}

func (api *BILL_TABLE_API) GetBillTable(ctx *gin.Context) {
	var Primarykey global.Primarykey
	_ = ctx.ShouldBindQuery(&Primarykey)
	if err := utils.Verify(Primarykey, utils.Primarykey); err != nil {
		utils.FailM(err.Error(), ctx)
		return
	}
	//载入api
	r2 := new(rbac_core.Permission)

	res, err2 := r2.GetItem(&Primarykey)
	if err2 != nil {
		utils.Fail(ctx)
		return
	}
	utils.OkDM(*res, "操作成功", ctx)
}
