import { message } from "antd";
import axios, { AxiosError } from "axios";
import { getToken, removeToken, setToken } from "./token";
import { router } from "../router";
import { userStore } from "../stores/user";
export const request = axios.create({
  // baseURL:"http://152.136.223.254:8080/",
  baseURL: "http://localhost:8080/",
});

request.interceptors.request.use((config) => {
  const token = getToken();
  if (token) {
    config.headers!["Authorization"] = token;
    config.headers["Access-Control-Expose-Headers"] = "Authorization";
  }
  return config;
});

request.interceptors.response.use(
  (res) => {
    if (res.data.err_no !== 0) {
      // throw Error(res.data.err_msg);
    }
    return res;
  },
  (err: AxiosError) => {
    if (err.response?.status === 401) {
      userStore.logout();
      router.navigate("/index");
      message.error("请重新登录");
      return new Promise(() => {});
    }
    throw err;
  }
);
