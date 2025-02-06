import { Plugin } from 'vue'
import UndrawUi from 'undraw-ui'
import 'undraw-ui/dist/style.css'

// const comps = [
//   UndrawUi
// ]

const plugins = [UndrawUi]
export const undrawUiInit: Plugin = {
  install(app) {
    // comps.forEach((comp) => {
    //   app.component(comp.name, comp)
    // })
    plugins.forEach((plugin) => {
      app.use(plugin)
    })
  }
}