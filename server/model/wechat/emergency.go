// 自动生成模板EmergencyAlert
package wechat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// EmergencyAlert 结构体
type EmergencyAlert struct {
	global.GVA_MODEL
	Uid           uint   `json:"uid" form:"uid" gorm:"column:uid;comment:发布人用户ID;"`
	LostLocation  string `json:"lostLocation" form:"lostLocation" gorm:"column:lost_location;comment:走失位置;"`
	City          string `json:"city" form:"city" gorm:"column:city;comment:城市;"` // 城市
	LostTime      string `json:"lostTime" form:"lostTime" gorm:"column:lost_time;comment:走丢时间;"`
	Longitude     string `json:"longitude" form:"longitude" gorm:"column:longitude;comment:经度;"`
	Latitude      string `json:"latitude" form:"latitude" gorm:"column:latitude;comment:纬度;"`
	BabyId        *int   `json:"babyId" form:"babyId" gorm:"column:baby_id;comment:宝贝ID;"`
	ContactPerson string `json:"contactPerson" form:"contactPerson" gorm:"column:contact_person;comment:紧急联系人;size:500;"`
	OtherFeatures string `json:"otherFeatures" form:"otherFeatures" gorm:"column:other_features;comment:其它特征;size:1000;"`
	Status        *int   `json:"status" form:"status" gorm:"column:status;comment:0代表危险，1代表安全;"`
	Baby          Baby   `json:"baby" form:"baby" gorm:"foreignKey:baby_id;references:id;"` // 一对一关系
}

// TableName EmergencyAlert 表名
func (EmergencyAlert) TableName() string {
	return "wechat_emergency_alert"
}
