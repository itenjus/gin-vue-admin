import service from '@/utils/request'
// @Tags Eventmt
// @Summary 创建事件管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Eventmt true "创建事件管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /EMT/createEventmt [post]
export const createEventmt = (data) => {
  return service({
    url: '/EMT/createEventmt',
    method: 'post',
    data
  })
}

// @Tags Eventmt
// @Summary 删除事件管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Eventmt true "删除事件管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /EMT/deleteEventmt [delete]
export const deleteEventmt = (params) => {
  return service({
    url: '/EMT/deleteEventmt',
    method: 'delete',
    params
  })
}

// @Tags Eventmt
// @Summary 批量删除事件管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除事件管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /EMT/deleteEventmt [delete]
export const deleteEventmtByIds = (params) => {
  return service({
    url: '/EMT/deleteEventmtByIds',
    method: 'delete',
    params
  })
}

// @Tags Eventmt
// @Summary 更新事件管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Eventmt true "更新事件管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /EMT/updateEventmt [put]
export const updateEventmt = (data) => {
  return service({
    url: '/EMT/updateEventmt',
    method: 'put',
    data
  })
}

// @Tags Eventmt
// @Summary 用id查询事件管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Eventmt true "用id查询事件管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /EMT/findEventmt [get]
export const findEventmt = (params) => {
  return service({
    url: '/EMT/findEventmt',
    method: 'get',
    params
  })
}

// @Tags Eventmt
// @Summary 分页获取事件管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取事件管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /EMT/getEventmtList [get]
export const getEventmtList = (params) => {
  return service({
    url: '/EMT/getEventmtList',
    method: 'get',
    params
  })
}
// @Tags Eventmt
// @Summary 获取数据源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /EMT/findEventmtDataSource [get]
export const getEventmtDataSource = () => {
  return service({
    url: '/EMT/getEventmtDataSource',
    method: 'get',
  })
}

// @Tags Eventmt
// @Summary 不需要鉴权的事件管理接口
// @accept application/json
// @Produce application/json
// @Param data query eventmReq.EventmtSearch true "分页获取事件管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /EMT/getEventmtPublic [get]
export const getEventmtPublic = () => {
  return service({
    url: '/EMT/getEventmtPublic',
    method: 'get',
  })
}
