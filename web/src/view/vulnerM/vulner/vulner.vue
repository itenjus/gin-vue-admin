
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
      
        <el-form-item label="脆弱性名称" prop="title">
         <el-input v-model="searchInfo.title" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="脆弱性上报日期" prop="datereported">
            
            <template #label>
            <span>
              脆弱性上报日期
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
            <el-date-picker v-model="searchInfo.startDatereported" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endDatereported ? time.getTime() > searchInfo.endDatereported.getTime() : false"></el-date-picker>
            —
            <el-date-picker v-model="searchInfo.endDatereported" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startDatereported ? time.getTime() < searchInfo.startDatereported.getTime() : false"></el-date-picker>

        </el-form-item>
        <el-form-item label="脆弱性处理的结束时间" prop="finishdate">
            
            <template #label>
            <span>
              脆弱性处理的结束时间
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
            <el-date-picker v-model="searchInfo.startFinishdate" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endFinishdate ? time.getTime() > searchInfo.endFinishdate.getTime() : false"></el-date-picker>
            —
            <el-date-picker v-model="searchInfo.endFinishdate" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startFinishdate ? time.getTime() < searchInfo.startFinishdate.getTime() : false"></el-date-picker>

        </el-form-item>
           <el-form-item label="风险等级" prop="risklevel">
            <el-select v-model="searchInfo.risklevel" clearable placeholder="请选择" @clear="()=>{searchInfo.risklevel=undefined}">
              <el-option v-for="(item,key) in risklevelOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
           <el-form-item label="脆弱性状态" prop="vulnerstatus">
            <el-select v-model="searchInfo.vulnerstatus" clearable placeholder="请选择" @clear="()=>{searchInfo.vulnerstatus=undefined}">
              <el-option v-for="(item,key) in eventstatusOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item label="脆弱性来源" prop="source">
         <el-input v-model="searchInfo.source" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="安全域" prop="ops">
         <el-input v-model="searchInfo.ops" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="org" prop="org">
         <el-input v-model="searchInfo.org" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="分配给" prop="assignto">
         <el-input v-model="searchInfo.assignto" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="资产" prop="assets">
         <el-input v-model="searchInfo.assets" placeholder="搜索条件" />

        </el-form-item>

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button v-auth="btnAuth.add" type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-button v-auth="btnAuth.batchDelete" icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            <ExportTemplate v-auth="btnAuth.exportTemplate" template-id="vulnerM_Vulner" />
            <ExportExcel v-auth="btnAuth.exportExcel" template-id="vulnerM_Vulner" />
            <ImportExcel v-auth="btnAuth.importExcel" template-id="vulnerM_Vulner" @on-success="getTableData" />
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="日期" prop="createdAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
          <el-table-column align="left" label="脆弱性名称" prop="title" width="120" />
          <el-table-column align="left" label="脆弱性描述" prop="desc" width="120" />
         <el-table-column align="left" label="脆弱性上报日期" prop="datereported" width="180">
            <template #default="scope">{{ formatDate(scope.row.datereported) }}</template>
         </el-table-column>
         <el-table-column align="left" label="脆弱性处理的结束时间" prop="finishdate" width="180">
            <template #default="scope">{{ formatDate(scope.row.finishdate) }}</template>
         </el-table-column>
        <el-table-column align="left" label="风险等级" prop="risklevel" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.risklevel,risklevelOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="脆弱性状态" prop="vulnerstatus" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.vulnerstatus,eventstatusOptions) }}
            </template>
        </el-table-column>
           <el-table-column label="脆弱性的图片" prop="pic" width="200">
              <template #default="scope">
                 <div class="multiple-img-box">
                    <el-image preview-teleported v-for="(item,index) in scope.row.pic" :key="index" style="width: 80px; height: 80px" :src="getUrl(item)" fit="cover"/>
                </div>
              </template>
           </el-table-column>
          <el-table-column align="left" label="脆弱性来源" prop="source" width="120" />
          <el-table-column align="left" label="解决方案" prop="solution" width="120" />
          <el-table-column align="left" label="安全域" prop="ops" width="120" />
          <el-table-column align="left" label="org" prop="org" width="120" />
          <el-table-column align="left" label="分配给" prop="assignto" width="120" />
          <el-table-column align="left" label="资产" prop="assets" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button v-auth="btnAuth.info" type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看详情</el-button>
            <el-button v-auth="btnAuth.edit" type="primary" link icon="edit" class="table-button" @click="updateVulnerFunc(scope.row)">变更</el-button>
            <el-button v-auth="btnAuth.delete" type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'添加':'修改'}}</span>
                <div>
                  <el-button type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="脆弱性名称:"  prop="title" >
              <el-input v-model="formData.title" :clearable="true"  placeholder="请输入脆弱性名称" />
            </el-form-item>
            <el-form-item label="脆弱性描述:"  prop="desc" >
              <el-input v-model="formData.desc" :clearable="true"  placeholder="请输入脆弱性描述" />
            </el-form-item>
            <el-form-item label="脆弱性上报日期:"  prop="datereported" >
              <el-date-picker v-model="formData.datereported" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="脆弱性处理的结束时间:"  prop="finishdate" >
              <el-date-picker v-model="formData.finishdate" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="风险等级:"  prop="risklevel" >
              <el-select v-model="formData.risklevel" placeholder="请选择风险等级" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in risklevelOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="脆弱性状态:"  prop="vulnerstatus" >
              <el-select v-model="formData.vulnerstatus" placeholder="请选择脆弱性状态" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in eventstatusOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="脆弱性的图片:"  prop="pic" >
                <SelectImage
                 multiple
                 v-model="formData.pic"
                 file-type="image"
                 />
            </el-form-item>
            <el-form-item label="脆弱性来源:"  prop="source" >
              <el-input v-model="formData.source" :clearable="true"  placeholder="请输入脆弱性来源" />
            </el-form-item>
            <el-form-item label="解决方案:"  prop="solution" >
              <el-input v-model="formData.solution" :clearable="true"  placeholder="请输入解决方案" />
            </el-form-item>
            <el-form-item label="处置过程:"  prop="processrecord" >
              <el-input v-model="formData.processrecord" :clearable="true"  placeholder="请输入处置过程" />
            </el-form-item>
            <el-form-item label="安全域:"  prop="ops" >
              <el-input v-model="formData.ops" :clearable="true"  placeholder="请输入安全域" />
            </el-form-item>
            <el-form-item label="org:"  prop="org" >
              <el-input v-model="formData.org" :clearable="true"  placeholder="请输入org" />
            </el-form-item>
            <el-form-item label="分配给:"  prop="assignto" >
              <el-input v-model="formData.assignto" :clearable="true"  placeholder="请输入分配给" />
            </el-form-item>
            <el-form-item label="资产:"  prop="assets" >
              <el-input v-model="formData.assets" :clearable="true"  placeholder="请输入资产" />
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="脆弱性名称">
                        {{ detailFrom.title }}
                    </el-descriptions-item>
                    <el-descriptions-item label="脆弱性上报日期">
                        {{ detailFrom.datereported }}
                    </el-descriptions-item>
                    <el-descriptions-item label="脆弱性处理的结束时间">
                        {{ detailFrom.finishdate }}
                    </el-descriptions-item>
                    <el-descriptions-item label="风险等级">
                        {{ detailFrom.risklevel }}
                    </el-descriptions-item>
                    <el-descriptions-item label="脆弱性状态">
                        {{ detailFrom.vulnerstatus }}
                    </el-descriptions-item>
                    <el-descriptions-item label="脆弱性的图片">
                            <el-image style="width: 50px; height: 50px; margin-right: 10px" :preview-src-list="returnArrImg(detailFrom.pic)" :initial-index="index" v-for="(item,index) in detailFrom.pic" :key="index" :src="getUrl(item)" fit="cover" />
                    </el-descriptions-item>
                    <el-descriptions-item label="脆弱性来源">
                        {{ detailFrom.source }}
                    </el-descriptions-item>
                    <el-descriptions-item label="安全域">
                        {{ detailFrom.ops }}
                    </el-descriptions-item>
                    <el-descriptions-item label="org">
                        {{ detailFrom.org }}
                    </el-descriptions-item>
                    <el-descriptions-item label="分配给">
                        {{ detailFrom.assignto }}
                    </el-descriptions-item>
                    <el-descriptions-item label="资产">
                        {{ detailFrom.assets }}
                    </el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createVulner,
  deleteVulner,
  deleteVulnerByIds,
  updateVulner,
  findVulner,
  getVulnerList
} from '@/api/vulnerM/vulner'
import { getUrl } from '@/utils/image'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
// 引入按钮权限标识
import { useBtnAuth } from '@/utils/btnAuth'

// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'


defineOptions({
    name: 'Vulner'
})
// 按钮权限实例化
    const btnAuth = useBtnAuth()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const risklevelOptions = ref([])
const eventstatusOptions = ref([])
const formData = ref({
            title: '',
            desc: '',
            datereported: new Date(),
            finishdate: new Date(),
            risklevel: '',
            vulnerstatus: '',
            pic: [],
            source: '',
            solution: '',
            processrecord: '',
            ops: '',
            org: '',
            assignto: '',
            assets: '',
        })



// 验证规则
const rule = reactive({
               title : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               datereported : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               risklevel : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               vulnerstatus : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               assets : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
        datereported : [{ validator: (rule, value, callback) => {
        if (searchInfo.value.startDatereported && !searchInfo.value.endDatereported) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startDatereported && searchInfo.value.endDatereported) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startDatereported && searchInfo.value.endDatereported && (searchInfo.value.startDatereported.getTime() === searchInfo.value.endDatereported.getTime() || searchInfo.value.startDatereported.getTime() > searchInfo.value.endDatereported.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change' }],
        finishdate : [{ validator: (rule, value, callback) => {
        if (searchInfo.value.startFinishdate && !searchInfo.value.endFinishdate) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startFinishdate && searchInfo.value.endFinishdate) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startFinishdate && searchInfo.value.endFinishdate && (searchInfo.value.startFinishdate.getTime() === searchInfo.value.endFinishdate.getTime() || searchInfo.value.startFinishdate.getTime() > searchInfo.value.endFinishdate.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change' }],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getVulnerList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    risklevelOptions.value = await getDictFunc('risklevel')
    eventstatusOptions.value = await getDictFunc('eventstatus')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteVulnerFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteVulnerByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateVulnerFunc = async(row) => {
    const res = await findVulner({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteVulnerFunc = async (row) => {
    const res = await deleteVulner({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        title: '',
        desc: '',
        datereported: new Date(),
        finishdate: new Date(),
        risklevel: '',
        vulnerstatus: '',
        pic: [],
        source: '',
        solution: '',
        processrecord: '',
        ops: '',
        org: '',
        assignto: '',
        assets: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createVulner(formData.value)
                  break
                case 'update':
                  res = await updateVulner(formData.value)
                  break
                default:
                  res = await createVulner(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}


const detailFrom = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findVulner({ ID: row.ID })
  if (res.code === 0) {
    detailFrom.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}


</script>

<style>

</style>