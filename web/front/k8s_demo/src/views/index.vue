<template>
    <h1>请选择操作</h1>
    <router-link to="/mysql">Go to Mysql</router-link>
    <br>
    <router-link to="/redis">Go to Redis</router-link>
    <h1>设置服务器为不健康状态</h1>
    <button @click="unHealth()">点击</button>
</template>

<script setup>
import { setUnHealth } from "@/api/health"

const unHealth = () => {
    setUnHealth().then(res => {
        if (res.code == 200) {
            alert("设置成功")
        }
    })
}
</script>