import { RouteObject, Router } from "oh-router";
import { BasicLayout } from "../layouts/basic";
import Index from "../pages/Index";
import User from "../pages/User";
import WineInfo from "../pages/WineInfo";
import RecordInfo from "../pages/RecordInfo";
import { Admin, SuperAdmin } from "../shared/constant";
import { LoginCheckMiddleWare } from "./middleware/loginChcek";
import { FetchUserMiddleware } from "./middleware/fetchUser";
import HomePage from "../pages/HomePage";

export interface Meta {
  // 1是superAdmin,2是admin
  role?: (1 | 2)[];
  menu?: {
    name: string;
    path: string;
  };
}

export const router = new Router({
  middlewares: [new LoginCheckMiddleWare(), new FetchUserMiddleware()],
  routes: [
    {
      path: "/index",
      element: <Index />,
    },
    {
      path: "/",
      name: "base",
      element: <BasicLayout />,
      children: [],
    },
    {
      path: "/user",
    },
    {
      path: "/homepage",
    },
    {
      path: "/wineInfo",
    },
    {
      path: "/RecordInfo",
    },
    {
      paht: "*",
      element: "404",
    },
  ],
});

const permissionRoutes: RouteObject<Meta>[] = [
  {
    // Index: true,
    path: "/homepage",
    element: <HomePage />,
    meta: {
      menu: {
        name: "首页",
        path: "/homepage",
      },
    },
  },
  {
    path: "/user",
    element: <User />,
    meta: {
      role: [Admin],
      menu: {
        name: "用户管理",
        path: "/user",
      },
    },
  },
  {
    path: "/wineInfo",
    element: <WineInfo />,
    meta: {
      role: [SuperAdmin, Admin],
      menu: {
        name: "今日酒价(待开发)",
        path: "/wineInfo",
      },
    },
  },
  {
    path: "/RecordInfo",
    element: <RecordInfo />,
    meta: {
      role: [SuperAdmin, Admin],
      menu: {
        name: "中奖记录(待开发)",
        path: "/RecordInfo",
      },
    },
  },
];

export function resetPermissionRoutes(user?: any) {
  const baseRoute = router.getRoutes().find((route) => route.name === "base")!;
  if (user) {
    const tmpRoutes: RouteObject<Meta>[] = [];
    for (const route of permissionRoutes) {
      if (route.meta?.role && !route.meta.role.includes(user.role)) continue;
      tmpRoutes.push(route);
    }
    baseRoute.children = tmpRoutes;
  } else {
    baseRoute.children = [];
  }
}
