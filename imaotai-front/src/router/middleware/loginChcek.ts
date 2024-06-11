import { Middleware, MiddlewareContext } from "oh-router";
import { getToken } from "../../shared/token";
import { router } from "..";

export class LoginCheckMiddleWare extends Middleware {
  async handler(
    ctx: MiddlewareContext<{}>,
    next: () => Promise<any>
  ): Promise<void> {
    // 判断用户是否登录
    const token = getToken();
    if (ctx.to.pathname === "/index") {
      if (token) {
        router.navigate("/");
      } else {
        next();
      }
      return;
    }
    if (token) {
      next();
    } else {
      router.navigate("/index");
    }
  }
}
