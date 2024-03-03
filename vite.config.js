import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
// VueJsx
import vueJsx from "@vitejs/plugin-vue-jsx";
// src for @vuejs
import { join } from "path";
// unplugin-vue-components 按需导入插件
import Components from "unplugin-vue-components/vite";
import { AntDesignVueResolver } from "unplugin-vue-components/resolvers";
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    // Vue
    vue(),
    // 按需导入组件
    Components({
      resolvers: [
        AntDesignVueResolver({
          importStyle: false, // css in js
        }),
      ],
    }),
    // Vue jsx
    vueJsx(),
  ],
  // scss编译器
  css: {
    preprocessorOptions: {},
  },
  resolve: {
    alias: {
      "@": join(__dirname, "src"),
    },
  },
});
