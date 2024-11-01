package assetm

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	AssetsApi
	AssettypeApi
}

var (
	AMService = service.ServiceGroupApp.AssetmServiceGroup.AssetsService
	ATService = service.ServiceGroupApp.AssetmServiceGroup.AssettypeService
)
