import { request } from "../utils/request.js"

export function getCategoryies() {
    return request(
        {
            url: "/api/vista/categories", 
            method: "get", 
        }
    )
}
