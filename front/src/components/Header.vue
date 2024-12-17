<script setup>
    import { logout } from "../api/user.js"
    import { useRouter } from "vue-router"
    import { getUserInfo } from "../api/user"
    import { ref } from "@vue/reactivity"

    let router = useRouter()

    let isLogin = ref(false)

    let userInfo = ref({
        user_name: "", 
        nickname: "", 
        avatar: "", 
    })

    function handleLogout() {
        logout({ device: "web" }).then((response) => {
            if (!response.success) {
                console.log(response.message)
                return 
            }
            localStorage.removeItem("token")
            router.push(
                {
                    name: "login", 
                }
            )
        })
    }

    function loadUserInfo() {
        let userId = localStorage.getItem("user_id")
        getUserInfo(userId).then((response) => {
            if (!response.success) {
                return 
            }
            isLogin.value = true
            userInfo.value = response.data
        })

    }
    loadUserInfo()

    function toUploadVideo() {
        router.push(
            {
                name: "uploadVideo", 
            }
        )
    }

    function toIndex() {
        router.push(
            {
                name: "main", 
            }
        )
    }

    function toLoginPage() {
        router.push(
            {
                name: "login", 
            }
        )
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
    <div id="top">
        <div id="left" @click="toIndex">
            <!--网站头像 + 网站名称-->
            <div><img id="icon" src="../assets/circled-play.png" alt/></div>
            <!-- <span>小米星球</span> -->
        </div>
        <div id="middle">
            <!-- <span>导航栏</span> -->
            <div id="uploadVideo">
                <button @click="toUploadVideo">视频上传</button>
            </div>
        </div>
        <div id="right">
            <div id="isNotLogin" v-if="!isLogin">
                <button @click="toLoginPage">登录</button>
                &nbsp;&nbsp;&nbsp;&nbsp;
                <button @click="toRegisterPage">注册</button>
            </div>
            <div id="login" v-else>
                <div>
                    <div><img id="avatar" v-bind:src="userInfo.avatar" alt/></div>
                    <!-- <div><strong>{{ userInfo.user_name }}</strong></div> -->
                </div>
                <div><button id="logout" @click="handleLogout">注销</button></div>
            </div>
        </div>
    </div>
</template>

<style scoped>
    #top{
        background: while;
        display: flex;
        align-items: center;
    }
    #left{
        /* background: green; */
        height: 100px;
        width: 300px;
    }
    #middle{
        background: white;
        height: 100px;
        width: 2000px;
    }
    #icon{
        width: 100px;
        height: 100px;
    }
    #uploadVideo{
        margin-top: 33px;
    }
    #right{
        /* background: green; */
        height: 100px;
        width: 400px;
    }
    #avatar{
        margin-top: 15px;
        width: 50px;
        height: 50px;
    }
    #login{
        display: flex;
    }
    #logout{
        margin-left: 20px;
        margin-top: 15px;
    }
    #isNotLogin{
        margin-top: 33px;
    }
    #login{
        margin-top: 33px;
    }
</style>
