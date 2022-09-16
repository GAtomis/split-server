/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-21 18:23:47
 * @LastEditTime: 2022-09-14 14:27:17
 * @LastEditors: Gavin
 */
package response

import (
	"split-server/model/RBAC/request"

	"github.com/golang-jwt/jwt/v4"
)

//JTW模版
type MyClaims struct {
	UserInfo request.SysUserLogin
	jwt.StandardClaims
}
