import{f as e}from"./fetch-BU9cnwi6.js";const o=t=>e(`/api/mysql_get?id=${t}`,{method:"GET"}).then(s=>s.json()),r=t=>e("/api/mysql_set",{method:"POST",headers:{"Content-Type":"application/json;charset=utf-8;"},body:JSON.stringify(t)}).then(s=>s.json());export{o as g,r as s};
