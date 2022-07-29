<template>
  <el-row class="row-bg"  justify="center">
    <el-col :span="6">
      <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <span>用户登录</span>
          </div>
        </template>
        <el-form
          ref="loginFormRef"
          :model="loginFormData"
          :rules="loginRules"
          label-width="120px"
          class="demo-ruleForm"
          :size="formSize"
          status-icon
        >
        <el-form-item label="用户名" prop="username"  >
          <el-input v-model="loginFormData.username" />
        </el-form-item>
        <el-form-item label="密码" prop="password"  >
          <el-input v-model="loginFormData.password" type="password"/>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm()">登录</el-button>
          <el-button @click="resetForm()"> 重置</el-button>
        </el-form-item>
      </el-form>
      </el-card>
    </el-col>
  </el-row>

</template>

<script scope setup>

import { reactive, ref } from 'vue'
import router from '@/route.js'

const checkUsername = (rule, value, callback) => {
  if (value=="") {
    return callback(new Error('请输入用户名'))
  } else {
    callback()
  }
}

const checkPassword = (rule, value, callback) => {
  if (value=="") {
    return callback(new Error('请输入密码'))
  } else {
    callback()
  }
}


const formSize = ref('default')
const loginFormRef = ref('')
const loginFormData = reactive({
  username: '',
  password: '',
})

const loginRules = reactive({
  username: [{ validator: checkUsername, trigger: 'blur' }],
  password: [{ validator: checkPassword, trigger: 'blur' }],
})

import axios from 'axios' // 引入axios
axios.defaults.baseURL = '/api/'




const submitForm =  () => {
  loginFormRef.value.validate(async(valid) => {
    if (valid) {
      const flag = await axios.post('login',loginFormData);

      console.log('submit!', flag)
      if (flag.data.success){
        await router.push("/home")

      }
    } else {
      console.log('error submit!')
    }
  })
}

const resetForm = () => {
  loginFormRef.value.resetFields()
}

</script>


<style>


</style>
