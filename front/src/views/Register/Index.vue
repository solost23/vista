<template>
    <div class="login">
      <Particles id="tsparticles" class="login__particles" :options="options" />
   
      <div class="loginPart">
        <h2>注册</h2>
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
          <el-button class="btn" type="primary" @click="register">注册</el-button>
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
  // import { getCodeService, registerService } from "@/api/user";
  import { ElNotification } from "element-plus";

  import * as Api from "@/api";
   
  // const isTime = ref(true);
  // const currentTime = ref(10);
  // const getCode = async () => {
  //   console.log("开始发送验证码");
  //   isTime.value = false;
  //   currentTime.value = 10;
  //   const email = registerForm.value.email;
  //   console.log(email);
  //   const interval = setInterval(() => {
  //     currentTime.value--;
  //     if (currentTime.value === 0) {
  //       clearInterval(interval);
  //       isTime.value = !isTime.value;
  //     }
  //   }, 1000);
  //   try {
  //     await getCodeService(email);
  //     ElMessage.success("验证码发送成功");
  //   } catch (error) {
  //     console.error("发送验证码时出现错误：", error);
  //   }
  // };
   
  // const registerForm = ref({
  //   username: "", 
  //   password: "",
  //   confirmPassword: "",
  // });
   
  // const registerFormRef = ref({
  //   email: "",
  //   password: "",
  //   confirmPassword: "",
  //   code: "",
  // });
   
  // const rules: validationRules = {
  //   // email: [
  //   //   { required: true, message: "请输入邮箱", trigger: "blur" },
  //   //   { type: "email", message: "邮箱格式不正确", trigger: ["blur", "change"] },
  //   // ],
  //   username: [{ required: true, message: "请输入用户名", trigger: "blur" }],
  //   password: [{ required: true, message: "请输入密码", trigger: "blur" }],
  //   confirmPassword: [
  //     { required: true, message: "请确认密码", trigger: "blur" },
  //     {
  //       validator: (rule: any, value: string, callback: (error?: Error)) => {
  //         if (value == "") {
  //           callback(new Error("请再次输入密码"));
  //         } else if (value !== registerForm.value.password) {
  //           callback(new Error("两次密码不一致"));
  //         } else {
  //           callback();
  //         }
  //       },
  //       trigger: "blur",
  //     },
  //   ],
  //   // code: [{ required: true, message: "请输入验证码", trigger: "blur" }],
  // };

  // 定义验证规则的类型
  interface ValidationRule {
    required?: boolean;
    message: string;
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
  }

  const registerForm = ref<registerForm>({
    username: '',
    password: '',
    confirmPassword: '',
  });

  // 定义验证规则，注意这里不直接引用 formState，而是在验证器内部通过闭包访问
  const rules: ValidationRules = {
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
  confirmPassword: [
    { required: true, message: '请确认密码', trigger: 'blur' },
    {
      validator: (rule, value, callback) => {
        if (value === '') {
          callback(new Error('请再次输入密码'));
        } else if (value !== registerForm.password) {
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
   
  // const register = async () => {
  //   registerForm.value.email = registerForm.value.email.trim();
  //   registerForm.value.password = registerForm.value.password.trim();
  //   registerForm.value.code = registerForm.value.code.trim();
  //   const res = await registerService(registerForm.value);
  //   //ElMessage.success(res.message ? res.message : '注册成功!')
  //   ElMessage.success("注册成功!");
  //   goToLogin;
  // };

  // 接口交互
  const register = async () => {
    const data = await Api.register({
      username: registerForm.value.username.trim(), 
      password: registerForm.value.password.trim(), 
      role: 1,   
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
   
  const options = {
    fpsLimit: 60,
    interactivity: {
      detectsOn: "canvas",
      events: {
        onClick: {
          // 开启鼠标点击的效果
          enable: true,
          mode: "push",
        },
        onHover: {
          // 开启鼠标悬浮的效果(线条跟着鼠标移动)
          enable: true,
          mode: "grab",
        },
        resize: true,
      },
      modes: {
        // 配置动画效果
        bubble: {
          distance: 400,
          duration: 2,
          opacity: 0.8,
          size: 40,
        },
        push: {
          quantity: 4,
        },
        grab: {
          distance: 200,
          duration: 0.4,
        },
        attract: {
          // 鼠标悬浮时，集中于一点，鼠标移开时释放产生涟漪效果
          distance: 200,
          duration: 0.4,
          factor: 5,
        },
      },
    },
    particles: {
      color: {
        value: "#BA55D3", // 粒子点的颜色
      },
      links: {
        color: "#FFBBFF", // 线条颜色
        distance: 150, //线条距离
        enable: true,
        opacity: 0.4, // 不透明度
        width: 1.2, // 线条宽度
      },
      collisions: {
        enable: true,
      },
      move: {
        attract: { enable: false, rotateX: 600, rotateY: 1200 },
        bounce: false,
        direction: "none",
        enable: true,
        out_mode: "out",
        random: false,
        speed: 0.5, // 移动速度
        straight: false,
      },
      number: {
        density: {
          enable: true,
          value_area: 800,
        },
        value: 80, //粒子数
      },
      opacity: {
        //粒子透明度
        value: 0.7,
      },
      shape: {
        //粒子样式
        type: "star",
      },
      size: {
        //粒子大小
        random: true,
        value: 3,
      },
    },
    detectRetina: true,
  };
  </script>
   
  <style lang="less" scoped>
  .login {
    height: 100%;
    width: 100%;
    overflow: hidden;
  }
   
  .login__particles {
    height: 100%;
    width: 100%;
    background-size: cover;
    background-repeat: no-repeat;
    // background-image: url("@/assets/images/loginbg.png");
    opacity: 0.9;
    position: fixed;
    pointer-events: none;
  }
   
  .loginPart {
    position: absolute;
    /*定位方式绝对定位absolute*/
    top: 50%;
    left: 50%;
    /*顶和高同时设置50%实现的是同时水平垂直居中效果*/
    transform: translate(-50%, -50%);
    /*实现块元素百分比下居中*/
    width: 450px;
    padding: 50px;
    // background: rgba(255, 204, 255, 0.3);
    /*背景颜色为黑色，透明度为0.8*/
    box-sizing: border-box;
    /*box-sizing设置盒子模型的解析模式为怪异盒模型，
      将border和padding划归到width范围内*/
    box-shadow: 0px 15px 25px rgba(0, 0, 0, 0.5);
    /*边框阴影  水平阴影0 垂直阴影15px 模糊25px 颜色黑色透明度0.5*/
    border-radius: 15px;
    /*边框圆角，四个角均为15px*/
  }
   
  h2 {
    margin: 0 0 30px;
    padding: 0;
    color: #fff;
    text-align: center;
    /*文字居中*/
  }
   
  .btn {
    transform: translate(170px);
    width: 80px;
    height: 40px;
    font-size: 15px;
  }
  .code-btn {
    transform: translate(20px);
    width: 90px;
    height: 40px;
    font-size: 10px;
  }
  </style>