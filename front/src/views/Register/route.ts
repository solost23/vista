import { WEB_NAME } from '@/common/static'

const Register = () => import('@/views/Register/Index.vue')
export default {
  path: '/register',
  name: 'Register',
  component: Register,
  meta: {
    title: WEB_NAME + '-注册'
  },
  props: true,
}