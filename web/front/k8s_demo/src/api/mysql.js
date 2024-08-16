import {
    fetchApi
} from "./fetch"

export const getMysql = (id) => {
    return fetchApi(`/api/mysql_get?id=${id}`, {
        method: 'GET',
    }).then((response) => response.json());
};


export const setMysql = (data) => {
    return fetchApi(`/api/mysql_set`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json;charset=utf-8;'
        },
        body: JSON.stringify(data)
    }).then((response) => response.json());
};