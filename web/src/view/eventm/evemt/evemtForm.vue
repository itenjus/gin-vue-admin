
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="事件时间:" prop="eventTime">
          <el-date-picker v-model="formData.eventTime" type="date" placeholder="选择日期" :clearable="false"></el-date-picker>
       </el-form-item>
        <el-form-item label="事件:" prop="event">
          <el-input v-model="formData.event" :clearable="true"  placeholder="请输入事件" />
       </el-form-item>
        <el-form-item label="详细说明:" prop="description">
          <RichEdit v-model="formData.description"/>
       </el-form-item>
        <el-form-item label="状态:" prop="result">
           <el-select v-model="formData.result" placeholder="请选择状态" style="width:100%" :clearable="false" >
              <el-option v-for="(item,key) in eventstatusOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="处理人:" prop="handler">
        <el-select  v-model="formData.handler" placeholder="请选择处理人" style="width:100%" :clearable="false" >
          <el-option v-for="(item,key) in dataSource.handler" :key="key" :label="item.label" :value="item.value" />
        </el-select>
       </el-form-item>
        <el-form-item label="报告人:" prop="reporter">
        <el-select  v-model="formData.reporter" placeholder="请选择报告人" style="width:100%" :clearable="false" >
          <el-option v-for="(item,key) in dataSource.reporter" :key="key" :label="item.label" :value="item.value" />
        </el-select>
       </el-form-item>
        <el-form-item label="类型:" prop="eventtype">
           <el-select v-model="formData.eventtype" placeholder="请选择类型" style="width:100%" :clearable="false" >
              <el-option v-for="(item,key) in event_typeOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="完成时间:" prop="repairttime">
          <el-date-picker v-model="formData.repairttime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
    getEventmtDataSource,
  createEventmt,
  updateEventmt,
  findEventmt
} from '@/api/eventm/evemt'

defineOptions({
    name: 'EventmtForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'


const route = useRoute()
const router = useRouter()

const type = ref('')
const eventstatusOptions = ref([])
const event_typeOptions = ref([])
const formData = ref({
            eventTime: new Date(),
            event: '',
            description: '',
            result: '未进行',
            handler: '',
            reporter: '',
            eventtype: '',
            repairttime: new Date(),
        })
// 验证规则
const rule = reactive({
               eventTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               event : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               result : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               eventtype : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getEventmtDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findEventmt({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create';
    }
    eventstatusOptions.value = await getDictFunc('eventstatus')
    event_typeOptions.value = await getDictFunc('event_type')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createEventmt(formData.value)
               break
             case 'update':
               res = await updateEventmt(formData.value)
               break
             default:
               res = await createEventmt(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
