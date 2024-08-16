export const fetchApi = (url, opt) => {
    return fetch(import.meta.env.VITE_APP_API + url, opt)
}