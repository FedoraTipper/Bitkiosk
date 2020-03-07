import { UserProfile } from "../../../models/userprofile";
import gqlfactory from "@/utils/gqlclient/gqlfactory";
import UserQueries from "@/modules/graphql/user/userqueries";
import UserMutations from "@/modules/graphql/user/usermutations";
import NotificationUtil from "@/utils/notification/notificationutil";
import {
  IRegisterDetails,
  IRegisterResponse
} from "@/models/authentication/authdetails";
import GQLClientFactory from "@/utils/gqlclient/gqlfactory";
import User from "@/models/user";

export default class UserGQL {
  constructor() {}

  async fetchUserProfile(email: string | null): Promise<User> {
    return new Promise<User>(async resolve => {
      let user: User = new User();

      let GQLClient = new gqlfactory().newGQLClient();

      const inputData = {
        email: email
      };

      await GQLClient.request(UserQueries.getUserProfile, inputData)
        .then(response => {
          if (response) {
            console.log(response);
            user = new User();
            user.setUserFromResponseObject(response["userProfile"]);
          }
        })
        .catch(error => {
          // TODO ADD ENV LOGGER
        });

      resolve(user);
    });
  }

  async registerUser(
    registerDetails: IRegisterDetails
  ): Promise<IRegisterResponse> {
    return new Promise<IRegisterResponse>(async resolve => {
      let GQLClient = new GQLClientFactory().newGQLClient();
      await GQLClient.request(UserMutations.signUpNewUser, {
        input: registerDetails
      })
        .then(response => {
          resolve({ success: true, message: "" } as IRegisterResponse);
        })
        .catch(response => {
          let error = JSON.parse(JSON.stringify(response, undefined, 2))[
            "response"
          ];
          if (error == undefined) {
            new NotificationUtil().displayError("Unable to parse response :(");
            resolve({success: false, message: ""})
          }

          resolve({
            success: false,
            message: (() => {
              let errorMessages: string[] = [];
              error["errors"].forEach((err: { [x: string]: string }) => {
                errorMessages.push(err["message"]);
              });
              return errorMessages.join("\n");
            })()
          } as IRegisterResponse);
        });
    });
  }
}
