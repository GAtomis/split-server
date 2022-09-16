/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-08-18 23:07:43
 * @LastEditTime: 2022-09-15 15:28:37
 * @LastEditors: Gavin
 */
package request

import (
	"split-server/model/global"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type BilTable struct {
	ID uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	global.NoIDModel
	Name       string             `json:"name" gorm:"type:varchar(64);comment:标题; not null;"`
	Describe   string             `json:"describe" gorm:"type:varchar(255);comment:描述;"`
	StartTime  *global.LocalTime  `json:"startTime" gorm:"comment:开始时间;not null" `
	EndTime    *global.LocalTime  `json:"endTime" gorm:"comment:结束时间"`
	Area       string             `json:"area" gorm:"type:varchar(255);comment:地区;not null"`
	State      uint               `json:"state" gorm:"comment:状态 1运行时 2已结束 3已过期;default:1;not null"`
	BilRecords []BilRecord        `json:"bilRecords" gorm:"foreignKey:TableId;comment:账单记录;"`
	UserNum    uint               `json:"userNum" gorm:"comment:用户个数;not null"`
	SysUsers   []SysUserAssociate `json:"sysUsers" gorm:"many2many:bil_table_user;comment:绑定用户;"`
	CreatorId  uuid.UUID          `json:"creatorId" gorm:"comment:创建者ID"`
}

func (u *BilTable) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewV4()
	return
}

type BilRecord struct {
	ID uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	global.NoIDModel
	Price     string            `json:"price" gorm:"comment:价格;not null;"`
	Describe  string            `json:"describe" gorm:"type:varchar(255);comment:描述;"`
	StartTime *global.LocalTime `json:"startTime" gorm:"comment:开始时间;"`
	Existing  uint              `json:"existing" gorm:"comment:1 个人分摊 2均摊"`
	Type      uint              `json:"type" gorm:"comment:1 交通 2 餐饮 3日常 4住宿 5地点服务 6其他"`
	Area      string            `json:"area" gorm:"type:varchar(255);comment:地区;"`
	Img       string            `json:"img" gorm:"comment:封面"`
	CreatorId uuid.UUID         `json:"userId" gorm:"comment:创建者ID"`
	TableId   uuid.UUID         `json:"tableId" gorm:"comment:封面"`
}

func (u *BilRecord) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewV4()
	return
}
