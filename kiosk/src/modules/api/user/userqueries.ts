export default {
  getUserProfile: `
      query getUserProfile($email: String!) {
        userProfile(email: $email) {
          firstName
          lastName
          dateOfBirth
        }
      }
   `
};
