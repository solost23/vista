import router from '@/entry/index/router'
import { AxiosInstance } from 'axios'
import { ElNotification } from 'element-plus'
import { LoginStore } from '@/stores/login.store'
/**
 * axios辅助函数
 */
export default class AxiosUtils {
  private instance: AxiosInstance | null = null
  constructor(instance: AxiosInstance) {
    this.instance = instance
    this.response()
    this.request()
  }
  /**
   * 响应拦截器
   */
  private response() {
    if (this.instance === null) return
    this.instance.interceptors.response.use(
      (response) => {
        // store.state.loading = true

        // 检查是否登陆过期
        if (response.data && (response.data.code === 1999 || response.data.code === 1998)) {
          ElNotification({
            type: 'error',
            title: 'Token 过期',
            message: '登录过期，请重新登录'
          })

          const loginStore = LoginStore()
          loginStore.delToken()
          loginStore.delUser()

          router.push({name: 'Login'})
          return Promise.reject(new Error('Token 过期'))
        }
        return response
      },
      (error) => {
        const status = error.toString()
        if (
          [
            'timeout ',
            'Invalid URL',
            '401',
            '403',
            '404',
            'Network Error'
          ].some((item) => status.includes(item))
        ) {
          // store.dispatch('logout')
          ElNotification({
            type: 'error',
            title: '资源获取错误',
            message: '请求源存在问题'
          })
          // router.push({ name: 'Setting' })
        }
        return Promise.reject()
      }
    )
  }
  /**
   * 请求拦截器
   */
  private request() {
    if (this.instance === null) return
    this.instance.interceptors.request.use((request) => {
      // store.state.loading = false
      // const token = this.getToken()
      // if (token) {
      //   request.headers['authorization'] = this.getToken()
      // }

      const loginStore = LoginStore()
      if (loginStore.getToken()) {
        request.headers['token'] = loginStore.getToken();
      }

      return request
    })
  }
}
