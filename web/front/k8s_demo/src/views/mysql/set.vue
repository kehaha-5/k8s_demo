
<template>
    <div>
        <h1>添加数据</h1>
        <h2>用户名：
            <input type="text" v-model="user.user_name">
        </h2>
        <h2>手机号： <input type="number" v-model="user.user_phone">
        </h2>
        <h2>年龄： <input type="number" v-model="user.age">
        </h2>
        <h2>性别： <select v-model="user.sex">
                <option value="0">请选择</option>
                <option value="1">男</option>
                <option value="2">女</option>
            </select>
        </h2>
        <button @click="submit()">提交</button>
    </div>
</template>

<script setup>
import { reactive } from "vue";
import { setMysql } from "@/api/mysql"

const user = reactive({
    user_name: "",
    user_phone: "",
    sex: "0",
    age: 0
});

const submit = () => {
    setMysql(user).then((res) => {
        if (res.code == 200) {
            alert("添加成功")
        }
    })
};
</script>