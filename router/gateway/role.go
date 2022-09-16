/*
 * @Description: 路由层 RBAC
 * @Author: Gavin
 * @Date: 2022-07-19 16:30:11
 * @LastEditTime: 2022-09-15 00:14:26
 * @LastEditors: Gavin
 */
package gateway

import (
	"split-server/api/RBAC"
	"split-server/router/interceptor"

	"github.com/gin-gonic/gin"
)

type Router struct{}

func (r *Router) InitRoleRouter(g *gin.RouterGroup) {
	roleGateway := g.Group("role").Use(interceptor.MiddleCommon()).Use(interceptor.JWTAuth())
	api := new(RBAC.ROLE_API)
	roleGateway.POST("role", api.CreateRole)
	roleGateway.GET("role", api.GetRole)
	roleGateway.PUT("role", api.UpdateRole)
	roleGateway.GET("list", api.GetRoleList)

}
