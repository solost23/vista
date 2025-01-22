import { jsonParse, arrChildSwap } from '@sorarain/utils';
import { ref } from 'vue';
import { StorageWatcherCling } from './storageWatcher.class';
const COMIC_FAV_STORE_KEY = 'COMIC_FAV_STORE';
class ComicFav extends StorageWatcherCling {
    constructor() {
        super(COMIC_FAV_STORE_KEY);
    }
    fav_ = ref([]);
    get fav() {
        return this.fav_.value;
    }
    getTime() {
        return new Date().getTime();
    }
    /**
     * 判断是否存在此集
     * @param comicId
     * @returns
     */
    has(comicId) {
        return !!~this.fav_.value.findIndex((item) => item.comicId === comicId);
    }
    index(comicId) {
        return this.fav_.value.findIndex((item) => item.comicId === comicId);
    }
    favHandler(comic) {
        const removeIndex = this.index(comic.comicId);
        if (!!~removeIndex) {
            this.fav_.value.splice(removeIndex, 1);
        }
        else {
            this.fav_.value.unshift({
                ...comic,
                favDate: this.getTime()
            });
        }
    }
    saveStore() {
        localStorage.setItem(COMIC_FAV_STORE_KEY, JSON.stringify(this.fav_.value));
    }
    getStore() {
        const data = jsonParse(localStorage.getItem(COMIC_FAV_STORE_KEY), []);
        if (data instanceof Array) {
            this.fav_.value = data;
        }
    }
    exChange(aIndex, bIndex) {
        arrChildSwap(this.fav_.value, aIndex, bIndex);
    }
}
let instance;
export function createComicFav() {
    if (!instance) {
        instance = new ComicFav();
    }
    return instance;
}
export function getComicFavInstance() {
    return instance;
}
