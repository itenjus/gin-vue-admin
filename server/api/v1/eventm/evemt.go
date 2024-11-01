package eventm

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/eventm"
    eventmReq "github.com/flipped-aurora/gin-vue-admin/server/model/eventm/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type EventmtApi struct {}



// CreateEventmt 创建事件管理
// @Tags Eventmt
// @Summary 创建事件管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body eventm.Eventmt true "创建事件管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /EMT/createEventmt [post]
func (EMTApi *EventmtApi) CreateEventmt(c *gin.Context) {
	var EMT eventm.Eventmt
	err := c.ShouldBindJSON(&EMT)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    EMT.CreatedBy = utils.GetUserID(c)
	err = EMTService.CreateEventmt(&EMT)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteEventmt 删除事件管理
// @Tags Eventmt
// @Summary 删除事件管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body eventm.Eventmt true "删除事件管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /EMT/deleteEventmt [delete]
func (EMTApi *EventmtApi) DeleteEventmt(c *gin.Context) {
	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := EMTService.DeleteEventmt(ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteEventmtByIds 批量删除事件管理
// @Tags Eventmt
// @Summary 批量删除事件管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /EMT/deleteEventmtByIds [delete]
func (EMTApi *EventmtApi) DeleteEventmtByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := EMTService.DeleteEventmtByIds(IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateEventmt 更新事件管理
// @Tags Eventmt
// @Summary 更新事件管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body eventm.Eventmt true "更新事件管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /EMT/updateEventmt [put]
func (EMTApi *EventmtApi) UpdateEventmt(c *gin.Context) {
	var EMT eventm.Eventmt
	err := c.ShouldBindJSON(&EMT)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    EMT.UpdatedBy = utils.GetUserID(c)
	err = EMTService.UpdateEventmt(EMT)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindEventmt 用id查询事件管理
// @Tags Eventmt
// @Summary 用id查询事件管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query eventm.Eventmt true "用id查询事件管理"
// @Success 200 {object} response.Response{data=eventm.Eventmt,msg=string} "查询成功"
// @Router /EMT/findEventmt [get]
func (EMTApi *EventmtApi) FindEventmt(c *gin.Context) {
	ID := c.Query("ID")
	reEMT, err := EMTService.GetEventmt(ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reEMT, c)
}

// GetEventmtList 分页获取事件管理列表
// @Tags Eventmt
// @Summary 分页获取事件管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query eventmReq.EventmtSearch true "分页获取事件管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /EMT/getEventmtList [get]
func (EMTApi *EventmtApi) GetEventmtList(c *gin.Context) {
	var pageInfo eventmReq.EventmtSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := EMTService.GetEventmtInfoList(pageInfo)
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
// GetEventmtDataSource 获取Eventmt的数据源
// @Tags Eventmt
// @Summary 获取Eventmt的数据源
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /EMT/getEventmtDataSource [get]
func (EMTApi *EventmtApi) GetEventmtDataSource(c *gin.Context) {
    // 此接口为获取数据源定义的数据
    dataSource, err := EMTService.GetEventmtDataSource()
    if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
   		response.FailWithMessage("查询失败:" + err.Error(), c)
   		return
    }
   response.OkWithData(dataSource, c)
}

// GetEventmtPublic 不需要鉴权的事件管理接口
// @Tags Eventmt
// @Summary 不需要鉴权的事件管理接口
// @accept application/json
// @Produce application/json
// @Param data query eventmReq.EventmtSearch true "分页获取事件管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /EMT/getEventmtPublic [get]
func (EMTApi *EventmtApi) GetEventmtPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    EMTService.GetEventmtPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的事件管理接口信息",
    }, "获取成功", c)
}
