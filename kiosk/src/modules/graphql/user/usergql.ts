import { UserProfile } from "../../../models/userprofile";
import gqlfactory from "@/utils/gqlclient/gqlfactory";
import UserQueries from "@/modules/graphql/queries/userqueries";
import NotificationUtil from "@/utils/notification/notificationutil";
import {RegisterDetails} from "@/models/authentication/authdetails";

export default class Usergql {
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
            response = response['userProfile'];
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

  async registerUser(registerDetails: RegisterDetails): Promise<boolean> {
    return new Promise<boolean>(async resolve => {
      let signUpDetails =
    });
  }
}
