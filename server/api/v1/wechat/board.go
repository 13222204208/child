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

type BoardApi struct {
}

var boardService = service.ServiceGroupApp.WechatServiceGroup.BoardService

// CreateBoard 创建Board
// @Tags Board
// @Summary 创建Board
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wechat.Board true "创建Board"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /board/createBoard [post]
func (boardApi *BoardApi) CreateBoard(c *gin.Context) {
	var board wechat.Board
	err := c.ShouldBindJSON(&board)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Title": {utils.NotEmpty()},
		"Icon":  {utils.NotEmpty()},
	}
	if err := utils.Verify(board, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := boardService.CreateBoard(board); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBoard 删除Board
// @Tags Board
// @Summary 删除Board
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wechat.Board true "删除Board"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /board/deleteBoard [delete]
func (boardApi *BoardApi) DeleteBoard(c *gin.Context) {
	var board wechat.Board
	err := c.ShouldBindJSON(&board)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := boardService.DeleteBoard(board); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBoardByIds 批量删除Board
// @Tags Board
// @Summary 批量删除Board
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Board"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /board/deleteBoardByIds [delete]
func (boardApi *BoardApi) DeleteBoardByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := boardService.DeleteBoardByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBoard 更新Board
// @Tags Board
// @Summary 更新Board
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wechat.Board true "更新Board"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /board/updateBoard [put]
func (boardApi *BoardApi) UpdateBoard(c *gin.Context) {
	var board wechat.Board
	err := c.ShouldBindJSON(&board)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Title": {utils.NotEmpty()},
		"Icon":  {utils.NotEmpty()},
	}
	if err := utils.Verify(board, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := boardService.UpdateBoard(board); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBoard 用id查询Board
// @Tags Board
// @Summary 用id查询Board
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wechat.Board true "用id查询Board"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /board/findBoard [get]
func (boardApi *BoardApi) FindBoard(c *gin.Context) {
	var board wechat.Board
	err := c.ShouldBindQuery(&board)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reboard, err := boardService.GetBoard(board.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reboard": reboard}, c)
	}
}

// GetBoardList 分页获取Board列表
// @Tags Board
// @Summary 分页获取Board列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wechatReq.BoardSearch true "分页获取Board列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /board/getBoardList [get]
func (boardApi *BoardApi) GetBoardList(c *gin.Context) {
	var pageInfo wechatReq.BoardSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := boardService.GetBoardInfoList(pageInfo); err != nil {
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

// in api get board list
func (boardApi *BoardApi) GetList(c *gin.Context) {
	if list, err := boardService.GetBoardList(); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(list, c)
	}
}
