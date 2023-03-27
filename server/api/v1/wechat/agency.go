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

type AgencyApi struct {
}

var agencyService = service.ServiceGroupApp.WechatServiceGroup.AgencyService

// CreateAgency 创建Agency
// @Tags Agency
// @Summary 创建Agency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wechat.Agency true "创建Agency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /agency/createAgency [post]
func (agencyApi *AgencyApi) CreateAgency(c *gin.Context) {
	var agency wechat.Agency
	err := c.ShouldBindJSON(&agency)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":  {utils.NotEmpty()},
		"Phone": {utils.NotEmpty()},
	}
	if err := utils.Verify(agency, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := agencyService.CreateAgency(agency); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAgency 删除Agency
// @Tags Agency
// @Summary 删除Agency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wechat.Agency true "删除Agency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /agency/deleteAgency [delete]
func (agencyApi *AgencyApi) DeleteAgency(c *gin.Context) {
	var agency wechat.Agency
	err := c.ShouldBindJSON(&agency)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := agencyService.DeleteAgency(agency); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAgencyByIds 批量删除Agency
// @Tags Agency
// @Summary 批量删除Agency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Agency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /agency/deleteAgencyByIds [delete]
func (agencyApi *AgencyApi) DeleteAgencyByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := agencyService.DeleteAgencyByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAgency 更新Agency
// @Tags Agency
// @Summary 更新Agency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wechat.Agency true "更新Agency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /agency/updateAgency [put]
func (agencyApi *AgencyApi) UpdateAgency(c *gin.Context) {
	var agency wechat.Agency
	err := c.ShouldBindJSON(&agency)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":  {utils.NotEmpty()},
		"Phone": {utils.NotEmpty()},
	}
	if err := utils.Verify(agency, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := agencyService.UpdateAgency(agency); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAgency 用id查询Agency
// @Tags Agency
// @Summary 用id查询Agency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wechat.Agency true "用id查询Agency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /agency/findAgency [get]
func (agencyApi *AgencyApi) FindAgency(c *gin.Context) {
	var agency wechat.Agency
	err := c.ShouldBindQuery(&agency)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reagency, err := agencyService.GetAgency(agency.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reagency": reagency}, c)
	}
}

// GetAgencyList 分页获取Agency列表
// @Tags Agency
// @Summary 分页获取Agency列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wechatReq.AgencySearch true "分页获取Agency列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /agency/getAgencyList [get]
func (agencyApi *AgencyApi) GetAgencyList(c *gin.Context) {
	var pageInfo wechatReq.AgencySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := agencyService.GetAgencyInfoList(pageInfo); err != nil {
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

func (agencyApi *AgencyApi) Apply(c *gin.Context) {
	var agency wechat.Agency
	err := c.ShouldBindJSON(&agency)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":  {utils.NotEmpty()},
		"Phone": {utils.NotEmpty()},
	}
	if err := utils.Verify(agency, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := agencyService.Apply(agency); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
