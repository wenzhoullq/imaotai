import { ProLayout } from "@ant-design/pro-layout";
import { Button, Space, Typography } from "antd";
import { Link, Outlet, useLocation } from "oh-router-react";
import { useMemo } from "react";
import { router } from "../router";
import { userStore } from "../stores/user";
const { Text } = Typography;

export function BasicLayout() {
  const location = useLocation();

  const menus = useMemo(() => {
    return router
      .getRoutes()
      .find((route) => route.name === "base")!
      .children?.map((route) => {
        return route.meta?.menu;
      });
  }, []);

  return (
    <ProLayout
      logo={"../public/logo.svg"}
      title="imaotai小助手"
      location={location}
      style={{
        height: "100vh",
      }}
      rightContentRender={() => {
        return (
          <Space>
            <span>
              <>
                <Text strong>
                  {"用户："}
                  {userStore.user.mobile}
                </Text>
              </>
            </span>
            <Button
              type="primary"
              onClick={() => {
                userStore.logout();
                router.navigate("/index");
              }}
            >
              注销
            </Button>
          </Space>
        );
      }}
      menuItemRender={(it: any) => {
        return <Link to={it.path}>{it.name}</Link>;
      }}
      route={{
        path: "/",
        routes: menus,
      }}
    >
      <Outlet />
    </ProLayout>
  );
}
