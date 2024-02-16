import LoginViews from "../views/login.js";
import Main from "../views/main.js";
import React from "react";
// hola
import Hola from "../views/main_page/hola.jsx";
// 监控
import SysChartSize from "../views/main_page/sys_chart.jsx";
// 工具
import SystemTools from "../views/main_page/sys_tools.jsx";
// shell类目
import ShellTerminal from "../views/main_page/sys_shell_ter.jsx";
import ShellScriptEdit from "../views/main_page/sys_xshell_exit.jsx";
// 系统
import SystemLog from "../views/main_page/sys_log.jsx";
let Index = {
  path: "/main",
  element: <Main></Main>,
  children: [
    {
      path: "hola",
      key: "欢迎",
      component: <Hola></Hola>,
    },
    {
      path: "systemchart",
      key: "系统监控",
      component: <SysChartSize></SysChartSize>,
    },
    {
      path: "terminal",
      key: "原生终端",
      component: <ShellTerminal></ShellTerminal>,
    },
    {
      path: "scriptedit",
      key: "脚本编辑",
      component: <ShellScriptEdit></ShellScriptEdit>,
    },
    {
      path: "tools",
      key: "系统工具",
      component: <SystemTools></SystemTools>,
    },
    {
      path: "log",
      key: "系统日志",
      component: <SystemLog></SystemLog>,
    },
  ],
};

let Over = [
  // 404
  {
    path: "*",
    element: <span>404</span>,
  },
  // 登录
  {
    path: "/login",
    element: <LoginViews></LoginViews>,
  },
  {
    path: "/",
    element: <LoginViews></LoginViews>,
  },
];

let exp = { Index, Over };
export default exp;
