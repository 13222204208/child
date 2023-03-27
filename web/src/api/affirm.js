import service from '@/utils/request'

// @Tags Affirm
// @Summary 创建Affirm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Affirm true "创建Affirm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /affirm/createAffirm [post]
export const createAffirm = (data) => {
  return service({
    url: '/affirm/createAffirm',
    method: 'post',
    data
  })
}

// @Tags Affirm
// @Summary 删除Affirm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Affirm true "删除Affirm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /affirm/deleteAffirm [delete]
export const deleteAffirm = (data) => {
  return service({
    url: '/affirm/deleteAffirm',
    method: 'delete',
    data
  })
}

// @Tags Affirm
// @Summary 删除Affirm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Affirm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /affirm/deleteAffirm [delete]
export const deleteAffirmByIds = (data) => {
  return service({
    url: '/affirm/deleteAffirmByIds',
    method: 'delete',
    data
  })
}

// @Tags Affirm
// @Summary 更新Affirm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Affirm true "更新Affirm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /affirm/updateAffirm [put]
export const updateAffirm = (data) => {
  return service({
    url: '/affirm/updateAffirm',
    method: 'put',
    data
  })
}

// @Tags Affirm
// @Summary 用id查询Affirm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Affirm true "用id查询Affirm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /affirm/findAffirm [get]
export const findAffirm = (params) => {
  return service({
    url: '/affirm/findAffirm',
    method: 'get',
    params
  })
}

// @Tags Affirm
// @Summary 分页获取Affirm列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Affirm列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /affirm/getAffirmList [get]
export const getAffirmList = (params) => {
  return service({
    url: '/affirm/getAffirmList',
    method: 'get',
    params
  })
}
