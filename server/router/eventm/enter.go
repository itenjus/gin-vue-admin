package eventm

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	EventmtRouter
}

var EMTApi = api.ApiGroupApp.EventmApiGroup.EventmtApi
