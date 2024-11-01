package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type VulnerSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    Title  string `json:"title" form:"title" `
    StartDatereported  *time.Time  `json:"startDatereported" form:"startDatereported"`
    EndDatereported  *time.Time  `json:"endDatereported" form:"endDatereported"`
    StartFinishdate  *time.Time  `json:"startFinishdate" form:"startFinishdate"`
    EndFinishdate  *time.Time  `json:"endFinishdate" form:"endFinishdate"`
    Risklevel  string `json:"risklevel" form:"risklevel" `
    Vulnerstatus  string `json:"vulnerstatus" form:"vulnerstatus" `
    Source  string `json:"source" form:"source" `
    Ops  string `json:"ops" form:"ops" `
    Org  string `json:"org" form:"org" `
    Assignto  string `json:"assignto" form:"assignto" `
    Assets  string `json:"assets" form:"assets" `
    request.PageInfo
}
