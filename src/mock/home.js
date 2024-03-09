export default {
  bigMaster: {
    color: [
      "#d666e498",
      "#ff67af98",
      "#ff8f7b98",
      "#44465598",
      "#888eba98",
      "#00b0ff98",
      "#6dfacd98",
      "#f0255798",
      "#f5757698",
    ],
    tooltip: {
      trigger: "item",
      formatter: "{a} <br/>{b}: {c} ({d}%)",
    },
    legend: {
      data: [
        "Master节点",
        "node1节点",
        "node2节点",
        "node3节点",
        "node4节点",
        "node5节点",
        "node6节点",
        "node7节点",
        "node8节点",
        "node9节点",
      ],
    },
    series: [
      {
        name: "节点数据流",
        type: "pie",
        selectedMode: "single",
        radius: [0, "30%"],
        label: {
          position: "inner",
          color: "#444655",
          fontWeight: "bold",
          fontSize: 14,
        },
        labelLine: {
          show: false,
        },
        data: [
          { value: 1548, name: "node2节点" },
          { value: 775, name: "Master" },
          { value: 679, name: "node1节点", selected: true },
        ],
      },
      {
        name: "节点数据流",
        type: "pie",
        radius: ["45%", "60%"],
        labelLine: {
          length: 30,
        },
        label: {
          formatter: "{a|{a}}{abg|}\n{hr|}\n  {b|{b}:}{c}  {per|{d}%}  ",
          backgroundColor: "#8685ef",
          borderColor: "#F6F8FC",
          color: "#FFFFFF",
          borderWidth: 1,
          borderRadius: 5,
          rich: {
            a: {
              color: "#fff",
              lineHeight: 22,
              align: "center",
            },
            hr: {
              borderColor: "#fff",
              width: "80%",
              align: "left",
              borderWidth: 0.2,
              height: 0,
            },
            b: {
              color: "#4C5058",
              color: "#fff",
              fontSize: 14,
              lineHeight: 33,
            },
            per: {
              color: "#fff",
              backgroundColor: "#4C5058",
              padding: [3, 4],
              borderRadius: 4,
            },
          },
        },
        data: [
          { value: 1048, name: "node6节点" },
          { value: 335, name: "Master" },
          { value: 310, name: "node3节点" },
          { value: 251, name: "node7节点" },
          { value: 234, name: "node4节点" },
          { value: 147, name: "node8节点" },
          { value: 135, name: "node5节点" },
          { value: 102, name: "node9节点" },
        ],
      },
    ],
  },
};
