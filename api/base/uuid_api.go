/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-09-24 12:02:48
 * @LastEditTime: 2022-09-24 12:05:34
 * @LastEditors: Gavin
 */
package base

import (
	"split-server/utils"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func GetUUID(ctx *gin.Context) {
	s := uuid.NewV4().String()
	utils.OkDM(s, "生成UUID", ctx)
}
