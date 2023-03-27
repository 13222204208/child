// 自动生成模板Affirm
package wechat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Affirm 结构体
type Affirm struct {
	global.GVA_MODEL
	Title    string `json:"title" form:"title" gorm:"column:title;comment:标题;"`
	Icon     string `json:"icon" form:"icon" gorm:"column:icon;comment:图标地址;"`
	Contents string `json:"contents" form:"contents" gorm:"column:contents;comment:跳转的路径;size:1000;"`
	Sort     *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序按数字大小;"`
	Status   *int   `json:"status" form:"status" gorm:"column:status;comment:显示的状态 0隐藏 1显示;"`
}

// TableName Affirm 表名
func (Affirm) TableName() string {
	return "wechat_affirm"
}
