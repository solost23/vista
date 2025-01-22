import { jsonParse } from '@sorarain/utils';
import { ref } from 'vue';
import { StorageWatcherCling } from './storageWatcher.class';
const PLAY_HISTORY_STORE_KEY = 'PLAY_HISTORY_STORE';
/**
 * 播放历史缓存
 */
class PlayHistory extends StorageWatcherCling {
    constructor() {
        super(PLAY_HISTORY_STORE_KEY);
    }
    /** 缓存列表 */
    cache_ = ref([]);
    get cache() {
        return this.cache_.value;
    }
    /**
     * 添加历史数据
     * @param item 动漫信息
     */
    add(item) {
        const index = this.cache_.value.findIndex((ch) => ch.id === item.id);
        if (!!~index) {
            this.cache_.value.splice(index, 1);
        }
        this.cache_.value.unshift({
            ...item,
            date: new Date().getTime()
        });
    }
    /**
     * 删除指定id数据
     * @param id 动漫id
     */
    remove(id) {
        const index = this.cache_.value.findIndex((ch) => ch.id === id);
        if (!!~index) {
            this.cache_.value.splice(index, 1);
        }
    }
    /**
     * 数据-本地保存
     */
    saveStore() {
        localStorage.setItem(PLAY_HISTORY_STORE_KEY, JSON.stringify(this.cache_.value));
    }
    /**
     * 数据-本地获取
     */
    getStore() {
        const data = jsonParse(localStorage.getItem(PLAY_HISTORY_STORE_KEY), []);
        if (data instanceof Array) {
            this.cache_.value = data;
        }
    }
    /**
     * 数据-本地清空
     */
    clearStore() {
        this.cache_.value.splice(0);
        localStorage.removeItem(PLAY_HISTORY_STORE_KEY);
    }
}
let instance;
export function createPlayHistory() {
    if (!instance) {
        instance = new PlayHistory();
    }
    return instance;
}
export function getPlayHistoryInstance() {
    return instance;
}
