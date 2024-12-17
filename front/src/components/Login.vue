<script setup>
    import { login } from "../api/user.js"
    import { useRouter } from 'vue-router';
    import { ref } from "@vue/reactivity";
    let router = useRouter()
    let user = ref({
        msg: "", 
        username: "", 
        password: "", 
    })

    function handleLogin() {
                let username = user.value.username
                let password = user.value.password
                if (username === "" || password === "") {
                    user.value.msg = "用户名或密码不为空!"
                    return
                }
                // 请求后端接口登录
                login({ userName: username, password: password, device: "web"}).then((response) => {
                    if (!response.success) {
                        user.value.msg = response.message
                        return 
                    }
                    // 将用户信息放入本地存储
                    localStorage.setItem("user", response.data)
                    localStorage.setItem("user_id", response.data.ID)
                    localStorage.setItem("avatar", response.data.avatar)
                    localStorage.setItem("token", response.data.token)

                    // 跳转路由, 使用path也可以
                    router.push(
                        {
                            name: "main", 
                        }
                    )

                })
            }

    function toRegisterPage() {
        router.push(
            {
                name: "register",    
            }
        )
    }
</script>
    
<template>
    <div id="login">
        <table id="input">
            <tr>
                <th>账号:</th>
                <td><input type="text" placeholder="请输入账号" v-model="user.username"/></td>
            </tr>
            <tr>
                <th>密码:</th>
                <td><input type="password" placeholder="请输入密码" v-model="user.password"/></td>
            </tr>
            <tr>
                <th style="width: max-content"></th>
                <td><span stype="color: red; font-size: smaller;">{{ user.msg }}</span></td>
            </tr>
            <tr>
                <th style="width: max-content;"></th>
                <td><button @click="handleLogin">登录</button>&nbsp;&nbsp;&nbsp;&nbsp;<button id="registerBut" @click="toRegisterPage">注册</button></td>
            </tr>
        </table>
    </div>
</template>

<style scoped>
    #input{
        margin: auto;
    }
</style>
