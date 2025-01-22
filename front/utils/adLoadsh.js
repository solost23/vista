import { ElNotification } from 'element-plus';
/**
 * 对象自动赋值（两者需要相同的类型依赖）
 * @param a
 * @param b
 */
export function autoObjAssign(a, b) {
    Object.keys(a).forEach((k) => {
        if (typeof b[k] !== 'undefined') {
            a[k] = b[k];
        }
    });
}
/**
 * 节点事件节流，基于requestAnimationFrame，常用于scroll
 * @param cb
 * @returns
 */
export function eventThrottle(cb) {
    let bool = false;
    return function (...args) {
        if (bool)
            return;
        bool = true;
        window.requestAnimationFrame(() => {
            cb.apply(this, args);
            bool = false;
        });
    };
}
/**
 * 将内容复制到粘贴板
 * @param text
 */
export function copyText(text) {
    const save = (e) => {
        e.clipboardData.setData('text/plain', text.toString());
        e.preventDefault();
    };
    document.addEventListener('copy', save);
    document.execCommand('copy');
    document.removeEventListener('copy', save);
    ElNotification({
        title: '内容复制',
        message: '内容已经成功复制到粘贴板~',
        type: 'success'
    });
}
