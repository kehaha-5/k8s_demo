import {
    fetchApi
} from "./fetch"


export const setUnHealth = () => {
    return fetchApi(`/api/set_un_health`, {
        method: 'GET',
    }).then((response) => response.json());
};