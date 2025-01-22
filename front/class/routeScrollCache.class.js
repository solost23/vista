import { nextTick } from 'vue';
/**
 * 路由滚动缓存管理
 */
class RouteScrollCacheManage {
    cache_ = new Map();
    get cache() {
        return this.cache_;
    }
    /**
     * 增加缓存
     * @param routeName 路由名称或地址
     * @param routeMeta 路由的meta值
     * @returns
     */
    addCache(routeName, routeMeta) {
        const elName = String(routeMeta.elName || '');
        if (!routeName || !elName)
            return;
        const el = document.querySelector(elName);
        this.cache_.set(routeName, {
            elName,
            scrollTop: el?.scrollTop || 0
        });
    }
    /**
     * 获取缓存
     * @param routeName 路由名称或地址
     * @returns 缓存
     */
    readCache(routeName) {
        return this.cache_.get(routeName);
    }
    /**
     * 路由滚动更新
     * @param routeName 路由名称或地址
     */
    async setScroll(routeName) {
        if (this.cache_.has(routeName)) {
            await nextTick();
            const cache = this.readCache(routeName);
            const el = document.querySelector(cache.elName);
            el && (el.scrollTop = cache.scrollTop);
        }
    }
    /**
     * 缓存清空
     */
    clearCache() {
        this.cache_.clear();
    }
}
let instance;
export function createRouteSCM() {
    if (!instance) {
        instance = new RouteScrollCacheManage();
    }
    return instance;
}
export function getRouteSCMInstance() {
    return instance;
}
