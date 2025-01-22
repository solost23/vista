import { mountComponent } from '@sorarain/use';
import ImagePreviewComponent from './ImagePreview.vue';
let instance;
function initInstance() {
    ;
    ({ instance } = mountComponent(ImagePreviewComponent));
}
const ImagePreview = (options) => {
    !instance && initInstance();
    instance.init(options);
    return instance;
};
export { ImagePreview };
