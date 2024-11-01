package assetm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AssetsRouter struct {}

// InitAssetsRouter 初始化 资产管理 路由信息
func (s *AssetsRouter) InitAssetsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	AMRouter := Router.Group("AM").Use(middleware.OperationRecord())
	AMRouterWithoutRecord := Router.Group("AM")
	AMRouterWithoutAuth := PublicRouter.Group("AM")
	{
		AMRouter.POST("createAssets", AMApi.CreateAssets)   // 新建资产管理
		AMRouter.DELETE("deleteAssets", AMApi.DeleteAssets) // 删除资产管理
		AMRouter.DELETE("deleteAssetsByIds", AMApi.DeleteAssetsByIds) // 批量删除资产管理
		AMRouter.PUT("updateAssets", AMApi.UpdateAssets)    // 更新资产管理
	}
	{
		AMRouterWithoutRecord.GET("findAssets", AMApi.FindAssets)        // 根据ID获取资产管理
		AMRouterWithoutRecord.GET("getAssetsList", AMApi.GetAssetsList)  // 获取资产管理列表
	}
	{
	    AMRouterWithoutAuth.GET("getAssetsPublic", AMApi.GetAssetsPublic)  // 资产管理开放接口
	}
}
