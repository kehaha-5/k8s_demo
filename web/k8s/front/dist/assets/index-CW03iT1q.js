import{f as c}from"./fetch-BU9cnwi6.js";import{c as u,a,w as l,b as t,F as i,r as d,o as h,d as r}from"./index-F8PVqnuV.js";const m=()=>c("/api/set_un_health",{method:"GET"}).then(e=>e.json()),p=t("h1",null,"请选择操作",-1),f=t("br",null,null,-1),k=t("h1",null,"设置服务器为不健康状态",-1),C={__name:"index",setup(e){const _=()=>{m().then(o=>{o.code==200&&alert("设置成功")})};return(o,n)=>{const s=d("router-link");return h(),u(i,null,[p,a(s,{to:"/mysql"},{default:l(()=>[r("Go to Mysql")]),_:1}),f,a(s,{to:"/redis"},{default:l(()=>[r("Go to Redis")]),_:1}),k,t("button",{onClick:n[0]||(n[0]=x=>_())},"点击")],64)}}};export{C as default};
