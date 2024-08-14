
<template>
    <div>
        <h1>添加数据</h1>
        <h2>key:
            <input type="text" v-model="data.key">
        </h2>
        <h2>value: <input type="text" v-model="data.value"></h2>
        <button @click="submit()">提交</button>
    </div>
</template>

<script setup>
import { reactive } from "vue";
import { setRedis } from "@/api/redis"

const data = reactive({
    key: "",
    value: ""
});

const submit = () => {
    setRedis(data).then((res) => {
        if (res.code == 200) {
            alert("设置成功")
        }
    })
};
</script>