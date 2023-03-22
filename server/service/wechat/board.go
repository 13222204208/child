/*
 * @Author: 13222204208 573516293@qq.com
 * @Date: 2023-03-22 14:57:03
 * @LastEditors: 13222204208 573516293@qq.com
 * @LastEditTime: 2023-03-22 15:10:09
 * @FilePath: /server/service/wechat/board.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package wechat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	wechatReq "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/request"
)

type BoardService struct {
}

// CreateBoard 创建Board记录
// Author [piexlmax](https://github.com/piexlmax)
func (boardService *BoardService) CreateBoard(board wechat.Board) (err error) {
	err = global.GVA_DB.Create(&board).Error
	return err
}

// DeleteBoard 删除Board记录
// Author [piexlmax](https://github.com/piexlmax)
func (boardService *BoardService) DeleteBoard(board wechat.Board) (err error) {
	err = global.GVA_DB.Delete(&board).Error
	return err
}

// DeleteBoardByIds 批量删除Board记录
// Author [piexlmax](https://github.com/piexlmax)
func (boardService *BoardService) DeleteBoardByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]wechat.Board{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateBoard 更新Board记录
// Author [piexlmax](https://github.com/piexlmax)
func (boardService *BoardService) UpdateBoard(board wechat.Board) (err error) {
	err = global.GVA_DB.Save(&board).Error
	return err
}

// GetBoard 根据id获取Board记录
// Author [piexlmax](https://github.com/piexlmax)
func (boardService *BoardService) GetBoard(id uint) (board wechat.Board, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&board).Error
	return
}

// GetBoardInfoList 分页获取Board记录
// Author [piexlmax](https://github.com/piexlmax)
func (boardService *BoardService) GetBoardInfoList(info wechatReq.BoardSearch) (list []wechat.Board, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&wechat.Board{})
	var boards []wechat.Board
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

	err = db.Limit(limit).Offset(offset).Find(&boards).Error
	return boards, total, err
}

// 获取Board记录 sort进行倒序
func (boardService *BoardService) GetBoardList() (boards []wechat.Board, err error) {
	err = global.GVA_DB.Order("sort desc").Find(&boards).Error
	return
}
