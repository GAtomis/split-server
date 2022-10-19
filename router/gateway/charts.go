/*
 * @Description: 图表
 * @Author: Gavin
 * @Date: 2022-09-26 16:43:19
 * @LastEditTime: 2022-10-15 18:43:00
 * @LastEditors: Gavin
 */
package gateway

import (
	"split-server/router/interceptor"

	"github.com/gin-gonic/gin"
)

func (r *Router) InitCharts(g *gin.RouterGroup) {
	api := g.Group("charts").Use(interceptor.MiddleCommon()).Use(interceptor.JWTAuth())

}
