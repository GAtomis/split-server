/*
 * @Description: 注册
 * @Author: Gavin
 * @Date: 2022-08-17 16:47:01
 * @LastEditTime: 2022-09-16 13:04:28
 * @LastEditors: Gavin
 */
package base

import (
	"fmt"
	"split-server/api/RBAC"
	"split-server/model/base/request"
	"split-server/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Register(c *gin.Context) {

	var register request.Register
	err := c.ShouldBindBodyWith(&register, binding.JSON)

	if err != nil {
		utils.Fail(c)
		return
	}

	if utils.CaptchaVerify(c, register.Code) {
		fmt.Printf("register.Code: %v\n", register.Code)
		api := new(RBAC.USER_LOGIN_API)
		api.CreateUser(c)
		return
	}
	utils.FailDM("", "验证码错误", c)
	return

}
