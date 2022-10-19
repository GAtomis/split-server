/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-09-20 11:14:50
 * @LastEditTime: 2022-10-03 13:41:46
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

type BILL_RECORD_API struct {
}

func (b BILL_RECORD_API) GetBillRecordList(ctx *gin.Context) {

	var body request.BilRecord
	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.FailM(err.Error(), ctx)
		return
	}

	var pageInfo global.PageInfo
	_ = ctx.ShouldBindQuery(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		utils.FailM(err.Error(), ctx)
		return
	}

	api := new(rbac_core.BilRecord)
	if res, err := api.GetRecordList(&body, &pageInfo); err != nil {

		utils.FailM(err.Error(), ctx)
		return

	} else {
		utils.OkDM(res, "操作成功", ctx)
	}

}

func (b BILL_RECORD_API) CreateItem(ctx *gin.Context) {
	//声明一个SysPermission
	var body request.BilRecord
	//成功JSON化
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		utils.FailDM(err.Error(), "参数验证", ctx)
		return
	}
	//载入api
	r2 := new(rbac_core.BilRecord)
	res, err2 := r2.CreateItem(&body)
	if err2 != nil {
		utils.FailDM(err2.Error(), "请求错误", ctx)
		return
	}
	utils.OkDM(res.ID, "操作成功", ctx)

}
func (b BILL_RECORD_API) UpdateItem(ctx *gin.Context) {
	//声明一个SysPermission
	var body request.BilRecord
	//成功JSON化
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		utils.FailDM(err.Error(), "参数验证", ctx)
		return
	}
	//载入api
	r2 := new(rbac_core.BilRecord)
	res, err2 := r2.UpdateItem(&body)
	if err2 != nil {
		utils.FailDM(err2.Error(), "请求错误", ctx)
		return
	}
	utils.OkDM(res.ID, "操作成功", ctx)

}
func (b BILL_RECORD_API) DeleteItem(ctx *gin.Context) {
	//声明一个SysPermission
	var body request.BilRecord
	//成功JSON化
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		utils.FailDM(err.Error(), "参数验证", ctx)
		return
	}
	//载入api
	r2 := new(rbac_core.BilRecord)
	res, err2 := r2.DeleteItem(&body)
	if err2 != nil {
		utils.FailDM(err2.Error(), "请求错误", ctx)
		return
	}
	utils.OkDM(res.ID, "操作成功", ctx)

}
