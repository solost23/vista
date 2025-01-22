import { defineStore } from 'pinia';
import BaseList from '@/class/baseList.class';
export const useSearchHistory = defineStore('searchHistory', {
    state: () => ({
        handler: new BaseList('SEARCH_HISTORY', 'value')
    }),
    getters: {
        list(state) {
            return [...state.handler.list].sort((a, b) => b.date - a.date);
        }
    },
    actions: {
        add(text) {
            const date = new Date().getTime();
            this.handler.addItem(text, {
                value: text,
                date
            });
        },
        delete(text) {
            return this.handler.removeItem(text);
        },
        save() {
            this.handler.saveStore();
        },
        clear() {
            this.handler.clear();
        }
    }
});
