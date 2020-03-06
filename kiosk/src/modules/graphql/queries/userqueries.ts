export default {
  getUserProfile: `
      query getUserProfile($email: String!) {
        userProfile(email: $email) {
          firstName
          lastName
        }
      }
   `,
  signUpNewUser: `
    mutation signUpNewUser($input: NewUser!) {
        createUser(input: $input)
    }
   `
};
