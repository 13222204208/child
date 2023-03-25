import service from '@/utils/request'

// @Tags Baby
// @Summary 创建Baby
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Baby true "创建Baby"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /baby/createBaby [post]
export const createBaby = (data) => {
  return service({
    url: '/baby/createBaby',
    method: 'post',
    data
  })
}

// @Tags Baby
// @Summary 删除Baby
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Baby true "删除Baby"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /baby/deleteBaby [delete]
export const deleteBaby = (data) => {
  return service({
    url: '/baby/deleteBaby',
    method: 'delete',
    data
  })
}

// @Tags Baby
// @Summary 删除Baby
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Baby"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /baby/deleteBaby [delete]
export const deleteBabyByIds = (data) => {
  return service({
    url: '/baby/deleteBabyByIds',
    method: 'delete',
    data
  })
}

// @Tags Baby
// @Summary 更新Baby
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Baby true "更新Baby"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /baby/updateBaby [put]
export const updateBaby = (data) => {
  return service({
    url: '/baby/updateBaby',
    method: 'put',
    data
  })
}

// @Tags Baby
// @Summary 用id查询Baby
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Baby true "用id查询Baby"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /baby/findBaby [get]
export const findBaby = (params) => {
  return service({
    url: '/baby/findBaby',
    method: 'get',
    params
  })
}

// @Tags Baby
// @Summary 分页获取Baby列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Baby列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /baby/getBabyList [get]
export const getBabyList = (params) => {
  return service({
    url: '/baby/getBabyList',
    method: 'get',
    params
  })
}
