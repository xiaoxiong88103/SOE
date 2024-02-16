import React from "react";
import { Routes, Route } from "react-router-dom";
import { useNavigate } from "react-router-dom";
import { Watermark } from "antd";
import RouterList from "../router/router.js";
// import { useSelector, useDispatch } from "react-redux";
// import { decrement, increment } from "../store/system";

// 页面
// import {
//   LaptopOutlined,
//   NotificationOutlined,
//   UserOutlined,
// } from "@ant-design/icons";
import { Breadcrumb, Layout, Menu, theme } from "antd";
const { Header, Content, Sider } = Layout;
const LeftMuenList = RouterList.Index.children.map((it) => {
  return {
    key: RouterList.Index.path + "/" + it.path,
    label: it.key,
  };
});
const Main = () => {
  const navigate = useNavigate();
  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken();
  const CheckRouter = (el) => {
    console.log(navigate);
    navigate(el.key);
  };
  return (
    <Watermark content="测试水印">
      <Layout style={{ height: "100vh" }}>
        <Header
          style={{
            display: "flex",
            alignItems: "center",
          }}
        >
          {/* <Menu
            theme="dark"
            mode="horizontal"
            defaultSelectedKeys={["1"]}
            items={items1}
            style={{
              flex: 1,
              minWidth: 0,
            }}
          /> */}
          header
        </Header>
        <Layout>
          <Sider
            width={200}
            style={{
              background: colorBgContainer,
            }}
          >
            <Menu
              onClick={(el) => {
                CheckRouter(el);
              }}
              mode="inline"
              defaultSelectedKeys={["1"]}
              defaultOpenKeys={["sub1"]}
              style={{
                height: "100%",
                borderRight: 0,
              }}
              items={LeftMuenList}
            />
          </Sider>
          <Layout
            style={{
              padding: "0 24px 24px",
            }}
          >
            <Breadcrumb
              style={{
                margin: "16px 0",
              }}
            >
              <Breadcrumb.Item>Home</Breadcrumb.Item>
              <Breadcrumb.Item>List</Breadcrumb.Item>
              <Breadcrumb.Item>App</Breadcrumb.Item>
            </Breadcrumb>
            <Content
              style={{
                padding: 24,
                margin: 0,
                minHeight: 280,
                background: colorBgContainer,
                borderRadius: borderRadiusLG,
              }}
            >
              {/* <button
                aria-label="Increment value"
                onClick={() => dispatch(increment())}
              >
                Increment
              </button>
              <span>{count}</span>
              <button
                aria-label="Decrement value"
                onClick={() => dispatch(decrement())}
              >
                Decrement
              </button> */}
              {/* {JSON.stringify(RouterList.Index)} */}
              {/* {JSON.stringify(RouterList.Index)} */}
              {/* {GetRouterList.Index.children.map((it) => {
                return 1;
              })} */}
              <Routes>
                {RouterList.Index.children.map((it, i) => {
                  return (
                    <Route
                      path={`/${it.path}`}
                      key={i}
                      element={it.component}
                    />
                  );
                })}
              </Routes>
            </Content>
          </Layout>
        </Layout>
      </Layout>
    </Watermark>
  );
};

export default Main;
