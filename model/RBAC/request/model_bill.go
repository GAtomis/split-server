/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-08-18 23:07:43
 * @LastEditTime: 2022-11-15 00:02:35
 * @LastEditors: Gavin 850680822@qq.com
 */
package request

import (
	"split-server/model/global"
)

type BilTable struct {
	global.DBModel
	Name       string        `json:"name" gorm:"type:varchar(64);comment:标题; not null;"`
	Describe   string        `json:"describe" gorm:"type:varchar(255);comment:描述;"`
	StartTime  uint64        `json:"startTime" gorm:"comment:开始时间;" `
	EndTime    uint64        `json:"endTime" gorm:"comment:结束时间;not null"`
	Area       string        `json:"area" gorm:"type:varchar(255);comment:地区;not null"`
	State      uint          `json:"state" gorm:"comment:状态 1运行时 2已结束 3已过期;default:1;not null"`
	BilRecords []BilRecord   `json:"bilRecords" gorm:"foreignKey:TableId;comment:账单记录;"`
	UserNum    uint          `json:"userNum" gorm:"comment:用户个数;not null"`
	Total      string        `json:"total" gorm:"comment:价格;not null;"`
	SysUsers   []SysUserInfo `json:"sysUsers" gorm:"many2many:bil_table_user;comment:绑定用户;"`
	CreatorId  string        `json:"creatorId" gorm:"type:char(36);comment:创建者ID"`
	Creator    SysUserInfo   `json:"creator" gorm:"foreignKey:CreatorId"`
}

type BilRecord struct {
	global.DBModel
	Price     string      `json:"price" gorm:"comment:价格;not null;"`
	Describe  string      `json:"describe" gorm:"type:varchar(255);comment:描述;"`
	StartTime uint64      `json:"startTime" gorm:"comment:开始时间;" `
	EndTime   uint64      `json:"endTime" gorm:"comment:结束时间;not null"`
	Existing  uint        `json:"existing" gorm:"comment:1 个人分摊 2均摊"`
	Type      uint        `json:"type" gorm:"comment:1 交通 2 餐饮 3日常 4住宿 5地点服务 6其他"`
	Area      string      `json:"area" gorm:"type:varchar(255);comment:地区;"`
	Img       string      `json:"img" gorm:"comment:封面"`
	CreatorId string      `json:"creatorId" gorm:"type:char(36);comment:创建者ID"`
	Creator   SysUserInfo `json:"creator" gorm:"foreignKey:CreatorId"`
	TableId   string      `json:"tableId" gorm:"type:char(36);comment:封面"`
}
