<script lang="jsx">
import { defineComponent } from "vue";
// import TestJsx from "@/components/test.jsx";
// import TestVue from "@/components/test.vue";
export default defineComponent({
  setup(props, ctx) {
    return () => (
      <div className="main">
        {/* <TestJsx />
        <TestVue /> */}
        <router-view></router-view>
      </div>
    );
  },
});
</script>
<style>
.main{
  width: 100VW;
  height: 100VH;
}
</style>