/*
 * @Description: gin启动服务
 * @Author: Gavin
 * @Date: 2022-07-20 16:10:20
 * @LastEditTime: 2022-09-25 10:22:37
 * @LastEditors: Gavin
 */
package router

import (
	"split-server/config"
	"split-server/router/gateway"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	rg := r.Group(config.ROOT)
	r2 := new(gateway.Router)
	r2.InitCaptcha(rg)
	r2.InitTool(rg)
	r2.InitRoleRouter(rg)
	r2.InitUserRouter(rg)
	r2.InitPermissionRouter(rg)
	r2.InitComment(rg)
	r2.InitBill(rg)

	r.Run(":" + config.GetPort())
}
