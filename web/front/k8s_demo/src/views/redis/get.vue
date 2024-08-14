
<template>
    <div>
        <input type="text" placeholder="请输入key" v-model="data.key">
        <button @click="query()">查询</button>
    </div>
    <div>
        <h1>查询结果：</h1>
        <h2>key：{{ data.key }}</h2>
        <h2>value：{{ data.value }}</h2>
    </div>
</template>

<script setup>
import { reactive } from "vue";
import { getRedis } from "@/api/redis"

const data = reactive({
    key: "",
    value: ""
});

const query = () => {
    getRedis(data.key).then((res) => {
        if (res.code == 200) {
            data.value = res.data
        }else{
            alert(res.msg)
        }
    })
};
</script>