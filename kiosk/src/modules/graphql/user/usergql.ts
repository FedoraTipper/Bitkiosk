import { UserProfile } from "../../../models/userprofile";
import gqlfactory from "@/utils/gqlclient/gqlfactory";
import UserQueries from "@/modules/graphql/queries/userqueries";
import NotificationUtil from "@/utils/notification/notificationutil";
import {
  IRegisterDetails,
  IRegisterResponse
} from "@/models/authentication/authdetails";
import GQLClientFactory from "@/utils/gqlclient/gqlfactory";

export default class UserGQL {
  constructor() {}

  async fetchUserProfile(email: string | null): Promise<UserProfile> {
    return new Promise<UserProfile>(async resolve => {
      let userProfile: UserProfile = new UserProfile();

      let GQLClient = new gqlfactory().newGQLClient();

      const inputData = {
        email: email
      };

      await GQLClient.request(UserQueries.getUserProfile, inputData)
        .then(response => {
          if (response) {
            // userProfile = JSON.parse(response['userProfile']);
            response = response["userProfile"];
            userProfile = new UserProfile();
            userProfile.setUserProfileFromResponse(response);
          }
        })
        .catch(error => {
          // TODO ADD ENV LOGGER
        });

      resolve(userProfile);
    });
  }

  async registerUser(
    registerDetails: IRegisterDetails
  ): Promise<IRegisterResponse> {
    return new Promise<IRegisterResponse>(async resolve => {
      let GQLClient = new GQLClientFactory().newGQLClient();
      await GQLClient.request(UserQueries.signUpNewUser, {
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
