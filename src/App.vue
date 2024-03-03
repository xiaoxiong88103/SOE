<script lang="jsx">
import { defineComponent } from "vue";
import TestJsx from "@/components/test.jsx";
import TestVue from "@/components/test.vue";
export default defineComponent({
  setup(props, ctx) {
    return () => (
      <>
        <TestJsx />
        <TestVue />
        <router-view></router-view>
      </>
    );
  },
});
</script>
