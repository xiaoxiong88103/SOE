import { configureStore } from "@reduxjs/toolkit";
// 系统保留控制数据中心
import system from "./store/system";
export default configureStore({
  reducer: {
    system,
  },
});
