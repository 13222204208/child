// 自动生成模板WechatUser
package wechat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// WechatUser 结构体
type WechatUser struct {
      global.GVA_MODEL
      Openid  string `json:"openid" form:"openid" gorm:"column:openid;comment:;"`
      Unionid  string `json:"unionid" form:"unionid" gorm:"column:unionid;comment:;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;"`
      Phone  string `json:"phone" form:"phone" gorm:"column:phone;comment:;"`
      Address  string `json:"address" form:"address" gorm:"column:address;comment:;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:;"`
}


// TableName WechatUser 表名
func (WechatUser) TableName() string {
  return "wechat_user"
}

