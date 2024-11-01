package eventm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/eventm"
    eventmReq "github.com/flipped-aurora/gin-vue-admin/server/model/eventm/request"
    "gorm.io/gorm"
)

type EventmtService struct {}
// CreateEventmt 创建事件管理记录
// Author [yourname](https://github.com/yourname)
func (EMTService *EventmtService) CreateEventmt(EMT *eventm.Eventmt) (err error) {
	err = global.GVA_DB.Create(EMT).Error
	return err
}

// DeleteEventmt 删除事件管理记录
// Author [yourname](https://github.com/yourname)
func (EMTService *EventmtService)DeleteEventmt(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&eventm.Eventmt{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&eventm.Eventmt{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteEventmtByIds 批量删除事件管理记录
// Author [yourname](https://github.com/yourname)
func (EMTService *EventmtService)DeleteEventmtByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&eventm.Eventmt{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&eventm.Eventmt{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateEventmt 更新事件管理记录
// Author [yourname](https://github.com/yourname)
func (EMTService *EventmtService)UpdateEventmt(EMT eventm.Eventmt) (err error) {
	err = global.GVA_DB.Model(&eventm.Eventmt{}).Where("id = ?",EMT.ID).Updates(&EMT).Error
	return err
}

// GetEventmt 根据ID获取事件管理记录
// Author [yourname](https://github.com/yourname)
func (EMTService *EventmtService)GetEventmt(ID string) (EMT eventm.Eventmt, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&EMT).Error
	return
}

// GetEventmtInfoList 分页获取事件管理记录
// Author [yourname](https://github.com/yourname)
func (EMTService *EventmtService)GetEventmtInfoList(info eventmReq.EventmtSearch) (list []eventm.Eventmt, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&eventm.Eventmt{})
    var EMTs []eventm.Eventmt
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
        if info.StartRepairttime != nil && info.EndRepairttime != nil {
            db = db.Where("repairt_time BETWEEN ? AND ? ",info.StartRepairttime,info.EndRepairttime)
        }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["event_time"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&EMTs).Error
	return  EMTs, total, err
}
func (EMTService *EventmtService)GetEventmtDataSource() (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)
	
	   handler := make([]map[string]any, 0)
	   
       
       global.GVA_DB.Table("sys_users").Select("nick_name as label,id as value").Scan(&handler)
	   res["handler"] = handler
	   reporter := make([]map[string]any, 0)
	   
       
       global.GVA_DB.Table("sys_users").Select("nick_name as label,id as value").Scan(&reporter)
	   res["reporter"] = reporter
	return
}
func (EMTService *EventmtService)GetEventmtPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
