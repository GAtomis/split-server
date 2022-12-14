/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-08-21 22:17:48
 * @LastEditTime: 2022-09-25 14:10:10
 * @LastEditors: Gavin
 */
package bill

import (
	"fmt"
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
	var BilTable request.BilTable
	r2 := new(rbac_core.BilTable)
	if res, err2 := r2.GetList(&BilTable, &pageInfo); err2 != nil {
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
	//声明一个SysPermission
	var newBilTable request.BilTable
	//成功JSON化
	err := ctx.ShouldBindJSON(&newBilTable)
	if err != nil {
		utils.FailM(err.Error(), ctx)
		return
	}
	fmt.Printf("newBilTable: %v\n", newBilTable)

	r2 := new(rbac_core.BilTable)

	res, err2 := r2.UpdateBilTableInfo(&newBilTable)
	if err2 != nil {
		utils.FailM(err2.Error(), ctx)
		return
	}

	utils.OkDM(res.ID, "操作成功", ctx)
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
		utils.FailDM(err.Error(), "参数验证", ctx)
		return
	}
	//载入api
	r2 := new(rbac_core.BilTable)
	res, err2 := r2.CreateItem(&newBilTable)
	if err2 != nil {
		utils.FailDM(err2.Error(), "请求错误", ctx)
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
	var PrimaryUUID global.PrimaryUUID
	_ = ctx.ShouldBindQuery(&PrimaryUUID)
	if err := utils.Verify(PrimaryUUID, utils.Primarykey); err != nil {
		utils.FailM(err.Error(), ctx)
		return
	}
	//载入api
	r2 := new(rbac_core.BilTable)

	res, err2 := r2.GetItem(&PrimaryUUID)
	if err2 != nil {
		utils.FailM(err2.Error(), ctx)
		return
	}
	utils.OkDM(*res, "操作成功", ctx)
}

/**
 * @description: 通过user的Id查询他参与关联的table列表
 * @param {*gin.Context} ctx
 * @return {*}
 * @Date: 2022-09-17 14:00:21
 */
func (api *BILL_TABLE_API) GetBillTableListByUserId(ctx *gin.Context) {
	var PrimaryUUID global.PrimaryUUID
	jwt := new(utils.JWT)
	tokenInfo := jwt.GetUserInfo(ctx)
	id := tokenInfo.UserInfo.ID
	PrimaryUUID.ID = id
	r2 := new(rbac_core.BilTable)

	var pageInfo global.PageInfo
	_ = ctx.ShouldBindQuery(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		utils.FailM(err.Error(), ctx)
		return
	}

	res, err2 := r2.GetBillTableListByUserId(&PrimaryUUID, &pageInfo)
	if err2 != nil {
		utils.FailM(err2.Error(), ctx)
		return
	}
	utils.OkDM(res, "success", ctx)

}
