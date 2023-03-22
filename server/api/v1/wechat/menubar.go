package wechat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	wechatReq "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MenubarApi struct {
}

var menubarService = service.ServiceGroupApp.WechatServiceGroup.MenubarService

// CreateMenubar 创建Menubar
// @Tags Menubar
// @Summary 创建Menubar
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wechat.Menubar true "创建Menubar"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menubar/createMenubar [post]
func (menubarApi *MenubarApi) CreateMenubar(c *gin.Context) {
	var menubar wechat.Menubar
	err := c.ShouldBindJSON(&menubar)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Title":  {utils.NotEmpty()},
		"Cion":   {utils.NotEmpty()},
		"Status": {utils.NotEmpty()},
	}
	if err := utils.Verify(menubar, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := menubarService.CreateMenubar(menubar); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMenubar 删除Menubar
// @Tags Menubar
// @Summary 删除Menubar
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wechat.Menubar true "删除Menubar"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /menubar/deleteMenubar [delete]
func (menubarApi *MenubarApi) DeleteMenubar(c *gin.Context) {
	var menubar wechat.Menubar
	err := c.ShouldBindJSON(&menubar)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := menubarService.DeleteMenubar(menubar); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMenubarByIds 批量删除Menubar
// @Tags Menubar
// @Summary 批量删除Menubar
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Menubar"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /menubar/deleteMenubarByIds [delete]
func (menubarApi *MenubarApi) DeleteMenubarByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := menubarService.DeleteMenubarByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMenubar 更新Menubar
// @Tags Menubar
// @Summary 更新Menubar
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wechat.Menubar true "更新Menubar"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /menubar/updateMenubar [put]
func (menubarApi *MenubarApi) UpdateMenubar(c *gin.Context) {
	var menubar wechat.Menubar
	err := c.ShouldBindJSON(&menubar)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Title":  {utils.NotEmpty()},
		"Cion":   {utils.NotEmpty()},
		"Status": {utils.NotEmpty()},
	}
	if err := utils.Verify(menubar, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := menubarService.UpdateMenubar(menubar); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMenubar 用id查询Menubar
// @Tags Menubar
// @Summary 用id查询Menubar
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wechat.Menubar true "用id查询Menubar"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /menubar/findMenubar [get]
func (menubarApi *MenubarApi) FindMenubar(c *gin.Context) {
	var menubar wechat.Menubar
	err := c.ShouldBindQuery(&menubar)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if remenubar, err := menubarService.GetMenubar(menubar.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remenubar": remenubar}, c)
	}
}

// GetMenubarList 分页获取Menubar列表
// @Tags Menubar
// @Summary 分页获取Menubar列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wechatReq.MenubarSearch true "分页获取Menubar列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menubar/getMenubarList [get]
func (menubarApi *MenubarApi) GetMenubarList(c *gin.Context) {
	var pageInfo wechatReq.MenubarSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := menubarService.GetMenubarInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetMenubarList
func (menubarApi *MenubarApi) GetList(c *gin.Context) {
	if list, err := menubarService.GetMenubarList(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"list": list}, c)
	}
}
