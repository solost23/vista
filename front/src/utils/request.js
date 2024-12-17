import axios from "axios"

export const request = axios.create(
    {
        timeout: 30000, 
        withCredentials: true, // 异步请求携带cookie
        headers: {
            // 设置后端需要的传参类型
            "Content-Type": "application/json", 
            "token": localStorage.getItem("token") || "", 
            "X-Requested-With": "XMLHttpRequest", 
        }
    }
);

// 请求拦截器
request.interceptors.request.use(
    function(request) {
        // 在发送请求之前做些什么
        console.log(request)
        return request;
    },
    function(error) {
        // 对请求错误做些什么
        console.log(error);
        return Promise.reject(error);
    }
); 

// 响应拦截器
request.interceptors.response.use(
    function(response) {
        /*
            2xx 范围内的状态码都会触发该函数。
            对响应数据做操作
         */
        return response.data;
    }, 
    function(error) {
        /*
            超出 2xx 范围的状态码都会触发该函数。
            对响应错误做点什么
        */
        return Promise.reject(error);
    }
);
