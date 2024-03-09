<template>
  <div ref="echart" class="echart"></div>
</template>

<script>
import * as echarts from "echarts";
export default {
  props:{
    options: {
      type: Object,
      default:  {
        xAxis: {
          data: ["1月", "2月", "3月"],
        },
        yAxis: {
          type: "value",
        },
        series: [
          {
            name: "销量",
            type: "bar",
            data: [100, 200, 300],
          },
        ],
      },
    }
  },
  data() {
    return {
      Observer: "", //resize监听器
      timer: "", //定时器
      echartElement: "", //echart实例
    };
  },

  mounted() {
    //this.loadEchart()
    //不需要执行上一行的原因是，下面的监听器会在挂载的时候自动执行一次。
    this.Observer = new ResizeObserver((entries) => {
      if (this.timer) {
        clearTimeout(this.timer);
      }
      this.timer = setTimeout(this.loadEchart(), 50);
    });
    this.Observer.observe(this.$refs.echart);
  },

  beforeDestory() {
    clearTimeout(this.timer);
    /*
		清除定时器。定时任务不会随着组件销毁而销毁，所以组件销毁后，
		如果还有定时任务没有执行，就会继续调用loadEchart函数，
		而此时this.$refs.echart是undefined，echart就会报错：
		“Initialize failed: invalid dom”，意思就是“初始化失败：无效的dom”
		*/
    this.Observer.unobserve(this.$refs.echart);
  },

  methods: {
    //加载echart
    loadEchart() {
      if (this.echartElement) {
        //如果echart已经初始化过，需要销毁，否则会报错：重复初始化
        this.echartElement.dispose();
      }
      this.echartElement = echarts.init(this.$refs.echart);
      this.echartElement.setOption(this.options);
    },
  },
};
</script>
<style scoped>
.echart {
  width: 100%;
  height: 100%;
}
</style>
