import { mountComponent } from '@sorarain/use';
import ContextmenuItx from './Index.vue';
// eslint-disable-next-line @typescript-eslint/ban-types
let instance;
function initInstance() {
    ;
    ({ instance } = mountComponent(ContextmenuItx));
}
const Contextmenu = (options) => {
    !instance && initInstance();
    instance.init(options);
    return instance;
};
export { Contextmenu };
