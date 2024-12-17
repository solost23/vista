import { request } from "../utils/request.js"

export function getComments(params) {
    return request(
        {
            url: "/api/comments", 
            method: "get", 
            params: params, 
        }
    )
}

export function createComment(data) {
    return request(
        {
            url: "/api/comments", 
            method: "post", 
            data: data, 
        }
    )
}
