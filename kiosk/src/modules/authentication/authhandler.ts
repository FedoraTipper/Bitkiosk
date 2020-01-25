import {
  authPayload,
  LoginDetails
} from "@/modules/authentication/authdetails";
import Axios from "axios";
import Cookies from "js-cookie";
import {CookieDetails} from "@/utils/cookie/cookiedetails";
import cookieutil from "@/utils/cookie/cookieutil";

const config = require("@/utils/config/config");

export default class AuthHandler {
  constructor() {}

  // @ts-ignore
  async Login(details: LoginDetails): Promise<boolean> {
    await this.postLogin(details).then(authPayload => {
      return new Promise<boolean>(resolve => {

      });
    });
  }

  Register(): boolean {
    return false;
  }

  // @ts-ignore
  async postLogin(details: LoginDetails): Promise<T> {
    await Axios.post(
      config.default.BASE_PATH + config.default.PATH.AUTHENTICATION,
      {
        identification: details.identification,
        token: details.token,
        authMethodId: details.authMethodId
      },
      {
        withCredentials: true
      }
    )
      .then(response => {
        console.log(response);
        // let authorizationHeader = response.headers['authorization'];
        // let cookieToSet : CookieDetails = new cookieutil().convertHeaderToCookieJson(authorizationHeader) as CookieDetails;
        // console.log(cookieToSet)
        // Cookies.set('Authorization', cookieToSet.value, {path: cookieToSet.path, domain: cookieToSet.domain, expires: cookieToSet.expires, secure: true, sameSite: "lax"});
        return response;
      })
      .catch(error => {
        if (error.response.status === 401) {
          return error.response;
        }
      });
  }
}
