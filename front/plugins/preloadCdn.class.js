import { DF_CDNS } from '@/common/cdns';
/**
 * link cdn挂载
 */
export default class PreloadCdn {
    preloadCdns = [...DF_CDNS];
    preloadMap = {
        stylesheet: {
            as: 'style'
        },
        font: {
            as: 'font'
        },
        javascript: {
            as: 'script'
        }
    };
    constructor() {
        this.preload();
    }
    preload() {
        // 不要合并下列循环
        // 预加载挂载
        this.preloadCdns.forEach(({ link, rel }) => {
            const el = document.createElement('link');
            el.rel = 'preload';
            el.as = this.preloadMap[rel].as;
            el.href = link;
            document.head.appendChild(el);
        });
        // 使用挂载
        this.preloadCdns.forEach(({ link, rel, unable }) => {
            if (unable && unable === true)
                return;
            if (rel === 'javascript') {
                const el = document.createElement('script');
                el.type = 'text/javascript';
                el.src = link;
                document.body.append(el);
            }
            else {
                const el = document.createElement('link');
                el.rel = rel;
                el.href = link;
                document.head.appendChild(el);
            }
        });
    }
}
/**
 * cdn挂载
 * @returns
 */
export function createPreloadCdn() {
    return new PreloadCdn();
}
