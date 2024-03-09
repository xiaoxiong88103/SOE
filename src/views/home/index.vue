<template>
  <div class="home">
    <Charts :options="options" />
    <Charts :options="options" />
  </div>
</template>

<script>
import { reactive, toRefs, ref } from "vue";
import Charts from "@/components/charts.vue";
import home from "@/mock/home.js";
console.log(home);
export default {
  setup() {
    const state = reactive({
      count: 0,
    });
    const options = ref(home.bigMaster);
    return {
      ...toRefs(state),
      options,
    };
  },
  components: { Charts },
};
</script>

<style lang="scss" scoped>
.home {
  width: 100%;
  height: 100%;
}
</style>
