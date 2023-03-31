package wechat

import (
	"strconv"

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

type EmergencyAlertApi struct {
}

var emergencyAlertService = service.ServiceGroupApp.WechatServiceGroup.EmergencyAlertService

// CreateEmergencyAlert 创建EmergencyAlert
// @Tags EmergencyAlert
// @Summary 创建EmergencyAlert
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wechat.EmergencyAlert true "创建EmergencyAlert"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /emergencyAlert/createEmergencyAlert [post]
func (emergencyAlertApi *EmergencyAlertApi) CreateEmergencyAlert(c *gin.Context) {
	var emergencyAlert wechat.EmergencyAlert
	err := c.ShouldBindJSON(&emergencyAlert)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"LostLocation": {utils.NotEmpty()},
		"LostTime":     {utils.NotEmpty()},
		"BabyId":       {utils.NotEmpty()},
	}
	if err := utils.Verify(emergencyAlert, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := emergencyAlertService.CreateEmergencyAlert(emergencyAlert); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteEmergencyAlert 删除EmergencyAlert
// @Tags EmergencyAlert
// @Summary 删除EmergencyAlert
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wechat.EmergencyAlert true "删除EmergencyAlert"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /emergencyAlert/deleteEmergencyAlert [delete]
func (emergencyAlertApi *EmergencyAlertApi) DeleteEmergencyAlert(c *gin.Context) {
	var emergencyAlert wechat.EmergencyAlert
	err := c.ShouldBindJSON(&emergencyAlert)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := emergencyAlertService.DeleteEmergencyAlert(emergencyAlert); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteEmergencyAlertByIds 批量删除EmergencyAlert
// @Tags EmergencyAlert
// @Summary 批量删除EmergencyAlert
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EmergencyAlert"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /emergencyAlert/deleteEmergencyAlertByIds [delete]
func (emergencyAlertApi *EmergencyAlertApi) DeleteEmergencyAlertByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := emergencyAlertService.DeleteEmergencyAlertByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateEmergencyAlert 更新EmergencyAlert
// @Tags EmergencyAlert
// @Summary 更新EmergencyAlert
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wechat.EmergencyAlert true "更新EmergencyAlert"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /emergencyAlert/updateEmergencyAlert [put]
func (emergencyAlertApi *EmergencyAlertApi) UpdateEmergencyAlert(c *gin.Context) {
	var emergencyAlert wechat.EmergencyAlert
	err := c.ShouldBindJSON(&emergencyAlert)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// verify := utils.Rules{
	// 	"LostLocation": {utils.NotEmpty()},
	// 	"LostTime":     {utils.NotEmpty()},
	// 	"BabyId":       {utils.NotEmpty()},
	// }
	// if err := utils.Verify(emergencyAlert, verify); err != nil {
	// 	response.FailWithMessage(err.Error(), c)
	// 	return
	// }
	if err := emergencyAlertService.UpdateEmergencyAlert(emergencyAlert); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindEmergencyAlert 用id查询EmergencyAlert
// @Tags EmergencyAlert
// @Summary 用id查询EmergencyAlert
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wechat.EmergencyAlert true "用id查询EmergencyAlert"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /emergencyAlert/findEmergencyAlert [get]
func (emergencyAlertApi *EmergencyAlertApi) FindEmergencyAlert(c *gin.Context) {
	var emergencyAlert wechat.EmergencyAlert
	err := c.ShouldBindQuery(&emergencyAlert)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reemergencyAlert, err := emergencyAlertService.GetEmergencyAlert(emergencyAlert.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reemergencyAlert": reemergencyAlert}, c)
	}
}

// GetEmergencyAlertList 分页获取EmergencyAlert列表
// @Tags EmergencyAlert
// @Summary 分页获取EmergencyAlert列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wechatReq.EmergencyAlertSearch true "分页获取EmergencyAlert列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /emergencyAlert/getEmergencyAlertList [get]
func (emergencyAlertApi *EmergencyAlertApi) GetEmergencyAlertList(c *gin.Context) {
	var pageInfo wechatReq.EmergencyAlertSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := emergencyAlertService.GetEmergencyAlertInfoList(pageInfo); err != nil {
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

func (emergencyAlertApi *EmergencyAlertApi) Issue(c *gin.Context) {
	var emergencyAlert wechat.EmergencyAlert
	err := c.ShouldBindJSON(&emergencyAlert)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"LostLocation": {utils.NotEmpty()},
		"LostTime":     {utils.NotEmpty()},
		"BabyId":       {utils.NotEmpty()},
	}
	if err := utils.Verify(emergencyAlert, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := c.MustGet("id").(uint)
	emergencyAlert.Uid = uid
	if err := emergencyAlertService.Issue(emergencyAlert); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("发布成功", c)
	}
}

// GetEmergencyAlertList api
func (emergencyAlertApi *EmergencyAlertApi) List(c *gin.Context) {
	if list, err := emergencyAlertService.GetEmergencyAlertList(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"list": list}, "获取成功", c)
	}
}

// findEmergencyAlertById
func (emergencyAlertApi *EmergencyAlertApi) FindEmergencyAlertById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.FailWithMessage("id不能为空", c)
		return
	}
	//id 转为int
	idInt, _ := strconv.Atoi(id)
	if reemergencyAlert, err := emergencyAlertService.GetEmergencyAlertById(idInt); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"info": reemergencyAlert}, c)
	}
}
