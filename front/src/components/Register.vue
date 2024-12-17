<script setup>
    import { register, uploadAvatar } from "../api/user.js"
    import { useRouter } from "vue-router"
    import { ref } from "@vue/reactivity"
    let router = useRouter()
    let user = ref(
        {
            username: "", 
            password: "", 
            nickname: "",
            role: "USER",
            avatar: "",
            introduce: "", 
            device: "web", 
        }
    )
    let msg = ref("")

    function toLoginPage() {
        router.push(
            {
                name: "login", 
            }
        )
    }

    function handleRegister() {
        let username = user.value.username 
        let password = user.value.password
        if (username === "" || password === "") {
            msg.value = "用户名或密码不为空"
            return 
        }
        register(user.value).then((response) => {
            if (!response.success) {
                msg.value = response.message
                return 
            }
            // 跳转登录路由
            router.push(
                {
                    name: "login", 
                }
            )
        })
    }

    function UploadImg(event) {
        let file = event.target.files[0]
        uploadAvatar({file: file}).then((response) => {
            if (!response.success) {
                msg.value = response.message
                return 
            }
            user.value.avatar = response.data
        })
    }

</script>

<template>
    <div id="register">
        <table id="input">
            <tr>
                <th>账号:</th>
                <td><input type="text" placeholder="请输入账号" v-model="user.username"/></td>
            </tr>
            <tr>
                <!-- 昵称 -->
                <th>昵称</th>
                <td><input type="text" placeholder="请输入昵称" v-model="user.nickname"/></td>
            </tr>
            <tr>
                <!-- 头像 -->
                <th>头像</th>
                <td><input ref="input" type="file" @change="UploadImg"/></td>
            </tr>
            <tr>
                <!-- 介绍 -->
                <th>介绍</th>
                <td><textarea v-model="user.introduce"></textarea></td>
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
                <td><button @click="toLoginPage">登录</button>&nbsp;&nbsp;&nbsp;&nbsp;<button id="registerBut" @click="handleRegister">注册</button></td>
            </tr>
        </table>
    </div>
</template>

<style scoped>
    #input{
        margin: auto;
    }
</style>
