import {
  LoginDetails,
  authPayload
} from "@/modules/authentication/authdetails";
import Axios from "axios";
import NotificationUtil from "@/utils/notification/notificationutil";

const config = require("@/utils/config/config");

export default class AuthHandler {
  constructor() {}

  // @ts-ignore
  async Login(details: LoginDetails): Promise<boolean> {
    return new Promise<boolean>(async resolve => {
      await this.postLogin(details).then((response: authPayload) => {
        // eslint-disable-next-line no-empty
        if (response.error.length > 0) {
          new NotificationUtil().displayError("Incorrect login details");
          resolve(false);
        }

        resolve(true);
      });
    });
  }

  Register(): boolean {
    return false;
  }

  // @ts-ignore
  async postLogin(details: LoginDetails): Promise<authPayload> {
    return new Promise(async resolve => {
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
          resolve(response.data);
        })
        .catch(error => {
          if (error.response.status === 401) {
            console.log(error.response.data);
            resolve(error.response.data);
          }
        });
    });
  }
}
