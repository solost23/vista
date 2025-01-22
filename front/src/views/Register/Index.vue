<template>
  <div class="register-main">
    <div class="register-main__part">
      <h2 class="register-main__title">注册</h2>
      <el-form
        aria-autocomplete="off"
        :model="registerForm"
        :rules="rules"
        label-width="100px"
        style="transform: translate(-30px)"
      >
        <!-- <el-form-item label="邮箱" prop="email">
          <el-input
            v-model="registerForm.email"
            placeholder="请输入邮箱"
          ></el-input>
        </el-form-item> -->
        <el-form-item label="用户名" prop="username">
          <el-input
            v-model="registerForm.username"
            placeholder="请输入用户名"
          ></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input
            type="password"
            v-model="registerForm.password"
            placeholder="请输入密码"
            show-password
          ></el-input>
        </el-form-item>
        <el-form-item label="确认密码" prop="confirmPassword">
          <el-input
            type="password"
            v-model="registerForm.confirmPassword"
            placeholder="请确认密码"
            show-password
          ></el-input>
        </el-form-item>
        <el-form-item label="头像" prop="avatar">
          <el-upload
            class="register-main__upload"
            action="/api/vista/upload"
            :show-file-list="false"
            :on-success="handleAvatarSuccess"
            :on-error="handleAvatarError"
            :before-upload="beforeAvatarUpload"
            :file-list="registerForm.fileList"
          >
            <img v-if="registerForm.avatar" :src="registerForm.avatar" class="register-main__upload__avatar"></img>
            <!-- <i v-else class="el-icon-plus avatar-uploader-icon"></i> -->
            <img v-else src="/android-chrome-192x192.png" class="register-main__upload__avatar"></img>
          </el-upload>
        </el-form-item>
        <!-- <el-form-item label="验证码" prop="code">
          <el-input
            style="width: 150px"
            v-model="registerForm.code"
            placeholder="请输入验证码"
            maxlength="6"
            clearable
          ></el-input>
          <el-button
            round
            class="code-btn"
            type="primary"
            v-if="isTime"
            @click="getCode(registerForm.email)"
            >获取验证码</el-button
          >
          <el-button
            round
            class="code-btn"
            size="Large"
            color="#c0c4c3"
            v-if="!isTime"
            >{{ currentTime }}后重新获取</el-button
          >
        </el-form-item> -->
        <el-button 
        class="register-main__btn" 
        type="primary" 
        @click="register"
        auto-insert-space 
        @keyup.enter="register"
        >注册</el-button
        >
        <div style="text-align: right; transform: translate(0, 30px)">
          <el-link type="success" @click="goToLogin">已有账号？去登录</el-link>
        </div>
      </el-form>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import { ElNotification } from "element-plus";
import * as Api from "@/api";

// 定义验证规则的类型
interface ValidationRule {
  required?: boolean;
  message?: string;
  trigger: string | string[];
  validator?: (rule: any, value: any, callback: (error?: Error) => void) => void;
}

// 定义整个验证规则对象的类型
interface ValidationRules {
  [key: string]: ValidationRule[];
}

// 假设的表单状态
interface registerForm {
  username: string;
  password: string;
  confirmPassword: string;
  // 可以添加其他字段
  avatar: string;
  fileList: any[];
}

const registerForm = ref<registerForm>({
  username: '',
  password: '',
  confirmPassword: '',
  avatar: '', 
  fileList: [], 
});

// 定义验证规则，注意这里不直接引用 formState，而是在验证器内部通过闭包访问
const rules: ValidationRules = {
username: [{ required: true, message: "请输入用户名", trigger: 'blur' }],
password: [{ required: true, message: "请输入密码", trigger: 'blur' }],
confirmPassword: [
  { required: true, message: "请确认密码", trigger: 'blur' },
  {
    validator: (rule: any, value: any, callback: (error?: Error) => void) => {
      if (value === '') {
        callback(new Error('请再次输入密码'));
      } else if (value !== registerForm.value.password) {
        callback(new Error('两次密码不一致'));
      } else {
        callback();
      }
    },
    trigger: 'blur',
  },
  ],
// 可以添加其他字段的规则
};

  
const router = useRouter();

// 接口交互
const register = async () => {
  const data = await Api.register({
    username: registerForm.value.username.trim(), 
    password: registerForm.value.password.trim(), 
    role: 1,   
    avatar: registerForm.value.avatar,
  })
  if (data) {
    router.push("/login");
    ElNotification({
      title: '注册成功',
      type: 'success',
    })
  }
}
  
const goToLogin = () => {
  router.push("/login");
};

const beforeAvatarUpload = (file: File) => {
  const isImage = file.type.startsWith('image/');;
  const isLt2M = file.size / 1024 / 1024 < 2;

  if (!isImage) {
    ElNotification({
      title: '上传头像失败',
      message: '上传头像只能是图片格式!',
      type: 'warning',
    })
  }
  if (!isLt2M) {
    ElNotification({
      title: '上传头像失败',
      message: '上传头像大小不能超过 2MB!',
      type: 'warning',
    })
  }
}

const handleAvatarSuccess = (res: any, file: any) => {
  registerForm.value.avatar = res.data;
}
const handleAvatarError = (err: any, file: any) => {
  ElNotification({
    title: '上传头像失败',
    message: '上传头像失败!',
    type: 'warning',
  })
}
</script>

<style lang="less" scoped>
@import './style/style.less';
</style>
