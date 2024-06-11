import { request } from "../shared/request";
export async function GetFollowerUserList(opts: {
  page: Number;
  page_size: Number;
}) {
  return request.post<any, any>("/imaotai/user/getFlowerUserList", opts);
}

export async function DeleteFollowerUser(opts: { uid: string }) {
  return request.post<any, any>("/imaotai/user/deleteFollowerUser", opts);
}

export async function StartFollowerUser(opts: { uid: string }) {
  return request.post<any, any>("/imaotai/user/startFollowerUser", opts);
}

export async function SuspendFollowerUser(opts: { uid: string }) {
  return request.post<any, any>("/imaotai/user/suspendFollowerUser", opts);
}

export async function GetVerifyCode(opts: { uid: string }) {
  return request.post<any, any>("/imaotai/user/getVerifyCode", opts);
}
export async function UpdateFollowerUserToken(opts: {
  uid: string;
  code: string;
}) {
  return request.post<any, any>("/imaotai/user/updateFollowerUserToken", opts);
}
export async function UpdateFollowerUserAddress(opts: {
  uid: string;
  address: string;
}) {
  return request.post<any, any>("/imaotai/user/updateFollowerUserAddress", opts);
}
