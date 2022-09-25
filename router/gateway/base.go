/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-21 15:57:48
 * @LastEditTime: 2022-09-24 12:09:47
 * @LastEditors: Gavin
 */

package gateway

import (
	"split-server/api/base"
	"split-server/router/interceptor"

	"github.com/gin-gonic/gin"
)

func (r *Router) InitCaptcha(g *gin.RouterGroup) {
	roleGateway := g.Group("base").Use(interceptor.Session("golang-tech-stack"))
	roleGateway.GET("code", base.GetCaptcha)
	roleGateway.POST("login", base.Login)
	roleGateway.POST("register", base.Register)

}
func (r *Router) InitComment(g *gin.RouterGroup) {
	roleGateway := g.Group("comment").Use(interceptor.MiddleCommon()).Use(interceptor.JWTAuth())
	ca := new(base.COMMENT_API)
	roleGateway.GET("getComments", ca.GetCommentsByUid)
	roleGateway.POST("comment", ca.CreateComment)
}
func (r *Router) InitTool(g *gin.RouterGroup) {
	baseRouter := g.Group("base").Use(interceptor.MiddleCommon()).Use(interceptor.JWTAuth())
	baseRouter.GET("getUUID", base.GetUUID)
}
