import AwDailogVue from './AwDailog.vue';
import { mountComponent } from '@sorarain/use';
let instance;
function initInstance() {
    ;
    ({ instance } = mountComponent(AwDailogVue));
}
const AwDailog = (options) => {
    !instance && initInstance();
    return instance.init(options);
};
const AwDailogPlugin = {
    install(app) {
        app.config.globalProperties.$awDailog = AwDailog;
    }
};
export { AwDailog, AwDailogPlugin };
