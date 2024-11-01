import service from '@/utils/request'
// @Tags Assettype
// @Summary 创建资产类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Assettype true "创建资产类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /AT/createAssettype [post]
export const createAssettype = (data) => {
  return service({
    url: '/AT/createAssettype',
    method: 'post',
    data
  })
}

// @Tags Assettype
// @Summary 删除资产类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Assettype true "删除资产类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /AT/deleteAssettype [delete]
export const deleteAssettype = (params) => {
  return service({
    url: '/AT/deleteAssettype',
    method: 'delete',
    params
  })
}

// @Tags Assettype
// @Summary 批量删除资产类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除资产类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /AT/deleteAssettype [delete]
export const deleteAssettypeByIds = (params) => {
  return service({
    url: '/AT/deleteAssettypeByIds',
    method: 'delete',
    params
  })
}

// @Tags Assettype
// @Summary 更新资产类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Assettype true "更新资产类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /AT/updateAssettype [put]
export const updateAssettype = (data) => {
  return service({
    url: '/AT/updateAssettype',
    method: 'put',
    data
  })
}

// @Tags Assettype
// @Summary 用id查询资产类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Assettype true "用id查询资产类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /AT/findAssettype [get]
export const findAssettype = (params) => {
  return service({
    url: '/AT/findAssettype',
    method: 'get',
    params
  })
}

// @Tags Assettype
// @Summary 分页获取资产类型列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取资产类型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /AT/getAssettypeList [get]
export const getAssettypeList = (params) => {
  return service({
    url: '/AT/getAssettypeList',
    method: 'get',
    params
  })
}

// @Tags Assettype
// @Summary 不需要鉴权的资产类型接口
// @accept application/json
// @Produce application/json
// @Param data query assetmReq.AssettypeSearch true "分页获取资产类型列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /AT/getAssettypePublic [get]
export const getAssettypePublic = () => {
  return service({
    url: '/AT/getAssettypePublic',
    method: 'get',
  })
}
