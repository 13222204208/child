<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="姓名:" prop="name">
          <el-input v-model="formData.name" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="姓别:" prop="gender">
          <el-select v-model="formData.gender" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in genderOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="出生日期:" prop="birthday">
          <el-input v-model="formData.birthday" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="身高:" prop="height">
          <el-input-number v-model="formData.height" :precision="2" :clearable="false"></el-input-number>
        </el-form-item>
        <el-form-item label="体重:" prop="weight">
          <el-input-number v-model="formData.weight" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="体貌特征:" prop="appearance">
          <el-input v-model="formData.appearance" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="所属用户ID:" prop="uid">
          <el-input v-model.number="formData.uid" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Baby'
}
</script>

<script setup>
import {
  createBaby,
  updateBaby,
  findBaby
} from '@/api/baby'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const genderOptions = ref([])
const formData = ref({
            name: '',
            gender: undefined,
            birthday: '',
            height: 0,
            weight: 0,
            appearance: '',
            uid: 0,
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               gender : [{
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
      const res = await findBaby({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rebaby
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    genderOptions.value = await getDictFunc('gender')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createBaby(formData.value)
               break
             case 'update':
               res = await updateBaby(formData.value)
               break
             default:
               res = await createBaby(formData.value)
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
