
export const getMysql = (id) => {
    return fetch(`/api/mysql_get?id=${id}`, {
        method: 'GET',
    }).then((response) => response.json());
};


export const setMysql = (data) => {
    return fetch(`/api/mysql_set`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json;charset=utf-8;'
        },
        body: JSON.stringify(data)
    }).then((response) => response.json());
};