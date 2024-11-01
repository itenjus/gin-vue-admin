package eventm

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	EventmtApi
}

var EMTService = service.ServiceGroupApp.EventmServiceGroup.EventmtService
