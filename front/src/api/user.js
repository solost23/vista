import { request } from "../utils/request.js"

export function register(data) {
    return request(
        {
            url: "/api/vista/register", 
            method: "post", 
            data: data, 
        }
    )
}

export function login(data) {
    return request(
        {
            url: "/api/vista/login", 
            method: "post", 
            data: data, 
        }
    )
}

export function logout(data) {
    return request(
        {
            url: "/api/vista/users/logout", 
            method: "post", 
            data: data, 
        }
    )
}

export function getUserInfo(id) {
    return request(
        {
            // or `/api/users/${id}`
            url: "/api/vista/users/" + id, 
            method: "get", 
        }
    )
}

export function uploadAvatar(data) {
    return request(
        {
            url: "/api/vista/register/avatar", 
            method: "post", 
            data: data, 
            headers: {
                "Content-Type": "multipart/form-data"
            }
        }
    )
}
