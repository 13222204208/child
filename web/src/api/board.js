import service from '@/utils/request'

// @Tags Board
// @Summary 创建Board
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Board true "创建Board"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /board/createBoard [post]
export const createBoard = (data) => {
  return service({
    url: '/board/createBoard',
    method: 'post',
    data
  })
}

// @Tags Board
// @Summary 删除Board
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Board true "删除Board"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /board/deleteBoard [delete]
export const deleteBoard = (data) => {
  return service({
    url: '/board/deleteBoard',
    method: 'delete',
    data
  })
}

// @Tags Board
// @Summary 删除Board
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Board"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /board/deleteBoard [delete]
export const deleteBoardByIds = (data) => {
  return service({
    url: '/board/deleteBoardByIds',
    method: 'delete',
    data
  })
}

// @Tags Board
// @Summary 更新Board
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Board true "更新Board"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /board/updateBoard [put]
export const updateBoard = (data) => {
  return service({
    url: '/board/updateBoard',
    method: 'put',
    data
  })
}

// @Tags Board
// @Summary 用id查询Board
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Board true "用id查询Board"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /board/findBoard [get]
export const findBoard = (params) => {
  return service({
    url: '/board/findBoard',
    method: 'get',
    params
  })
}

// @Tags Board
// @Summary 分页获取Board列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Board列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /board/getBoardList [get]
export const getBoardList = (params) => {
  return service({
    url: '/board/getBoardList',
    method: 'get',
    params
  })
}
