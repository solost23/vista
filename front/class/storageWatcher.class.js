/**
 * storage监听者
 */
class StorageWatcher {
    deps = new Map();
    constructor() {
        this.watch();
    }
    addDep(storageKey, cb) {
        const dep = this.deps.get(storageKey);
        if (dep) {
            dep.add(cb);
        }
        else {
            this.deps.set(storageKey, new Set([cb]));
        }
    }
    notify(storageKey) {
        const dep = this.deps.get(storageKey);
        if (!dep)
            return;
        dep.forEach((cb) => cb());
    }
    watch() {
        window.addEventListener('storage', (e) => {
            if (e.url.includes('main.html#')) {
                this.notify(e.key || '');
            }
        });
    }
}
const watcher = new StorageWatcher();
/**
 * 方便storage存储类的扩展继承类
 * @param key storage的键
 */
export class StorageWatcherCling {
    constructor(key) {
        watcher.addDep(key, () => {
            this.getStore();
        });
    }
    /** 需要在子类中重写此方法 */
    getStore() {
        //
    }
}
export default watcher;
