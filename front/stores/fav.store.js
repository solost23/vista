import { defineStore } from 'pinia';
import { getComicFavInstance } from '@/class/comicFav.class';
export const useFavStore = defineStore('favStore', {
    getters: {
        comicFavs() {
            return getComicFavInstance().fav;
        }
    },
    actions: {
        comicFav(comic) {
            getComicFavInstance().favHandler(comic);
        },
        isFavComic(id) {
            return getComicFavInstance().has(id);
        },
        saveComicFav() {
            getComicFavInstance().saveStore();
        },
        exChange(aId, bId) {
            let aIndex = 0, bIndex = 0;
            this.comicFavs.forEach((item, index) => {
                if (aId === item.comicId) {
                    aIndex = index;
                }
                if (bId === item.comicId) {
                    bIndex = index;
                }
            });
            getComicFavInstance().exChange(aIndex, bIndex);
        }
    }
});
