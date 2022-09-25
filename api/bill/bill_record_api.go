/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-09-20 11:14:50
 * @LastEditTime: 2022-09-24 22:05:56
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
	var PrimaryUUID global.PrimaryUUID
	_ = ctx.ShouldBindQuery(&PrimaryUUID)
	if err := utils.Verify(PrimaryUUID, utils.Primarykey); err != nil {
		utils.FailM(err.Error(), ctx)
		return
	}
	api := new(rbac_core.BilRecord)
	if res, err := api.GetRecordList(&PrimaryUUID); err != nil {

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
