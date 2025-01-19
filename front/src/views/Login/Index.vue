<template>
  <div class="login-main">
    <div class="login-main__part">
      <h2 class="login-main__title">登录</h2>
      <el-form
        :model="loginForm"
        :rules="rules"
        label-width="100px"
        style="transform: translate(-30px)"
      >
        <el-form-item label="用户名" prop="username">
          <el-input
            v-model="loginForm.username"
            placeholder="请输入用户名"
            clearable
          ></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input
            type="password"
            v-model="loginForm.password"
            placeholder="请输入密码"
            show-password
            clearable
          ></el-input>
        </el-form-item>
 
        <el-button
          class="login-main__btn"
          type="primary"
          @click="login"
          auto-insert-space
          @keyup.enter="login"
          >登录</el-button
        >
        <div class="login-main__register" style="text-align: right; transform: translate(0, 30px)">
          <el-link type="warning" @click="changeUrl('/register')"
            >没有账号？去注册</el-link
          >
        </div>
      </el-form>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import { ElNotification } from "element-plus";

import * as Api from '@/api'

import { LoginStore } from '@/stores/login.store'

const loginForm = ref({
  username: "",
  password: "",
});

const rules = {
  username: [{ required: true, message: "请输入用户名", trigger: "blur" }],
  password: [{ required: true, message: "请输入密码", trigger: "blur" }],
};

const router = useRouter();

// 接口交互
const login = async () => {
  const data = await Api.login({
    username: loginForm.value.username,
    password: loginForm.value.password,
    device: 'web', 
  })
  if (data) {
    // 保存 token
    const loginStore = LoginStore();
    loginStore.saveToken(data.token);
    loginStore.saveUser(data);
    router.push("/");
    ElNotification({
      title: '登录成功',
      message: '欢迎回来',
      type: 'success',
    })
  } 
};

const changeUrl = (url: string) => {
  router.replace(url);
};
</script>
 
<style lang="less" scoped>
@import './style/main';
</style>