import { jsonParse } from '@sorarain/utils';
import { ref } from 'vue';
import { StorageWatcherCling } from './storageWatcher.class';
const PLAY_PROGRESS_STORE_KEY = 'PLAY_PROGRESS_STORE';
class PlayProgress extends StorageWatcherCling {
    constructor() {
        super(PLAY_PROGRESS_STORE_KEY);
    }
    cache_ = ref([]);
    get cache() {
        return this.cache_.value;
    }
    get cacheMap() {
        return this.cache_.value.reduce((totol, item) => {
            totol[this.cacheItemToName(item)] = item;
            return totol;
        }, {});
    }
    get latestCacheMap() {
        return this.cache_.value.reduce((totol, item) => {
            if (!totol[item.comicId]) {
                totol[item.comicId] = item;
            }
            else if (item.date > totol[item.comicId].date) {
                totol[item.comicId] = item;
            }
            return totol;
        }, {});
    }
    createDate() {
        return new Date().getTime();
    }
    cacheItemToName(item) {
        return `${item.comicId}-${item.orgId}-${item.name}`;
    }
    /**
     * 判断是否存在此集
     * @param name
     * @returns
     */
    has(name) {
        return typeof this.cacheMap[name] !== 'undefined';
    }
    /**
     * 添加集数缓存，如果存在，则会被更新
     * @param item
     */
    add(item) {
        if (!item.orgId)
            return;
        const name = this.cacheItemToName(item);
        const cacheItem = {
            ...item,
            date: this.createDate()
        };
        if (!this.has(name)) {
            this.cache_.value.push(cacheItem);
        }
        else {
            const index = this.cache_.value.findIndex((ca) => this.cacheItemToName(ca) === this.cacheItemToName(item));
            this.cache_.value.splice(index, 1, cacheItem);
        }
    }
    /**
     * 删除指定集数缓存
     * @param item
     */
    remove(item) {
        const index = this.cache_.value.findIndex((ca) => this.cacheItemToName(ca) === this.cacheItemToName(item));
        if (index !== -1) {
            this.cache_.value.splice(index, 1);
        }
    }
    /**
     * 获取指定缓存进度
     * @param name
     * @returns
     */
    get(name) {
        if (this.has(name)) {
            return this.cacheMap[name].progress;
        }
        else {
            return 0;
        }
    }
    saveStore() {
        localStorage.setItem(PLAY_PROGRESS_STORE_KEY, JSON.stringify(this.cache_.value));
    }
    getStore() {
        const data = jsonParse(localStorage.getItem(PLAY_PROGRESS_STORE_KEY), []);
        if (data instanceof Array) {
            this.cache_.value = data;
        }
    }
    /**
     * 获取指定动漫最后一次记录
     * @param comicId
     * @returns
     */
    getLatestCache(comicId) {
        return this.latestCacheMap[comicId];
    }
}
let instance;
export function createPlayProgress() {
    if (!instance) {
        instance = new PlayProgress();
    }
    return instance;
}
export function getPlayProgressInstance() {
    return instance;
}
