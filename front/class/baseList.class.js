import { reactive } from 'vue';
import { jsonParse } from '@sorarain/utils';
export default class BaseList {
    /**
     * 仓库名称
     */
    storeKey = '';
    /**
     * 列表子项key名称
     */
    listItemKey = null;
    constructor(storeKey, listItemKey) {
        this.storeKey = storeKey;
        this.listItemKey = listItemKey;
        this.getStore();
    }
    state = reactive({
        list: new Map()
    });
    get list() {
        return [...this.state.list].map(([_, v]) => v);
    }
    addItem(key, value) {
        this.state.list.set(key, value);
    }
    removeItem(key) {
        return this.state.list.delete(key);
    }
    clear() {
        this.state.list.clear();
        localStorage.removeItem(this.storeKey);
    }
    saveStore() {
        localStorage.setItem(this.storeKey, JSON.stringify(this.list));
    }
    getStore() {
        const data = jsonParse(localStorage.getItem(this.storeKey), []);
        if (!(data instanceof Array))
            return;
        data.forEach((item) => {
            if (this.listItemKey && item[this.listItemKey]) {
                this.addItem(item[this.listItemKey], item);
            }
        });
    }
}
