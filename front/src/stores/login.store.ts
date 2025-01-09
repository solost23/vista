import { jsonParse } from '@sorarain/utils'
import { defineStore } from "pinia"

import * as Api from '@/api'

const TOKEN_STORE_KEY = 'TOKEN'
const USER_STORE_KEY = 'USER'

export function getToken() {
    return localStorage.getItem(TOKEN_STORE_KEY) || '';
}

export function getUser() {
    const data = localStorage.getItem(USER_STORE_KEY)
    const origin = {
        id: 0, 
        nickname: '',
        avatar: '',
        token: '', 
    }
    return jsonParse<Api.Login>(data, origin) || origin
}

export const LoginStore = defineStore('Login', {
    state: () => ({
        token: '', 
        user: {} as Api.Login
    }), 
    actions: {
        init() {
            this.token = getToken();
            this.user = getUser();
        }, 
        saveToken(token: string) {
            this.token = token;
            localStorage.setItem(TOKEN_STORE_KEY, token);
        }, 
        getToken() {
            return getToken();
        }, 
        saveUser(user: Api.Login) {
            this.user = user;
            localStorage.setItem(USER_STORE_KEY, JSON.stringify(user));
        },
        getUser() {
            return getUser();
        }
    }
})
