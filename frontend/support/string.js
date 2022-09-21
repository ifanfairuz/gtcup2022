export function filename(str) {
    return new String(str).substring(str.lastIndexOf('/') + 1);
}