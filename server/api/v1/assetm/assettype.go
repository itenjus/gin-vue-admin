package assetm

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/assetm"
    assetmReq "github.com/flipped-aurora/gin-vue-admin/server/model/assetm/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type AssettypeApi struct {}



// CreateAssettype 创建资产类型
// @Tags Assettype
// @Summary 创建资产类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body assetm.Assettype true "创建资产类型"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /AT/createAssettype [post]
func (ATApi *AssettypeApi) CreateAssettype(c *gin.Context) {
	var AT assetm.Assettype
	err := c.ShouldBindJSON(&AT)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = ATService.CreateAssettype(&AT)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteAssettype 删除资产类型
// @Tags Assettype
// @Summary 删除资产类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body assetm.Assettype true "删除资产类型"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /AT/deleteAssettype [delete]
func (ATApi *AssettypeApi) DeleteAssettype(c *gin.Context) {
	ID := c.Query("ID")
	err := ATService.DeleteAssettype(ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteAssettypeByIds 批量删除资产类型
// @Tags Assettype
// @Summary 批量删除资产类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /AT/deleteAssettypeByIds [delete]
func (ATApi *AssettypeApi) DeleteAssettypeByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := ATService.DeleteAssettypeByIds(IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateAssettype 更新资产类型
// @Tags Assettype
// @Summary 更新资产类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body assetm.Assettype true "更新资产类型"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /AT/updateAssettype [put]
func (ATApi *AssettypeApi) UpdateAssettype(c *gin.Context) {
	var AT assetm.Assettype
	err := c.ShouldBindJSON(&AT)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = ATService.UpdateAssettype(AT)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindAssettype 用id查询资产类型
// @Tags Assettype
// @Summary 用id查询资产类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query assetm.Assettype true "用id查询资产类型"
// @Success 200 {object} response.Response{data=assetm.Assettype,msg=string} "查询成功"
// @Router /AT/findAssettype [get]
func (ATApi *AssettypeApi) FindAssettype(c *gin.Context) {
	ID := c.Query("ID")
	reAT, err := ATService.GetAssettype(ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reAT, c)
}

// GetAssettypeList 分页获取资产类型列表
// @Tags Assettype
// @Summary 分页获取资产类型列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query assetmReq.AssettypeSearch true "分页获取资产类型列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /AT/getAssettypeList [get]
func (ATApi *AssettypeApi) GetAssettypeList(c *gin.Context) {
	var pageInfo assetmReq.AssettypeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := ATService.GetAssettypeInfoList(pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetAssettypePublic 不需要鉴权的资产类型接口
// @Tags Assettype
// @Summary 不需要鉴权的资产类型接口
// @accept application/json
// @Produce application/json
// @Param data query assetmReq.AssettypeSearch true "分页获取资产类型列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /AT/getAssettypePublic [get]
func (ATApi *AssettypeApi) GetAssettypePublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    ATService.GetAssettypePublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的资产类型接口信息",
    }, "获取成功", c)
}
