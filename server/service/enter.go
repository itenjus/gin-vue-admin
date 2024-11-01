package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/assetm"
	"github.com/flipped-aurora/gin-vue-admin/server/service/eventm"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/vulnerM"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	VulnerMServiceGroup vulnerM.ServiceGroup
	AssetmServiceGroup  assetm.ServiceGroup
	EventmServiceGroup  eventm.ServiceGroup
}
