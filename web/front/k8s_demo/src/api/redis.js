
export const getRedis = (key) => {
    return fetch(`/api/redis_get?key=${key}`, {
        method: 'GET',
    }).then((response) => response.json());
};


export const setRedis = (data) => {
    return fetch(`/api/redis_set`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json;charset=utf-8;'
        },
        body: JSON.stringify(data)
    }).then((response) => response.json());
};