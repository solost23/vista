import { jsonParse } from '@sorarain/utils';
import { defineStore } from "pinia";
const TOKEN_STORE_KEY = 'TOKEN';
const USER_STORE_KEY = 'USER';
export function getToken() {
    return localStorage.getItem(TOKEN_STORE_KEY) || '';
}
export function getUser() {
    const data = localStorage.getItem(USER_STORE_KEY);
    const origin = {
        id: 0,
        nickname: '',
        avatar: '',
        token: '',
    };
    return jsonParse(data, origin) || origin;
}
export const LoginStore = defineStore('Login', {
    state: () => ({
        token: '',
        user: {}
    }),
    actions: {
        init() {
            this.token = getToken();
            this.user = getUser();
        },
        saveToken(token) {
            this.token = token;
            localStorage.setItem(TOKEN_STORE_KEY, token);
        },
        getToken() {
            return getToken();
        },
        saveUser(user) {
            this.user = user;
            localStorage.setItem(USER_STORE_KEY, JSON.stringify(user));
        },
        getUser() {
            return getUser();
        }
    }
});
