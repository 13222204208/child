/*
 * @Author: 13222204208 573516293@qq.com
 * @Date: 2023-03-22 08:25:18
 * @LastEditors: 13222204208 573516293@qq.com
 * @LastEditTime: 2023-03-22 14:26:20
 * @FilePath: /server/service/wechat/menubar.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package wechat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	wechatReq "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/request"
)

type MenubarService struct {
}

// CreateMenubar 创建Menubar记录
// Author [piexlmax](https://github.com/piexlmax)
func (menubarService *MenubarService) CreateMenubar(menubar wechat.Menubar) (err error) {
	err = global.GVA_DB.Create(&menubar).Error
	return err
}

// DeleteMenubar 删除Menubar记录
// Author [piexlmax](https://github.com/piexlmax)
func (menubarService *MenubarService) DeleteMenubar(menubar wechat.Menubar) (err error) {
	err = global.GVA_DB.Delete(&menubar).Error
	return err
}

// DeleteMenubarByIds 批量删除Menubar记录
// Author [piexlmax](https://github.com/piexlmax)
func (menubarService *MenubarService) DeleteMenubarByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]wechat.Menubar{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateMenubar 更新Menubar记录
// Author [piexlmax](https://github.com/piexlmax)
func (menubarService *MenubarService) UpdateMenubar(menubar wechat.Menubar) (err error) {
	err = global.GVA_DB.Save(&menubar).Error
	return err
}

// GetMenubar 根据id获取Menubar记录
// Author [piexlmax](https://github.com/piexlmax)
func (menubarService *MenubarService) GetMenubar(id uint) (menubar wechat.Menubar, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&menubar).Error
	return
}

// GetMenubarInfoList 分页获取Menubar记录
// Author [piexlmax](https://github.com/piexlmax)
func (menubarService *MenubarService) GetMenubarInfoList(info wechatReq.MenubarSearch) (list []wechat.Menubar, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&wechat.Menubar{})
	var menubars []wechat.Menubar
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["sort"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	err = db.Limit(limit).Offset(offset).Find(&menubars).Error
	return menubars, total, err
}

// 获取菜单列表 status为1，sort 为倒序
func (menubarService *MenubarService) GetMenubarList() (list []wechat.Menubar, err error) {
	err = global.GVA_DB.Where("status = ?", 1).Order("sort desc").Find(&list).Error
	return
}
