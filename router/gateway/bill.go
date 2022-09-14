/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-08-21 22:14:38
 * @LastEditTime: 2022-09-13 22:42:34
 * @LastEditors: Gavin
 */
package gateway

import (
	"split-server/api/bill"
	"split-server/router/interceptor"

	"github.com/gin-gonic/gin"
)

func (r *Router) InitBill(g *gin.RouterGroup) {
	billGateway := g.Group("bill").Use(interceptor.MiddleCommon()).Use(interceptor.JWTAuth())
	bt := new(bill.BILL_TABLE_API)
	billGateway.POST("table", bt.CreateBillTable)
	billGateway.GET("table", bt.GetBillTable)
	billGateway.PUT("table", bt.UpdateBillTable)
	billGateway.DELETE("table", bt.DeleteBillTable)

}
