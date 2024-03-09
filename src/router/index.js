import { createRouter, createWebHistory } from "vue-router";
import ErrorPage from "@/views/404.vue";

const routes = [
  {
    path: "/",
    name: "404",
    component: ErrorPage,
    // beforeEnter: (to, from) => {
    //   console.log("独享");
    // },
  },
  {
    path: "/main",
    name: "404",
    component: () => import("@/views/main.vue"),
    // beforeEnter: (to, from) => {
    //   console.log("独享");
    // },
    children: [
      {
        path: "home",
        name: "home",
        component: () => import("@/views/home/index.vue"), // component: import('../views/reg.vue')
      },
    ],
  },
  // {
  //   path: "/",
  //   name: "home",
  //   alias: ["/home", "/home2"], // 别名，可以定义很多个
  //   component: () => import("../views/home.vue"),
  //   // 重定向
  //   // redirect: '/welcome',
  //   redirect: (to) => {
  //     console.log(to);
  //     return {
  //       path: "/welcome",
  //       query: {
  //         name: "欢迎",
  //       },
  //     };
  //   },
  //   children: [
  //     {
  //       path: "/welcome",
  //       name: "welcome",
  //       component: () => import("../views/welcome.vue"), // component: import('../views/reg.vue')
  //     },
  //   ],
  // },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

// 路由守卫
router.beforeEach((to, from, next) => {
  // ...
  // 返回 false 以取消导航
  // return false;
  next();
});
export default router;
