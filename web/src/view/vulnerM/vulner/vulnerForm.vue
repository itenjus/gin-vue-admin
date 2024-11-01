
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="脆弱性名称:" prop="title">
          <el-input v-model="formData.title" :clearable="true"  placeholder="请输入脆弱性名称" />
       </el-form-item>
        <el-form-item label="脆弱性描述:" prop="desc">
          <el-input v-model="formData.desc" :clearable="true"  placeholder="请输入脆弱性描述" />
       </el-form-item>
        <el-form-item label="脆弱性上报日期:" prop="datereported">
          <el-date-picker v-model="formData.datereported" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="脆弱性处理的结束时间:" prop="finishdate">
          <el-date-picker v-model="formData.finishdate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="风险等级:" prop="risklevel">
           <el-select v-model="formData.risklevel" placeholder="请选择风险等级" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in risklevelOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="脆弱性状态:" prop="vulnerstatus">
           <el-select v-model="formData.vulnerstatus" placeholder="请选择脆弱性状态" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in eventstatusOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="脆弱性的图片:" prop="pic">
           <SelectImage v-model="formData.pic" multiple file-type="image"/>
       </el-form-item>
        <el-form-item label="脆弱性来源:" prop="source">
          <el-input v-model="formData.source" :clearable="true"  placeholder="请输入脆弱性来源" />
       </el-form-item>
        <el-form-item label="解决方案:" prop="solution">
          <el-input v-model="formData.solution" :clearable="true"  placeholder="请输入解决方案" />
       </el-form-item>
        <el-form-item label="处置过程:" prop="processrecord">
          <el-input v-model="formData.processrecord" :clearable="true"  placeholder="请输入处置过程" />
       </el-form-item>
        <el-form-item label="安全域:" prop="ops">
          <el-input v-model="formData.ops" :clearable="true"  placeholder="请输入安全域" />
       </el-form-item>
        <el-form-item label="org:" prop="org">
          <el-input v-model="formData.org" :clearable="true"  placeholder="请输入org" />
       </el-form-item>
        <el-form-item label="分配给:" prop="assignto">
          <el-input v-model="formData.assignto" :clearable="true"  placeholder="请输入分配给" />
       </el-form-item>
        <el-form-item label="资产:" prop="assets">
          <el-input v-model="formData.assets" :clearable="true"  placeholder="请输入资产" />
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
  createVulner,
  updateVulner,
  findVulner
} from '@/api/vulnerM/vulner'

defineOptions({
    name: 'VulnerForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'


const route = useRoute()
const router = useRouter()

const type = ref('')
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
               }],
               datereported : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               risklevel : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               vulnerstatus : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               assets : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findVulner({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    risklevelOptions.value = await getDictFunc('risklevel')
    eventstatusOptions.value = await getDictFunc('eventstatus')
}

init()
// 保存按钮
const save = async() => {
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
