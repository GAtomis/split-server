/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-18 16:58:15
 * @LastEditTime: 2022-08-22 10:42:46
 * @LastEditors: Gavin
 */
package request

import "split-server/model/global"

type BizComment struct {
	global.DBModel
	Like        uint   `json:"like" gorm:"comment:点赞;"`
	Unlike      uint   `json:"unlike" gorm:"comment:点灭;"`
	Title       string `json:"title" gorm:"type:varchar(255);comment:留言标题;"`
	Content     string `json:"content" gorm:"type:varchar(255);comment:留言内容;"`
	UserId      uint   `json:"userId" gorm:"comment:关联用户;"`
	CreatorName string `json:"creatorName" gorm:"type:varchar(64);comment:创建者姓名;"`
	Avatar      string `json:"avatar" gorm:"type:varchar(255);comment:头像;"`
}
