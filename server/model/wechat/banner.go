/*
 * @Author: 13222204208 573516293@qq.com
 * @Date: 2023-03-22 08:44:56
 * @LastEditors: 13222204208 573516293@qq.com
 * @LastEditTime: 2023-03-22 13:37:17
 * @FilePath: /server/model/wechat/banner.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
// 自动生成模板Banner
package wechat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Banner 结构体
type Banner struct {
	global.GVA_MODEL
	Image  string `json:"image" form:"image" gorm:"column:image;comment:图片地址;"`
	Path   string `json:"path" form:"path" gorm:"column:path;comment:跳转的路径;"`
	Sort   *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序按数字大小;"`
	Status *int   `json:"status" form:"status" gorm:"column:status;comment:显示的状态 0隐藏 1显示;"`
}

// TableName Banner 表名
func (Banner) TableName() string {
	return "wechat_banner"
}
