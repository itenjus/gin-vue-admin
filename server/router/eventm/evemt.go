package eventm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EventmtRouter struct {}

// InitEventmtRouter 初始化 事件管理 路由信息
func (s *EventmtRouter) InitEventmtRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	EMTRouter := Router.Group("EMT").Use(middleware.OperationRecord())
	EMTRouterWithoutRecord := Router.Group("EMT")
	EMTRouterWithoutAuth := PublicRouter.Group("EMT")
	{
		EMTRouter.POST("createEventmt", EMTApi.CreateEventmt)   // 新建事件管理
		EMTRouter.DELETE("deleteEventmt", EMTApi.DeleteEventmt) // 删除事件管理
		EMTRouter.DELETE("deleteEventmtByIds", EMTApi.DeleteEventmtByIds) // 批量删除事件管理
		EMTRouter.PUT("updateEventmt", EMTApi.UpdateEventmt)    // 更新事件管理
	}
	{
		EMTRouterWithoutRecord.GET("findEventmt", EMTApi.FindEventmt)        // 根据ID获取事件管理
		EMTRouterWithoutRecord.GET("getEventmtList", EMTApi.GetEventmtList)  // 获取事件管理列表
	}
	{
	    EMTRouterWithoutAuth.GET("getEventmtDataSource", EMTApi.GetEventmtDataSource)  // 获取事件管理数据源
	    EMTRouterWithoutAuth.GET("getEventmtPublic", EMTApi.GetEventmtPublic)  // 事件管理开放接口
	}
}
