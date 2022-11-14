/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-27 22:17:49
 * @LastEditTime: 2022-11-14 23:55:26
 * @LastEditors: Gavin 850680822@qq.com
 */
package global

import (
	"database/sql/driver"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type DBModel struct {
	ID        string         `json:"id" gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `json:"createdAt" `
	UpdatedAt time.Time      `json:"updatedAt" `
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func (u *DBModel) BeforeCreate(tx *gorm.DB) (err error) {
	if !(len(u.ID) > 0) {
		u.ID = uuid.NewV4().String()
	}
	return
}

type PageInfo struct {
	Page     int `json:"page" form:"page"`         // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
}

// GetAuthorityId Get role by id structure
type Primarykey struct {
	ID string `json:"id" form:"id"` // 主键
}

// GetAuthorityId Get role by id structure

type PrimaryUUID struct {
	ID string `json:"id" form:"id"` // 主键

}

type SearchParams struct {
	Searchkey string `json:"searchkey" form:"searchkey"`
}

//	type UUID struct {
//		ID string `gorm:"primarykey; type:uuid;default:uuid_generate_v4()"`
//	}
type NoIDModel struct {
	CreatedAt time.Time      `json:"createdAt" `
	UpdatedAt time.Time      `json:"updatedAt" `
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

const TimeFormat = "2006-01-02 15:04:05"

type LocalTime time.Time

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 2 {
		*t = LocalTime(time.Time{})
		return
	}

	now, err := time.Parse(`"`+TimeFormat+`"`, string(data))
	*t = LocalTime(now)
	return
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

func (t LocalTime) Value() (driver.Value, error) {
	if t.String() == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return []byte(time.Time(t).Format(TimeFormat)), nil
}

func (t *LocalTime) Scan(v interface{}) error {
	tTime, _ := time.Parse("2006-01-02 15:04:05 +0800 CST", v.(time.Time).String())
	*t = LocalTime(tTime)
	return nil
}

func (t LocalTime) String() string {
	return time.Time(t).Format(TimeFormat)
}
