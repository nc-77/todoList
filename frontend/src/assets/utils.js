// 过滤obj的空属性
function filterEmptyAttr(obj) {
    Object.keys(obj).forEach(function (key) {
        if (obj[key] === '' || obj[key] === null) {
            delete (obj[key])
        }
    });
    return obj
}
// 过滤obj1中和obj2相同的属性 及空属性
function filterSameAttr(obj1, obj2) {
    Object.keys(obj1).forEach(function (key) {
        if (obj1[key] === obj2[key] || (obj1[key] === '' || obj1[key] === null)) {
            delete (obj1[key])
        }
    });
    return obj1
}
function fmtDate(dateString) {
    const date = new Date(dateString);
    const dates = date.toDateString().split(" ");
    const fmtDate = dates[0] + "," + dates[1] + " " + dates[2];
    return fmtDate;
}
// 判断是否是今天或者明天
function judgeDate(dateString) {
    const date = new Date(dateString);
    const today = new Date();
    const tomorrow = new Date(
        today.getFullYear(),
        today.getMonth(),
        today.getDate() + 1
    );
    if (
        date.getFullYear() == today.getFullYear() &&
        date.getMonth() == today.getMonth() &&
        date.getDate() == today.getDate()
    ) {
        return "Today";
    } else if (
        date.getFullYear() == tomorrow.getFullYear() &&
        date.getMonth() == tomorrow.getMonth() &&
        date.getDate() == tomorrow.getDate()
    ) {
        return "Tomorrow";
    }
    return dateString;
}
export { filterEmptyAttr, filterSameAttr, fmtDate, judgeDate }