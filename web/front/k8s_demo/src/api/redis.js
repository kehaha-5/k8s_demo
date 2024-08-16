import {
    fetchApi
} from "./fetch"

export const getRedis = (key) => {
    return fetchApi(`/api/redis_get?key=${key}`, {
        method: 'GET',
    }).then((response) => response.json());
};


export const setRedis = (data) => {
    return fetchApi(`/api/redis_set`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json;charset=utf-8;'
        },
        body: JSON.stringify(data)
    }).then((response) => response.json());
};