import { createApp } from "vue";
// Router
import router from "@/router/index.js";
// pinia
import { createPinia } from "pinia";
// ant-css
import "ant-design-vue/dist/reset.css";
// APP
import App from "./App.vue";

createApp(App).use(createPinia()).use(router).mount("#app");
