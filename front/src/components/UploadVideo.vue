<script setup>
    import { ref } from "@vue/reactivity"
    import { uploadVideo, uploadImage, createVideoInfo } from "../api/video.js"
    import { getCategoryies } from "../api/category.js"
    import { useRouter } from "vue-router"

    let router = useRouter()
    let videoInsertForm = ref({
        title: "", 
        introduce: "", 
        categoryId: 0, 
        imageUrl: "",
        videoUrl: "", 
    })

    let msg = ref("")

    let categories = ref([
        {
            id: 0, 
            title: "",
            introduce: "", 
        }
    ])

    function handleUploadVid(event) {
        let file = event.target.files[0]
        uploadVideo({file: file}).then((response) => {
            if (!response.success) {
                msg.value = response.message
                return 
            }
            videoInsertForm.value.videoUrl = response.data
        })
    }

    function handleUploadImg(event) {
        let file = event.target.files[0]
        uploadImage({file: file}).then((response) => {
            if (!response.success) {
                msg.value = response.message
                return 
            } 
            videoInsertForm.value.imageUrl = response.data
        })
    }

    function handleSubmit() {
        // 提交表单信息
        createVideoInfo(videoInsertForm.value).then((response) => {
            if (!response.success) {
                msg.value = response.messge
                return 
            }
            // 成功，跳转到视频详情
            router.push(
                {
                    path: `/videos/detail/${response.data}`, 
                }
            )
        })
    }

    function loadCategories() {
        getCategoryies().then((response) => {
            if (!response.success) {
                msg.value = response.message
                return 
            }
            categories.value = response.data.records
        })
    }
    loadCategories()
</script>

<template>
    <div class="nav">
        <form class="videoInfo">
            <table>
                <tr>
                    <label>视频标题: </label>
                    <input type="text" v-model="videoInsertForm.title" placeholder="请输入视频标题"/>
                </tr>
                <hr/>
                <tr>
                    <label>视频介绍: </label>
                    <textarea v-model="videoInsertForm.introduce"></textarea>
                </tr>
                <hr/>
                <tr>
                    <label>所属类别: </label>
                    <select v-model="videoInsertForm.categoryId">
                        <option v-for="category in categories" v-bind:key="category.id" v-bind:value="category.id">{{ category.title }}</option>
                    </select>
                </tr>
                <hr/>
                <tr>
                    <label>提交封面: </label>
                    <input ref="input" type="file" id="uploaderImg" @change="handleUploadImg"/>
                </tr>
                <tr>
                    <label>提交视频: </label>
                    <input ref="input" type="file" id="uploader" @change="handleUploadVid"/>
                </tr>
                <hr/>
                <p class="msg">{{ msg }}</p>
            </table>
            <input type="button" value="提交" @click="handleSubmit"/>
        </form>
    </div>
</template>

<style scoped>
    .videoInfo {
        margin-left: 450px;
    }
    .msg {
        color: red;
    }
</style>
