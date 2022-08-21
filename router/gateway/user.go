/*
 * @Description: user模块
 * @Author: Gavin
 * @Date: 2022-07-22 13:55:29
 * @LastEditTime: 2022-08-16 23:25:32
 * @LastEditors: Gavin
 */
package gateway

import (
	"split-server/api/RBAC"
	"split-server/router/interceptor"

	"github.com/gin-gonic/gin"
)

func (r *Router) InitUserRouter(g *gin.RouterGroup) {
	userGateway := g.Group("user").Use(interceptor.JWTAuth())
	userGateway.PUT("user", RBAC.UpdateUser)
	userGateway.GET("getUserInfo", RBAC.GetUserInfo)

}
