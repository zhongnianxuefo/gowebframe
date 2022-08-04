<template>
  <a-row type="flex" justify="center" align="top" style="margin-top:50px;">
    <a-col :span="6">
      <a-card title="用户登录" >
        <a-form
            ref="formRef"
            :model="formState"
            name="basic"
            :label-col="{ span: 8 }"
            :wrapper-col="{ span: 16 }"
            autocomplete="off"
            @finish="onFinish"
            @finishFailed="onFinishFailed"
        >

          <a-form-item
              label="用户名"
              name="username"
              :rules="[{ required: true, message: '请输入用户名！' }]"
          >
            <a-input v-model:value="formState.username" />
          </a-form-item>

          <a-form-item
              label="密码"
              name="password"
              :rules="[{ required: true, message: '请输入密码！' }]"
          >
            <a-input-password v-model:value="formState.password" />
          </a-form-item>

          <a-form-item :wrapper-col="{ offset: 8, span: 16 }">
            <a-button type="primary" html-type="submit">登录</a-button>
            <a-button style="margin-left: 10px" @click="resetForm">重置</a-button>
          </a-form-item>
        </a-form>
      </a-card>
    </a-col>
  </a-row>
</template>
<script setup>
import { reactive, ref} from 'vue';
import { message } from 'ant-design-vue';

import router from "@/base/route"
import request from "@/base/request"


const formRef = ref ();
const formState = reactive ({
  username: '',
  password: '',
});

const onFinish = async(values) => {
  if (values) {
    request.login(values).then(onRequestOK).catch(onRequestErr);
  }
};
const onRequestOK=(r) => {
  console.log('Request OK:', r);
  if (r.data["ResponseCode"]==0) {
     router.push("/home")
  }else{
      message.error(r.data["ErrMsg"]);
  }
}
const onRequestErr = (error) => {
  console.log("Request Err:",error);
  message.error("登录失败，请检查网络连接情况！");
}

const onFinishFailed = errorInfo => {
  console.log('Finish Failed:', errorInfo);
};

const resetForm = () => {
  formRef.value.resetFields();
};


</script>
