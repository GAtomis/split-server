/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-09-20 11:14:50
 * @LastEditTime: 2022-09-20 13:17:01
 * @LastEditors: Gavin
 */
package bill

import (
	"split-server/model/global"
	"split-server/service/rbac_core"
	"split-server/utils"

	"github.com/gin-gonic/gin"
)

type BILL_RECROD_API struct {
}

func (b BILL_RECROD_API) GetBillRecrodList(ctx *gin.Context) {
	var PrimaryUUID global.PrimaryUUID
	_ = ctx.ShouldBindQuery(&PrimaryUUID)
	if err := utils.Verify(PrimaryUUID, utils.Primarykey); err != nil {
		utils.FailM(err.Error(), ctx)
		return
	}
	api := new(rbac_core.BilRecrod)
	if res, err := api.GetRecrodList(&PrimaryUUID); err != nil {

		utils.FailM(err.Error(), ctx)
		return

	} else {
		utils.OkDM(res, "操作成功", ctx)
	}

}
