import { defineStore } from "pinia";

const useStore = defineStore("storeId", {
  state: () => {
    return {
      count: 1,
      arr: [], 
    };
  },
  getters: {},
  actions: {},
});

export default useStore;
