import router from '@/entry/index/router';
/**
 * 前往动漫详情
 * @param id
 * @param type
 * @param op { latest-0|1 是否最新集 }
 * @returns
 */
export function toComicMain(id, type = 'push', op = {
    latest: 0
}) {
    return router[type]({
        name: 'ComicMain',
        params: {
            id
        },
        query: {
            latest: op.latest
        }
    });
}
/**
 * 前往Pixiv图片详情
 * @param id
 * @param params
 * @param type
 * @returns
 */
export function toPixivMain(el, item, type = 'push') {
    const rect = el.getBoundingClientRect();
    return router[type]({
        name: 'PixivMain',
        params: {
            rect: JSON.stringify({
                width: rect.width | 0,
                height: rect.height | 0,
                x: rect.x | 0,
                y: rect.y | 0,
                path: item.preurl,
                radius: getComputedStyle(el).borderRadius
            }),
            id: item.id
        }
    });
}
export function toPixivPainter(id) {
    router.push({
        name: 'PixivPainter',
        params: {
            id
        }
    });
}
export function createNewMainPath(id) {
    return `${location.origin}/main.html#/${id}`;
}
export function openNewMain(id) {
    window.open(createNewMainPath(id));
}
