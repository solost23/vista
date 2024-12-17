<script setup>
    import { getVideos, getVideosByParams } from "../api/video.js"
    import { getCategoryies } from "../api/category.js"
    import { useRouter } from "vue-router"
    import { ref } from "@vue/reactivity"

    let router = useRouter()
    let token = localStorage.getItem("token")
    if (token === "" || token === null) {
        // 跳转登录路由
        router.push(
            {
                name: "login", 
            }
        )
    }

    let categories = ref([
        {
            id: 0, 
            title: "全部",
            introduce: "全部", 
        }, 
    ])

    let video = ref({
        keyword: "", 
        videoInfos: [
            {
                categoryId: 0, 
                categoryName: "", 
                commentCount: 0, 
                id: 0, 
                imageUrl: "",
                introduce: "", 
                thumbCount: 0, 
                title: "", 
                userId: 0, 
                userName: "",
                videoUrl: "", 
            }
        ], 
    })

    // 初始化分类列表
    let loadCategories = function loadCategories() {
        getCategoryies().then((response) => {
            if (!response.success) {
                alert(response.message)
                return 
            }
            // 拿到数据，展示
            categories.value = response.data.records
        })
    }
    loadCategories()

    // 初始化视频列表数据
    let loadVideos = function loadVideos() {
        let keyword = video.value.keyword
        getVideos({keyword: keyword, page: 1, size: 50}).then((response) => {
            if (!response.success) {
                alert(response.message)
                return 
            }
            // 拿到数据，展示数据
            video.value.videoInfos = response.data.records
        })
    }
    loadVideos()

    function toVideoDetailPage(videoId) {
        router.push(
            {
                path: `/videos/detail/${videoId}`, 
            }
        )
    }

    function handleGetVideoByCategoryId(categoryId) {
        // 请求接口
        getVideosByParams({categoryId: categoryId, page: 1, size: 50}).then((response) => {
            if (!response.success) {
                alert(response.message)
                return 
            }
            video.value.videoInfos = response.data.records
        }) 
    }
</script>

<template>
    <div id="index">
        <div>
            <div class="categories">
                <button v-for="category in categories" v-bind:key="category.id" @click="handleGetVideoByCategoryId(category.id)">{{ category.title }}</button>
            </div>
            <!-- 放个搜索框 -->
            <div class="search">
                <input type="text" placeholder="搜索" v-model="video.keyword" @keyup.enter="loadVideos"/>
            </div>
        </div>
        <!-- 视频列表 -->
        <div class="videoList">
            <ul>
               <li v-for="videoInfo in video.videoInfos" v-bind:key="videoInfo.id" @click="toVideoDetailPage(videoInfo.id)">
                    <!-- <video controls :src="videoInfo.videoUrl"></video> -->
                    <img :src="videoInfo.imageUrl"/>
                    <br/>
                    <div class="videoBaseInfo">
                        <strong>{{ videoInfo.title }}</strong>
                        <p>{{ videoInfo.introduce }}</p>
                    </div>
                </li> 
            </ul>
        </div>
    </div>
</template>

<style scoped>
    ul{
        max-width: 100%;
        display: flex;
        flex-wrap:wrap;
        list-style: none;
    }
    li{
        width: 300px;
        height: 200px;
    }
    img{
        width: 100px;
        height: auto;
    }
    p{
        margin-top: 0px;
    }
    button{
        margin-left: 5px;
    }
</style>
