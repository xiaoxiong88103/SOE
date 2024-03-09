<template>
  <a-layout style="height: 100vh">
    <a-layout-header class="header">
      <div class="logo" />
      <!-- <a-menu theme="dark" mode="horizontal" :style="{ lineHeight: '64px' }">
        <a-menu-item key="1">nav 1</a-menu-item>
        <a-menu-item key="2">nav 2</a-menu-item>
        <a-menu-item key="3">nav 3</a-menu-item>
      </a-menu> -->
      <span style="color: #fff">SEO运维系统</span>
    </a-layout-header>
    <a-layout>
      <a-layout-sider
        width="200"
        style="background: #fff; position: relative"
        v-model:collapsed="collapsed"
        :trigger="null"
        collapsible
      >
        <div class="triggerCount">
          <a-button
            style="height: 35px; margin: 5px 3px"
            @click="() => (collapsed = !collapsed)"
          >
            <MenuUnfoldOutlined v-if="collapsed" />
            <MenuFoldOutlined v-else />
          </a-button>
        </div>
        <a-menu
          v-model:openKeys="state.openKeys"
          v-model:selectedKeys="state.selectedKeys"
          mode="inline"
          :inline-collapsed="state.collapsed"
          :items="items"
        ></a-menu>
      </a-layout-sider>
      <a-layout style="padding: 0 24px 24px">
        <a-breadcrumb style="margin: 20px 40px">
          <a-breadcrumb-item>Home</a-breadcrumb-item>
          <a-breadcrumb-item>List</a-breadcrumb-item>
          <a-breadcrumb-item>App</a-breadcrumb-item>
        </a-breadcrumb>
        <a-layout-content
          :style="{
            background: '#fff',
            padding: '24px',
            margin: 0,
            minHeight: '280px',
            overflow: 'hidden',
            overflowY: 'auto',
          }"
        >
          <router-view />
        </a-layout-content>
      </a-layout>
    </a-layout>
  </a-layout>
</template>
<script setup>
import { reactive, watch, h, ref } from "vue";
import {
  MenuFoldOutlined,
  MenuUnfoldOutlined,
  PieChartOutlined,
  MailOutlined,
  DesktopOutlined,
  InboxOutlined,
  AppstoreOutlined,
} from "@ant-design/icons-vue";
const collapsed = ref(false);
const state = reactive({
  preOpenKeys: ["sub1"],
});
const items = reactive([
  {
    key: "1",
    icon: () => h(PieChartOutlined),
    label: "数据面板",
    title: "数据面板",
  },
  {
    key: "2",
    icon: () => h(DesktopOutlined),
    label: "系统信息",
    title: "Option 2",
  },
  {
    key: "3",
    icon: () => h(InboxOutlined),
    label: "定时任务",
    title: "Option 3",
  },
  {
    key: "sub1",
    icon: () => h(MailOutlined),
    label: "运维日志",
    title: "运维日志",
    children: [
      {
        key: "5",
        label: "运维日志1",
        title: "运维日志",
      },
      {
        key: "6",
        label: "运维日志2",
        title: "运维日志",
      },
    ],
  },
  {
    key: "sub2",
    icon: () => h(AppstoreOutlined),
    label: "小工具",
    title: "小工具",
    children: [
      {
        key: "9",
        label: "小工具1",
        title: "小工具1",
      },
      {
        key: "10",
        label: "小工具2",
        title: "小工具2",
      },
      {
        key: "sub3",
        label: "小工具3",
        title: "小工具3",
      },
    ],
  },
]);
watch(
  () => state.openKeys,
  (_val, oldVal) => {
    state.preOpenKeys = oldVal;
  }
);
const toggleCollapsed = () => {
  state.collapsed = !state.collapsed;
  state.openKeys = state.collapsed ? [] : state.preOpenKeys;
};
</script>
<style lang="scss">
.triggerCount {
  position: absolute;
  top: 0;
  right: -60px;
  width: 60px;
  height: 40px;
}
</style>
