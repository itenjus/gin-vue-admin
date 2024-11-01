// 自动生成模板Assets
package assetm
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 资产管理 结构体  Assets
type Assets struct {
    global.GVA_MODEL
    AssetId  *int `json:"assetId" form:"assetId" gorm:"primarykey;column:asset_id;comment:资产唯一标识;size:64;" binding:"required"`  //资产唯一标识 
    AssetName  string `json:"assetName" form:"assetName" gorm:"column:asset_name;comment:资产名称;"`  //资产名称 
    Category  string `json:"category" form:"category" gorm:"column:category;comment:资产类别;"`  //资产类别 
    Location  string `json:"location" form:"location" gorm:"column:location;comment:资产位置;"`  //资产位置 
    Status  string `json:"status" form:"status" gorm:"column:status;comment:资产状态;"`  //资产状态 
    Vendor  string `json:"vendor" form:"vendor" gorm:"column:vendor;comment:供应商;"`  //供应商 
    SerialNumber  string `json:"serialNumber" form:"serialNumber" gorm:"column:serial_number;comment:序列号;"`  //序列号 
    Warranty  string `json:"warranty" form:"warranty" gorm:"column:warranty;comment:保修信息;"`  //保修信息 
    Cost  *float64 `json:"cost" form:"cost" gorm:"column:cost;comment:资产成本;"`  //资产成本 
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 资产管理 Assets自定义表名 Assets
func (Assets) TableName() string {
    return "Assets"
}

