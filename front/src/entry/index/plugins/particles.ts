import { Plugin } from 'vue'

import VueParticles from 'particles.vue3'

const comps = [
    VueParticles,
]

const plugins = [VueParticles]
export const particlesInit: Plugin = {
    install(app) {
        comps.forEach((comp) => {
            app.component(comp.name, comp)
        })
        plugins.forEach((plugin) => {
            app.use(plugin)
        })
    }
}
