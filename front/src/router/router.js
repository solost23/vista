import { createRouter, createWebHistory } from "vue-router"

const router = createRouter({
    history: createWebHistory(), 
    routes: [
        {
            path: "/", 
            name: "index",
            component: () => import("../components/Videos.vue") 
        }, 
        {
            path: "/login", 
            name: "login", 
            component: () => import("../components/Login.vue")
        }, 
        {
            path: "/register", 
            name: "register", 
            component: () => import("../components/Register.vue")
        }, 
        {
            path: "/videos/detail//:id(\\d+)", 
            name: "VideoDetail", 
            component: () => import("../components/VideoDetail.vue")
        }, 
        {
            path: "/:pathMatch(.*)*", 
            name: "NotFound", 
            component: () => import("../components/NotFound.vue")
        }, 
        {
            path: "/videos/upload", 
            name: "uploadVideo", 
            component: () => import("../components/UploadVideo.vue")
        }
    ]
})

export default router
