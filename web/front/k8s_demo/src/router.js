import {
    createRouter,
    createWebHistory
} from "vue-router";

const router = createRouter({
    history: createWebHistory(),
    routes: [{
            path: "/",
            name: "index",
            component: () => import("@/views/index.vue"),
        },
        {
            path: "/mysql",
            name: "mysql",
            component: () => import("@/views/mysql/index.vue"),
            redirect: "/mysql/get",
            children: [{
                path: "/mysql/get",
                name: "mysql-get",
                component: () => import("@/views/mysql/get.vue"),
            }, {
                path: "/mysql/set",
                name: "mysql-set",
                component: () => import("@/views/mysql/set.vue"),
            }, ]
        },
        {
            path: "/redis",
            name: "redis",
            component: () => import("@/views/redis/index.vue"),
            redirect: "/redis/get",
            children: [{
                path: "/redis/get",
                name: "redis-get",
                component: () => import("@/views/redis/get.vue"),
            }, {
                path: "/redis/set",
                name: "redis-set",
                component: () => import("@/views/redis/set.vue"),
            }, ]
        },
    ],
});


export default router;