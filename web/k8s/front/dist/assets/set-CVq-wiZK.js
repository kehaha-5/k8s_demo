import{e as i,c as d,b as t,d as s,f as a,v as u,o as p}from"./index-F8PVqnuV.js";import{s as m}from"./redis-ChKnNJad.js";import"./fetch-BU9cnwi6.js";const v=t("h1",null,"添加数据",-1),y={__name:"set",setup(k){const l=i({key:"",value:""}),r=()=>{m(l).then(n=>{n.code==200&&alert("设置成功")})};return(n,e)=>(p(),d("div",null,[v,t("h2",null,[s("key: "),a(t("input",{type:"text","onUpdate:modelValue":e[0]||(e[0]=o=>l.key=o)},null,512),[[u,l.key]])]),t("h2",null,[s("value: "),a(t("input",{type:"text","onUpdate:modelValue":e[1]||(e[1]=o=>l.value=o)},null,512),[[u,l.value]])]),t("button",{onClick:e[2]||(e[2]=o=>r())},"提交")]))}};export{y as default};
