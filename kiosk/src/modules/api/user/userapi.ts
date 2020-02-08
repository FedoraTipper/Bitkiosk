import { UserProfile } from "../../../models/userprofile";
import gqlfactory from "@/utils/gqlclient/gqlfactory";
import UserQueries from "@/modules/api/user/userqueries";
import NotificationUtil from "@/utils/notification/notificationutil";

export default class UserAPI {
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
}
