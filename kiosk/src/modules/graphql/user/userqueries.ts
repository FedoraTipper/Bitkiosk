export default {
  getUserProfile: `
      query getUserProfile($email: String!) {
        userProfile(email: $email) {
          email
          role
          userProfile {
            firstName
            lastName
           }
        }
      }
   `
};
