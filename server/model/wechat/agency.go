// 自动生成模板Agency
package wechat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Agency 结构体
type Agency struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;"`
      Phone  string `json:"phone" form:"phone" gorm:"column:phone;comment:;"`
      Province  string `json:"province" form:"province" gorm:"column:province;comment:;"`
      City  string `json:"city" form:"city" gorm:"column:city;comment:;"`
}


// TableName Agency 表名
func (Agency) TableName() string {
  return "wechat_agency"
}

