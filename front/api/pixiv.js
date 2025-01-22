import { dfGetax, getax } from '@/common/request/index';
import { getVal } from '@sorarain/utils';
import { vilipixSearchResultFormat, vilipixImgPathFormat } from './utils';
// const BaseUrl = 'http://localhost:3001'
const BaseUrl = 'http://pixivapi.adicw.cn';
/**
 * 获取动漫的相关图片列表-来源pixiv
 * @param param
 */
export async function vilipixSearch({ name = '', limit = 20, offset = 0, sort = 'hot', painter = '' } = {}) {
    try {
        const realName = name
            .split(' ')[0]
            .split('/')[0]
            .split('(')[0]
            .replace(/(第一季|第二季|第四季|第五季|第六季|第七季|BD|无修|剧场版|？)/, '');
        const { data } = await dfGetax(`https://www.vilipix.com/api/v1/picture/public?limit=${limit}&tags=${realName}&sort=${sort}&offset=${offset}&author_user_id=${painter}`);
        return vilipixSearchResultFormat(data);
    }
    catch (e) {
        return {
            list: [],
            total: 0
        };
    }
}
/**
 * 获取动漫的相关图片列表-来源pixiv
 * @param param
 */
export async function vilipixRank({ limit = 20, offset = 0, type = 0 } = {}) {
    try {
        const { data } = await dfGetax(`https://www.vilipix.com/api/v1/picture/ranking?limit=${limit}&offset=${offset}&type=${type}`);
        return {
            list: getVal(() => data.data.rows, []).map((item) => ({
                date: item.created_at,
                title: item.title,
                orgurl: item.original_url,
                preurl: vilipixImgPathFormat(item.regular_url),
                id: item.picture_id,
                w: item.width,
                h: item.height,
                commentTotal: item.comment_total,
                likeTotal: item.like_total,
                tags: item.tags,
                total: item.page_total
            })),
            total: data?.data?.count || 0
        };
    }
    catch (e) {
        return {
            list: [],
            total: 0
        };
    }
}
/**
 * 获取动漫图片详情，和vilipixSearch关联
 * @param id
 */
export async function getVilipixPicMain(id) {
    try {
        const [{ data }, { data: { data: orgImgs = [] } }] = await Promise.all([
            dfGetax(`${BaseUrl}/illust/${id}`),
            getax(`api/vilipix/${id}`)
        ]);
        return {
            orgImgs,
            author: {
                id: data.authorId,
                avatar: data.authorAvatar,
                name: data.authorName
            },
            ...data,
            date: new Date(data.date)
        };
    }
    catch {
        return null;
    }
}
/**
 * 获取首页热门tag
 * @param id
 * @returns
 */
export async function getVilipixHotTags() {
    try {
        const { data = [] } = await dfGetax(`${BaseUrl}/hot-tags`);
        return data.slice(0, 10);
    }
    catch {
        return [];
    }
}
/**
 * 获取画师信息
 * @returns
 */
export async function getVilipixerInfo(id) {
    try {
        const { data } = await dfGetax(`${BaseUrl}/user/${id}`);
        // return data.slice(0, 10)
        return {
            avatar: data?.avatar || '',
            name: data?.name || '',
            desc: data?.desc || ''
        };
    }
    catch {
        return null;
    }
}
/**
 * 获取作者作品集-来源pixiv
 * @param param
 */
export async function vilipixerSearch({ painterId, limit = 20, offset = 0, sort = 'new' } = { painterId: '' }) {
    try {
        const { data } = await dfGetax(`https://www.vilipix.com/api/v1/picture/public?sort=${sort}&type=0&author_user_id=${painterId}&limit=${limit}&offset=${offset}`);
        return vilipixSearchResultFormat(data);
    }
    catch (e) {
        return {
            list: [],
            total: 0
        };
    }
}
