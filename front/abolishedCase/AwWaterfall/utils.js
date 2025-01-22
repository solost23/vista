export const WATERFALL_KEY = Symbol('WATERFALL_KEY');
export function arrayAverag(arr, num) {
    const resultArr = [];
    num = arr.length / num;
    for (let i = 0, len = arr.length; i < len; i += num) {
        resultArr.push(arr.slice(i, i + num));
    }
    return resultArr;
}
