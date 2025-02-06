import { createApp } from 'vue'

import App from './App.vue'
import router from './router'

import { createComicMouseright } from '@/class/comicMouseright.class'
import { createVueInit } from '@/utils/vue/index'
import { baseLoader } from '../base'
import { elementPlusInit } from './plugins/elementPlus'
import { undrawUiInit } from './plugins/undrawUi'

const app = createApp(App)

baseLoader(app)
createComicMouseright()

app
.use(elementPlusInit)
.use(undrawUiInit)
.use(createVueInit)
.use(router)
app.mount('#app')
