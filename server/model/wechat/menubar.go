/*
 * @Author: 13222204208 573516293@qq.com
 * @Date: 2023-03-22 08:25:18
 * @LastEditors: 13222204208 573516293@qq.com
 * @LastEditTime: 2023-03-22 13:37:33
 * @FilePath: /server/model/wechat/menubar.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
// 自动生成模板Menubar
package wechat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Menubar 结构体
type Menubar struct {
	global.GVA_MODEL
	Title  string `json:"title" form:"title" gorm:"column:title;comment:标题;"`
	Icon   string `json:"icon" form:"icon" gorm:"column:icon;comment:图标地址;"`
	Path   string `json:"path" form:"path" gorm:"column:path;comment:跳转的路径;"`
	Sort   *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序按数字大小;"`
	Status *int   `json:"status" form:"status" gorm:"column:status;comment:显示的状态 0隐藏 1显示;"`
}

// TableName Menubar 表名
func (Menubar) TableName() string {
	return "wechat_menubar"
}
