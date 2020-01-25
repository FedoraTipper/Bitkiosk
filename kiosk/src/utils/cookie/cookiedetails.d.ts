export const SAMESITE_LAX = "lax";
export const SAMESITE_STRICT = "strict";

export interface CookieDetails {
  value: string;
  path: string;
  expires: number;
  domain: string;
  secure: boolean;
  sameSite: string;
}
