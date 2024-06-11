import { message } from "antd";
import { login, getAdminInfo } from "../api/index";
import { resetPermissionRoutes } from "../router";
import { removeToken, setToken } from "../shared/token";
import { NavLink } from "oh-router-react";
export const userStore = new (class {
  async login(mobile: string, pass_word: string) {
    const resp = await login({ mobile, pass_word });
    setToken(resp.headers.authorization);
  }
  user:
    | {
        mobile: string;
        uid: string;
        role: number;
      }
    | undefined;

  async fetchUser(mobile: string) {
    await getAdminInfo({ mobile }).then((data) => {
      if (data.data.err_no !== 0) {
        message.error(data.data.err_msg);
        return;
      }
      this.user = {
        mobile: data.data.data.mobile,
        uid: data.data.data.uid,
        role: data.data.data.role,
      };
    });
    resetPermissionRoutes(this.user);
  }

  logout() {
    removeToken();
    this.user = undefined;
    resetPermissionRoutes();
  }
})();
