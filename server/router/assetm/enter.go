package assetm

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	AssetsRouter
	AssettypeRouter
}

var (
	AMApi = api.ApiGroupApp.AssetmApiGroup.AssetsApi
	ATApi = api.ApiGroupApp.AssetmApiGroup.AssettypeApi
)
