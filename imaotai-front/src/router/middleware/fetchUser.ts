import { Middleware, MiddlewareContext } from "oh-router";
import { router } from "..";
import { getToken } from "../../shared/token";
import { userStore } from "../../stores/user";
import { jwtDecode } from "jwt-decode";

export class FetchUserMiddleware extends Middleware {
  async handler(
    ctx: MiddlewareContext<{}>,
    next: () => Promise<any>
  ): Promise<void> {
    const token = getToken();
    if (token && userStore.user === undefined) {
      const decoded = jwtDecode(token);
      //获得user
      await userStore.fetchUser(String(decoded.mobile));
      router.rematch();
    } else {
      next();
    }
  }
}
