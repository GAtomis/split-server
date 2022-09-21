/*
 * @Description: user模块
 * @Author: Gavin
 * @Date: 2022-07-22 13:55:29
 * @LastEditTime: 2022-09-19 18:06:59
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

	// loginApi := new(RBAC.USER_LOGIN_API)
	// userGateway.PUT("user", loginApi.UpdateUser)
	infoApi := new(RBAC.USER_INFO_API)
	userGateway.GET("userInfo", infoApi.GetUserInfo)
	userGateway.PUT("userInfo", infoApi.UpdateUser)
	userGateway.POST("userInfo", infoApi.CreateUser)
	userGateway.POST("addUserInfo", infoApi.CreateInfoByToken)
	userGateway.POST("getInfoListByName", infoApi.GetInfoListByName)
}
