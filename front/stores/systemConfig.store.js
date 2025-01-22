import { jsonParse } from '@sorarain/utils';
import { defineStore } from 'pinia';
const REQUEST_BASEURL_STORE_KEY = 'REQUEST_BASEURL_STORE';
const STATIC_RESOURCES_STORE_KEY = 'STATIC_RESOURCES_STORE';
export function getServerIp() {
    return localStorage.getItem(REQUEST_BASEURL_STORE_KEY) || '';
}
export function getStaticResource() {
    const data = localStorage.getItem(STATIC_RESOURCES_STORE_KEY);
    const origin = {
        videoProgressCurIcon: ''
    };
    return jsonParse(data, origin) || origin;
}
export const useSystemConfigStore = defineStore('SystemConfig', {
    state: () => ({
        serverIp: '',
        staticResource: {
            videoProgressCurIcon: ''
        }
    }),
    actions: {
        init() {
            this.getServerIp();
            this.getStaticResource();
        },
        saveServerIp(url) {
            this.serverIp = url;
            localStorage.setItem(REQUEST_BASEURL_STORE_KEY, url);
        },
        getServerIp() {
            this.serverIp = getServerIp();
        },
        saveStaticResource(resource) {
            this.staticResource = resource;
            localStorage.setItem(STATIC_RESOURCES_STORE_KEY, JSON.stringify(resource));
        },
        getStaticResource() {
            this.staticResource = getStaticResource();
        }
    }
});
