import{e as p,c as i,b as e,d as s,f as l,v as u,g as d,o as m}from"./index-F8PVqnuV.js";import{s as _}from"./mysql-l5bSr6Uq.js";import"./fetch-BU9cnwi6.js";const v=e("h1",null,"添加数据",-1),x=e("option",{value:"0"},"请选择",-1),f=e("option",{value:"1"},"男",-1),h=e("option",{value:"2"},"女",-1),V=[x,f,h],U={__name:"set",setup(b){const o=p({user_name:"",user_phone:"",sex:"0",age:0}),r=()=>{_(o).then(a=>{a.code==200&&alert("添加成功")})};return(a,t)=>(m(),i("div",null,[v,e("h2",null,[s("用户名： "),l(e("input",{type:"text","onUpdate:modelValue":t[0]||(t[0]=n=>o.user_name=n)},null,512),[[u,o.user_name]])]),e("h2",null,[s("手机号： "),l(e("input",{type:"number","onUpdate:modelValue":t[1]||(t[1]=n=>o.user_phone=n)},null,512),[[u,o.user_phone]])]),e("h2",null,[s("年龄： "),l(e("input",{type:"number","onUpdate:modelValue":t[2]||(t[2]=n=>o.age=n)},null,512),[[u,o.age]])]),e("h2",null,[s("性别： "),l(e("select",{"onUpdate:modelValue":t[3]||(t[3]=n=>o.sex=n)},V,512),[[d,o.sex]])]),e("button",{onClick:t[4]||(t[4]=n=>r())},"提交")]))}};export{U as default};
