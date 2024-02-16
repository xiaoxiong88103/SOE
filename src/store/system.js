import { createSlice } from "@reduxjs/toolkit";
export const counterSlice = createSlice({
  name: "system",
  initialState: {
    value: 0,
  },
  reducers: {
    increment: (state) => {
      // Redux Toolkit 允许我们在 reducers 中编写 mutating 逻辑。
      // 它实际上并没有 mutate state 因为它使用了 Immer 库，
      // 它检测到草稿 state 的变化并产生一个全新的基于这些更改的不可变 state
      state.value += 1;
    },
    decrement: (state) => {
      state.value -= 1;
    },
    incrementByAmount: (state, action) => {
      state.value += action.payload;
    },
  },
});

// 为每个 case reducer 函数生成 Action creators
export const { increment, decrement, incrementByAmount } = counterSlice.actions;

export default counterSlice.reducer;
