
<template>
    <div>
        <input type="number" placeholder="请输入id" v-model="user.id">
        <button @click="query()">查询</button>
    </div>
    <div>
        <h1>查询结果：</h1>
        <h2>用户名：{{ user.name }}</h2>
        <h2>手机号：{{ user.phone }}</h2>
        <h2>年龄：{{ user.age }}</h2>
        <h2>性别：{{ user.sex }}</h2>
        <h2>创建时间：{{ user.createdAt }}</h2>
    </div>
</template>

<script setup>
import { reactive } from "vue";
import { getMysql } from "@/api/mysql"

const user = reactive({
    id: "",
    name: "",
    phone: "",
    sex: "",
    age: "",
    createdAt: ""
});

const query = () => {
    getMysql(user.id).then((res) => { 
        if (res.code == 200){
            user.id = res.data.ID
            user.createdAt = res.data.CreatedAt
            user.age = res.data.Age
            user.name = res.data.UserName
            user.phone = res.data.UserPhone
            user.sex = res.data.Sex
        }else{
            alert(res.msg)
        }
    })
};
</script>