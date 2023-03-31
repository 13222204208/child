// 自动生成模板Contact
package wechat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Contact 结构体
type Contact struct {
	global.GVA_MODEL
	Uid              uint   `json:"uid" form:"uid" gorm:"column:uid;comment:发布人用户ID;"` // 一对一关系
	Name             string `json:"name" form:"name" gorm:"column:name;comment:;"`
	Phone            string `json:"phone" form:"phone" gorm:"column:phone;comment:;"`
	Message          string `json:"message" form:"message" gorm:"column:message;comment:留言;size:1000;"`
	EmergencyAlertId uint   `json:"emergencyAlertId" form:"emergencyAlertId" gorm:"column:emergency_alert_id;comment:紧急联系人ID;"` // 一对多关系
	Status           *int   `json:"status" form:"status" gorm:"column:status;comment:;"`
}

// TableName Contact 表名
func (Contact) TableName() string {
	return "wechat_contact"
}
