import { WEB_NAME } from '@/common/static'

const Login = () => import('@/views/Login/Index.vue')
export default {
  path: '/login',
  name: 'Login',
  component: Login,
  meta: {
    title: WEB_NAME + '-登陆'
  },
  props: true,
}