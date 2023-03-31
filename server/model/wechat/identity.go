// 自动生成模板Identity
package wechat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Identity 结构体
type Identity struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:;"`
}


// TableName Identity 表名
func (Identity) TableName() string {
  return "wechat_identity"
}

