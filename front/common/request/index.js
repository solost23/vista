import axios from 'axios';
import { BASE_COMIC_URL, timeout } from './config';
import AxiosUtils from './axiosUtils.class';
// 创建公共实例
const instance = axios.create({
    baseURL: BASE_COMIC_URL,
    timeout
});
new AxiosUtils(instance);
/**
 * 请求 - get
 * @param url 请求地址
 * @returns
 */
export function getax(url) {
    // onDownloadProgress ???
    return instance.get(url);
}
/**
 * 请求 - post
 * @param url 请求地址
 * @param params 请求参数
 * @param config 请求配置
 * @returns
 */
export function postax(url, params = {}, config = {}) {
    return instance.post(url, params, config);
}
/**
 * 请求 - delete
 * @param url 请求地址
 * @returns
 */
export function deleteax(url) {
    return instance.delete(url);
}
/**
 * 基于axios的请求 - get
 * @param url 完整请求地址
 * @returns
 */
export function dfGetax(url) {
    const instance = axios.create({
        timeout
    });
    return instance.get(url);
}
/**
 * 创建一个请求驳回源
 * @returns
 */
export function createCancelToken() {
    const CancelToken = axios.CancelToken;
    return CancelToken.source();
}
