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

type IdentityApi struct {
}

var identityService = service.ServiceGroupApp.WechatServiceGroup.IdentityService

// CreateIdentity 创建Identity
// @Tags Identity
// @Summary 创建Identity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wechat.Identity true "创建Identity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /identity/createIdentity [post]
func (identityApi *IdentityApi) CreateIdentity(c *gin.Context) {
	var identity wechat.Identity
	err := c.ShouldBindJSON(&identity)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":   {utils.NotEmpty()},
		"Status": {utils.NotEmpty()},
	}
	if err := utils.Verify(identity, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := identityService.CreateIdentity(identity); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIdentity 删除Identity
// @Tags Identity
// @Summary 删除Identity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wechat.Identity true "删除Identity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /identity/deleteIdentity [delete]
func (identityApi *IdentityApi) DeleteIdentity(c *gin.Context) {
	var identity wechat.Identity
	err := c.ShouldBindJSON(&identity)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := identityService.DeleteIdentity(identity); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIdentityByIds 批量删除Identity
// @Tags Identity
// @Summary 批量删除Identity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Identity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /identity/deleteIdentityByIds [delete]
func (identityApi *IdentityApi) DeleteIdentityByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := identityService.DeleteIdentityByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIdentity 更新Identity
// @Tags Identity
// @Summary 更新Identity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wechat.Identity true "更新Identity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /identity/updateIdentity [put]
func (identityApi *IdentityApi) UpdateIdentity(c *gin.Context) {
	var identity wechat.Identity
	err := c.ShouldBindJSON(&identity)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":   {utils.NotEmpty()},
		"Status": {utils.NotEmpty()},
	}
	if err := utils.Verify(identity, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := identityService.UpdateIdentity(identity); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIdentity 用id查询Identity
// @Tags Identity
// @Summary 用id查询Identity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wechat.Identity true "用id查询Identity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /identity/findIdentity [get]
func (identityApi *IdentityApi) FindIdentity(c *gin.Context) {
	var identity wechat.Identity
	err := c.ShouldBindQuery(&identity)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reidentity, err := identityService.GetIdentity(identity.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reidentity": reidentity}, c)
	}
}

// GetIdentityList 分页获取Identity列表
// @Tags Identity
// @Summary 分页获取Identity列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wechatReq.IdentitySearch true "分页获取Identity列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /identity/getIdentityList [get]
func (identityApi *IdentityApi) GetIdentityList(c *gin.Context) {
	var pageInfo wechatReq.IdentitySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := identityService.GetIdentityInfoList(pageInfo); err != nil {
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

func (identityApi *IdentityApi) List(c *gin.Context) {
	if list, err := identityService.GetIdentityList(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"list": list}, "获取成功", c)
	}
}
