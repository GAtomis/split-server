/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-18 16:58:15
 * @LastEditTime: 2022-08-22 10:42:21
 * @LastEditors: Gavin
 */
package request

import (
	"split-server/model/business/request"
	"split-server/model/global"
)

type SysUser struct {
	global.DBModel
	Name        string               `json:"name" gorm:"type:varchar(32);comment:姓名;"`
	Username    string               `json:"username" gorm:"type:varchar(64);comment:姓名;"`
	Password    string               `json:"password" gorm:"type:varchar(32);comment:密码;"`
	Avatar      string               `json:"avatar" gorm:"type:varchar(255);comment:头像;default:https://avatars.githubusercontent.com/u/40788938?v=4;"`
	JobType     string               `json:"jobType" gorm:"type:varchar(32);comment:职位;"`
	Company     string               `json:"company" gorm:"type:varchar(32);comment:公司;"`
	CatchPhrase string               `json:"catchPhrase" gorm:"type:varchar(255);comment:个人简介;"`
	Salt        string               `json:"salt" gorm:"type:varchar(128);comment:盐;"`
	Gender      bool                 `json:"gender" gorm:"comment:性别;default:true"`
	Locked      bool                 `json:"locked" gorm:"comment:账号是否锁定，1：锁定，0未锁定;default:false"`
	RoleId      uint                 `json:"roleId" gorm:"comment:权限ID;default:1"`
	BilRecords  []BilRecord          `json:"bilRecords" gorm:"foreignKey:UserId;comment:评论条目;"`
	BizComments []request.BizComment `json:"bizComments" gorm:"foreignKey:UserId;comment:评论条目;"`
	BilTables   []BilTable           `json:"bilTables" gorm:"many2many:bil_table_user;comment:绑定评论;"`
}

type SysRole struct {
	global.DBModel
	Name           string          `json:"name" gorm:"type:varchar(128);comment:角色名;"`
	Available      bool            `json:"available" gorm:"comment:角色是否可用 1 Y ,0: N;default:true;"`
	Describe       string          `json:"describe" gorm:"type:varchar(255);comment:描述;"`
	Level          uint            `json:"level" gorm:"comment:级别 0-系统管理员 1-组带表;default:1"`
	SysUsers       []SysUser       `json:"sysUsers" gorm:"foreignKey:RoleId;comment:拥有用户;"`
	SysPermissions []SysPermission `json:"sysPermissions"  gorm:"many2many:sys_role_permission;comment:角色&权限;"`
}

type SysPermission struct {
	global.DBModel
	Name       string    `json:"name" gorm:"type:varchar(128);comment:资源名称;"`
	Type       string    `json:"type" gorm:"type:varchar(32);comment:资源类型：menu,button;"`
	Url        string    `json:"url" gorm:"type:varchar(128);comment:访问url地址;"`
	Percode    string    `json:"percode" gorm:"type:varchar(128);comment:权限代码字符串';"`
	Parentid   string    `json:"parentid" gorm:"type:varchar(20);comment:父结点id;"`
	Parentids  string    `json:"parentids" gorm:"type:varchar(128);comment:父结点id列表串;"`
	Sortstring string    `json:"sortstring" gorm:"type:varchar(128);comment:排序号;"`
	Available  bool      `json:"available" gorm:"comment:是否可用 1 Y ,0: N;default:true;"`
	SysRoles   []SysRole `json:"sysRoles" gorm:"many2many:sys_role_permission;comment:角色&权限;"`
}