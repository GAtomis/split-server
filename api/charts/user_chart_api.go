/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-09-26 16:55:44
 * @LastEditTime: 2022-10-15 18:46:02
 * @LastEditors: Gavin
 */

package charts

import (
	business "split-server/model/business/request"
	"split-server/utils"

	"github.com/gin-gonic/gin"
)

type USER_CHART_API struct {
}

func (uca *USER_CHART_API) GetRecordTypeByScope(ctx *gin.Context) {

	body := business.QuarterOrMonth{}
	ctx.ShouldBindQuery(&body)
	//获得当前访问用户ID
	jwt := new(utils.JWT)
	tokenInfo := jwt.GetUserInfo(ctx)
	id := tokenInfo.UserInfo.ID
	//数据操作
	dl := new(rbac_core.UserCharts)
	dl.GetRecordsBySeasons(id)
}
