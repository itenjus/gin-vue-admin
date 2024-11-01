// 自动生成模板Eventmt
package eventm

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 事件管理 结构体  Eventmt
type Eventmt struct {
	global.GVA_MODEL
	Eventtime   *time.Time `json:"eventTime" form:"eventTime" gorm:"column:event_time;comment:;" binding:"required"` //事件时间
	Event       string     `json:"event" form:"event" gorm:"column:event;comment:;" binding:"required"`              //事件
	Description string     `json:"description" form:"description" gorm:"column:description;comment:;type:text;"`     //详细说明
	Status      string     `json:"result" form:"result" gorm:"column:result;comment:;" binding:"required"`           //状态
	Handler     int        `json:"handler" form:"handler" gorm:"column:handler;comment:;"`                           //处理人
	Reporter    int        `json:"reporter" form:"reporter" gorm:"column:reporter;comment:;"`                        //报告人
	Type        string     `json:"eventtype" form:"eventtype" gorm:"column:event_type;comment:;" binding:"required"` //类型
	Repairttime *time.Time `json:"repairttime" form:"repairttime" gorm:"column:repairt_time;comment:;"`              //完成时间
	CreatedBy   uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 事件管理 Eventmt自定义表名 event_management
func (Eventmt) TableName() string {
	return "event_management"
}
