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

type AffirmApi struct {
}

var affirmService = service.ServiceGroupApp.WechatServiceGroup.AffirmService

// CreateAffirm 创建Affirm
// @Tags Affirm
// @Summary 创建Affirm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wechat.Affirm true "创建Affirm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /affirm/createAffirm [post]
func (affirmApi *AffirmApi) CreateAffirm(c *gin.Context) {
	var affirm wechat.Affirm
	err := c.ShouldBindJSON(&affirm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Title":    {utils.NotEmpty()},
		"Cion":     {utils.NotEmpty()},
		"Contents": {utils.NotEmpty()},
		"Status":   {utils.NotEmpty()},
	}
	if err := utils.Verify(affirm, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := affirmService.CreateAffirm(affirm); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAffirm 删除Affirm
// @Tags Affirm
// @Summary 删除Affirm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wechat.Affirm true "删除Affirm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /affirm/deleteAffirm [delete]
func (affirmApi *AffirmApi) DeleteAffirm(c *gin.Context) {
	var affirm wechat.Affirm
	err := c.ShouldBindJSON(&affirm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := affirmService.DeleteAffirm(affirm); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAffirmByIds 批量删除Affirm
// @Tags Affirm
// @Summary 批量删除Affirm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Affirm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /affirm/deleteAffirmByIds [delete]
func (affirmApi *AffirmApi) DeleteAffirmByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := affirmService.DeleteAffirmByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAffirm 更新Affirm
// @Tags Affirm
// @Summary 更新Affirm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wechat.Affirm true "更新Affirm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /affirm/updateAffirm [put]
func (affirmApi *AffirmApi) UpdateAffirm(c *gin.Context) {
	var affirm wechat.Affirm
	err := c.ShouldBindJSON(&affirm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Title":    {utils.NotEmpty()},
		"Cion":     {utils.NotEmpty()},
		"Contents": {utils.NotEmpty()},
		"Status":   {utils.NotEmpty()},
	}
	if err := utils.Verify(affirm, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := affirmService.UpdateAffirm(affirm); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAffirm 用id查询Affirm
// @Tags Affirm
// @Summary 用id查询Affirm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wechat.Affirm true "用id查询Affirm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /affirm/findAffirm [get]
func (affirmApi *AffirmApi) FindAffirm(c *gin.Context) {
	var affirm wechat.Affirm
	err := c.ShouldBindQuery(&affirm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reaffirm, err := affirmService.GetAffirm(affirm.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reaffirm": reaffirm}, c)
	}
}

// GetAffirmList 分页获取Affirm列表
// @Tags Affirm
// @Summary 分页获取Affirm列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wechatReq.AffirmSearch true "分页获取Affirm列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /affirm/getAffirmList [get]
func (affirmApi *AffirmApi) GetAffirmList(c *gin.Context) {
	var pageInfo wechatReq.AffirmSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := affirmService.GetAffirmInfoList(pageInfo); err != nil {
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

func (affirmApi *AffirmApi) List(c *gin.Context) {

	if list, err := affirmService.GetAffirmList(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"list": list}, "获取成功", c)
	}
}
