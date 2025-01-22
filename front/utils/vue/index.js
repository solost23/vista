import GlobalComp from '@comps/Global/index';
import { directs } from '@/utils/vue/directs';
/**
 * vue根实例附加初始化
 */
export default class VueInit {
    instance;
    constructor(instance) {
        this.instance = instance;
        this.useDirects();
        this.useComps();
    }
    useDirects() {
        for (const [k, v] of Object.entries(directs)) {
            this.instance.directive(k, v);
        }
        return this;
    }
    useComps() {
        GlobalComp.forEach((comp) => {
            this.instance.component(comp.name, comp);
        });
        return this;
    }
}
/**
 * 附加初始化
 * @param app 根vue
 * @returns
 */
export const createVueInit = {
    install(app) {
        new VueInit(app);
    }
};
