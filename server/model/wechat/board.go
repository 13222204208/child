// 自动生成模板Board
package wechat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Board 结构体
type Board struct {
      global.GVA_MODEL
      Title  string `json:"title" form:"title" gorm:"column:title;comment:标题;"`
      Path  string `json:"path" form:"path" gorm:"column:path;comment:跳转的路径;"`
      Icon  string `json:"icon" form:"icon" gorm:"column:icon;comment:;"`
      Sort  *int `json:"sort" form:"sort" gorm:"column:sort;comment:排序按数字大小;"`
}


// TableName Board 表名
func (Board) TableName() string {
  return "wechat_board"
}

