// 自动生成模板Assettype
package assetm
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 资产类型 结构体  Assettype
type Assettype struct {
    global.GVA_MODEL
    TypeName  string `json:"typeName" form:"typeName" gorm:"column:type_name;comment:;" binding:"required"`  //类型名 
    Description  string `json:"description" form:"description" gorm:"column:description;comment:;"`  //描述 
    Category  string `json:"category" form:"category" gorm:"column:category;comment:;" binding:"required"`  //资产大类 
    Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:;"`  //备注 
}


// TableName 资产类型 Assettype自定义表名 asset_type
func (Assettype) TableName() string {
    return "asset_type"
}

