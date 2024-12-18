import { request } from "../utils/request.js"

export function getVideos(params) {
    return request(
        {
            url: "/api/vista/videos/search", 
            method: "get",
            params: params,  
        }
    )
}

export function getVideoInfo(id) {
    return request(
        {
            url: "/api/vista/videos/" + id, 
            method: "get", 
        }
    )
}

export function getVideoComments(params) {
    return request(
        {
            url: "/api/vista/comments", 
            method: "get", 
        }
    )
}

export function createVideoInfo(data) {
    return request(
        {
            url: "/api/vista/videos", 
            method: "post", 
            data: data, 
            headers: {
                "Content-Type": "multipart/form-data"
            }
        }
    )
}

export function uploadVideo(data) {
    return request(
        {
            url: "/api/vista/videos/vid", 
            method: "post", 
            data: data, 
            headers: {
                "Content-Type": "multipart/form-data"
            }
        }
    )
}

export function uploadImage(data) {
    return request(
        {
            url: "/api/vista/videos/img", 
            method: "post", 
            data: data, 
            headers: {
                "Content-Type": "multipart/form-data"
            }
        }
    )
}

// 根据条件查询视频列表
export function getVideosByParams(params) {
    return request (
        {
            url:"/api/vista/videos", 
            method: "get", 
            params: params, 
        }
    )
}
