// 自动生成模板Message
package wechat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Message 结构体
type Message struct {
      global.GVA_MODEL
      Cotent  string `json:"cotent" form:"cotent" gorm:"column:cotent;comment:信息内容;size:5000;"`
      Category  *int `json:"category" form:"category" gorm:"column:category;comment:信息属于公告还是提醒;"`
}


// TableName Message 表名
func (Message) TableName() string {
  return "wechat_message"
}

