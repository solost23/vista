import { deleteax, getax, postax } from '@/common/request/index';
import { getVal } from '@sorarain/utils';
import { newError, badRequestNotify } from './utils';
import { ElNotification } from 'element-plus';
export * from './type';
export * from './pixiv';
/**
 * 根据名称获取动漫列表
 * @param param
 * @returns
 */
export async function searchComic(param) {
    try {
        const { data: { data: { results, pagetotal } } } = await getax(`/video/search?page=${param.page}&name=${param.name}`);
        if (results instanceof Array) {
            return {
                data: results,
                total: (pagetotal || 0) * 20
            };
        }
        else {
            throw newError();
        }
    }
    catch {
        badRequestNotify('/video/search');
        return {
            data: [],
            total: 0
        };
    }
}
/**
 * 动漫筛选
 * @param param
 * @returns
 */
export async function filterComic(param) {
    try {
        const api = Object.entries(param).reduce((total, [k, v]) => {
            return v !== '' ? `${total}&${k}=${v}` : total;
        }, '/video/filter?');
        const { data } = await getax(api);
        return {
            data: getVal(() => data.data.results, []).map((item) => ({
                cover: item.cover,
                id: item.id,
                season: item.season,
                title: item.title
            })),
            total: data?.data?.total || 0
        };
    }
    catch {
        badRequestNotify('/vieo/filter');
        return {
            data: [],
            total: 0
        };
    }
}
/**
 * 获取动漫详情
 * @param id
 * @returns
 */
export async function getComicMain(id) {
    try {
        const { data: { data } } = await getax(`video/${id}`);
        const playlist = new Map();
        Object.entries(data.playlist || {}).forEach(([k, v]) => {
            playlist.set(k, v.map((item, index) => ({
                name: String(item.title),
                value: index
            })));
        });
        return {
            playlist,
            title: data.title,
            season: data.season || '未知',
            region: data.region || '未知',
            rank: data.rank || '',
            master: data.master || '未知',
            lang: data.lang || '未知',
            firstDate: data.first_date || '未知',
            cover: data.cover || '',
            voiceActors: data.actors || [],
            cates: data.categories || []
        };
    }
    catch {
        badRequestNotify('/video/:id');
        return null;
    }
}
/**
 * 获取视频地址集
 * @param key
 * @returns
 */
export async function getVideoUrl(key) {
    try {
        const { data: { data } } = await await getax(`/video/${key}/playlist`);
        return Object.entries(data).map(([k, v]) => ({
            key: k,
            value: (v instanceof Array ? v : []).map((url) => url.replaceAll("'", '').split('?url=').pop())
        }));
    }
    catch (e) {
        badRequestNotify('/video/:key/playlist');
        console.error(e);
        return [];
    }
}
/**
 * 获取混合列表（热门、最新更新、轮播、每周更新列表、番外、完结日漫、国漫）
 * @returns
 */
export async function getHomeMixData() {
    try {
        const { data } = await getax('/video/index');
        const listFormat = (list) => list.slice(0, 10).map((item) => ({
            cover: item.cover,
            id: item.id,
            season: item.season,
            title: item.title
        }));
        return {
            hots: getVal(() => data.data.hots.results, []).map((item) => ({
                cover: item.cover,
                id: item.id,
                season: item.season,
                title: item.title
            })),
            latest: listFormat(getVal(() => data.data.latest, [])),
            banner: getVal(() => data.data.banner, []).map((item) => ({
                cover: item.cover,
                id: item.id || '-1',
                title: item.title || '未知'
            })),
            perweek: Object.entries(getVal(() => data.data.perweek, {})).map(([k, v]) => ({
                name: `周${['一', '二', '三', '四', '五', '六', '日'][+k]}`,
                key: k,
                value: (v || []).map((e) => ({
                    id: e.id,
                    season: e.season || '未知',
                    title: e.title || '未知'
                }))
            })),
            tv: listFormat(getVal(() => data.data.theatre_comic, [])),
            endJp: listFormat(getVal(() => data.data.japancomic, [])),
            cn: listFormat(getVal(() => data.data.chinese_comic, []))
        };
        // return
    }
    catch (e) {
        badRequestNotify('/video/index');
        console.error(e);
        return null;
    }
}
/**
 * 获取动漫筛选条件
 * @returns
 */
export async function getComicFilterConfig() {
    try {
        const { data } = await getax('/video/config');
        return getVal(() => data.data.filtersConfig, []).map((item) => ({
            id: item.id,
            name: item.name,
            value: (item.categories || []).map((key) => ({
                name: key.classname,
                value: key.classid
            }))
        }));
    }
    catch {
        badRequestNotify('/video/config');
        return [];
    }
}
/**
 * 登陆
 * @returns
 */
export async function login(params) {
    try {
        const { data } = await postax('/login', params);
        if (data.code != 0 && data.success == false) {
            ElNotification({
                title: '登陆失败',
                message: data.message,
                type: 'error'
            });
            return null;
        }
        return {
            id: data.data.ID,
            nickname: data.data.nickname,
            avatar: data.data.avatar,
            token: data.data.token,
        };
    }
    catch {
        badRequestNotify('/login');
        return null;
    }
}
/**
 * 注册
 * @returns
 */
export async function register(params) {
    try {
        const { data } = await postax('/register', params);
        if (data.code != 0 && data.success == false) {
            ElNotification({
                title: '注册失败',
                message: data.message,
                type: 'error'
            });
            return null;
        }
    }
    catch {
        badRequestNotify('/register');
    }
    return 'success';
}
/**
 * 删除视频
 */
export async function deleteVideo(key) {
    try {
        const { data } = await deleteax(`/video/${key}`);
        if (data.code != 0 && data.success == false) {
            ElNotification({
                title: "删除失败",
                message: data.message,
                type: 'error'
            });
            return null;
        }
    }
    catch {
        badRequestNotify('/video/delete');
    }
    return 'success';
}
