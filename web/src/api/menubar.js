import service from '@/utils/request'

// @Tags Menubar
// @Summary 创建Menubar
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Menubar true "创建Menubar"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menubar/createMenubar [post]
export const createMenubar = (data) => {
  return service({
    url: '/menubar/createMenubar',
    method: 'post',
    data
  })
}

// @Tags Menubar
// @Summary 删除Menubar
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Menubar true "删除Menubar"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /menubar/deleteMenubar [delete]
export const deleteMenubar = (data) => {
  return service({
    url: '/menubar/deleteMenubar',
    method: 'delete',
    data
  })
}

// @Tags Menubar
// @Summary 删除Menubar
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Menubar"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /menubar/deleteMenubar [delete]
export const deleteMenubarByIds = (data) => {
  return service({
    url: '/menubar/deleteMenubarByIds',
    method: 'delete',
    data
  })
}

// @Tags Menubar
// @Summary 更新Menubar
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Menubar true "更新Menubar"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /menubar/updateMenubar [put]
export const updateMenubar = (data) => {
  return service({
    url: '/menubar/updateMenubar',
    method: 'put',
    data
  })
}

// @Tags Menubar
// @Summary 用id查询Menubar
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Menubar true "用id查询Menubar"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /menubar/findMenubar [get]
export const findMenubar = (params) => {
  return service({
    url: '/menubar/findMenubar',
    method: 'get',
    params
  })
}

// @Tags Menubar
// @Summary 分页获取Menubar列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Menubar列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menubar/getMenubarList [get]
export const getMenubarList = (params) => {
  return service({
    url: '/menubar/getMenubarList',
    method: 'get',
    params
  })
}
