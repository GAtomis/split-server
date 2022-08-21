/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-08-18 23:07:43
 * @LastEditTime: 2022-08-21 22:40:02
 * @LastEditors: Gavin
 */
package request

import (
	"split-server/model/global"
	"time"
)

type BilTable struct {
	global.DBModel
	Name       string      `json:"name" gorm:"type:varchar(64);comment:描述;"`
	Describe   string      `json:"describe" gorm:"type:varchar(255);comment:描述;"`
	StartTime  time.Time   `json:"startTime" gorm:"comment:结束时间"`
	EndTime    time.Time   `json:"endTime" gorm:"comment:结束时间"`
	State      int         `json:"state" gorm:"comment:状态 0运行时 1已结束 2已过期;default:0;"`
	BilRecords []BilRecord `json:"bilRecords" gorm:"foreignKey:TableId;comment:账单记录;"`
	UserNum    int         `json:"userNum" gorm:"comment:用户个数;"`
	SysUsers   []SysUser   `json:"sysUsers" gorm:"many2many:bil_table_user;comment:绑定用户;"`
}
type BilRecord struct {
	global.DBModel
	Price     float32   `json:"price" gorm:"comment:价格;"`
	Describe  string    `json:"describe" gorm:"type:varchar(255);comment:描述;"`
	StartTime time.Time `json:"startTime" gorm:"comment:结束时间"`
	Existing  int       `json:"existing" gorm:"comment:0 个人分摊 1均摊"`
	Type      int       `json:"type" gorm:"comment:1 交通 2 餐饮 3日常 4住宿 5地点服务 6其他"`
	Img       string    `json:"img" gorm:"comment:封面"`
	UserId    uint      `json:"userId" gorm:"comment:归属用户"`
	TableId   uint      `json:"tableId" gorm:"comment:封面"`
}
