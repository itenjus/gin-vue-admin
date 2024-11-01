
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="资产唯一标识:" prop="assetId">
          <el-input v-model.number="formData.assetId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="资产名称:" prop="assetName">
          <el-input v-model="formData.assetName" :clearable="true"  placeholder="请输入资产名称" />
       </el-form-item>
        <el-form-item label="资产类别:" prop="category">
          <el-input v-model="formData.category" :clearable="true"  placeholder="请输入资产类别" />
       </el-form-item>
        <el-form-item label="资产位置:" prop="location">
          <el-input v-model="formData.location" :clearable="true"  placeholder="请输入资产位置" />
       </el-form-item>
        <el-form-item label="资产状态:" prop="status">
          <el-input v-model="formData.status" :clearable="true"  placeholder="请输入资产状态" />
       </el-form-item>
        <el-form-item label="供应商:" prop="vendor">
          <el-input v-model="formData.vendor" :clearable="true"  placeholder="请输入供应商" />
       </el-form-item>
        <el-form-item label="序列号:" prop="serialNumber">
          <el-input v-model="formData.serialNumber" :clearable="true"  placeholder="请输入序列号" />
       </el-form-item>
        <el-form-item label="保修信息:" prop="warranty">
          <el-input v-model="formData.warranty" :clearable="true"  placeholder="请输入保修信息" />
       </el-form-item>
        <el-form-item label="资产成本:" prop="cost">
          <el-input-number v-model="formData.cost" :precision="2" :clearable="true"></el-input-number>
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
  createAssets,
  updateAssets,
  findAssets
} from '@/api/assetm/Assets'

defineOptions({
    name: 'AssetsForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            assetId: undefined,
            assetName: '',
            category: '',
            location: '',
            status: '',
            vendor: '',
            serialNumber: '',
            warranty: '',
            cost: 0,
        })
// 验证规则
const rule = reactive({
               assetId : [{
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
      const res = await findAssets({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createAssets(formData.value)
               break
             case 'update':
               res = await updateAssets(formData.value)
               break
             default:
               res = await createAssets(formData.value)
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
