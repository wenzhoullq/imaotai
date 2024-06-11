import { request } from "../shared/request";

export async function login(opts: { mobile: string; pass_word: string }) {
  return request.post<any, any>("/imaotai/login", opts);
}

export async function register(opts: {
  mobile: string;
  pass_word: string;
  address: string;
  email: string;
}) {
  return request.post<any, any>("imaotai/register", opts);
}

export async function getAdminInfo(opts: { mobile: string }) {
  return request.post<any, any>("/imaotai/getAdminInfo", opts);
}
