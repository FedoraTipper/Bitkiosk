export default {
  createUser: `
    mutation RegisterNewUser($input: NewUser!) {
      createUser(input: $input)
    }
  `,
  signUpNewUser: `
    mutation signUpNewUser($input: NewUser!) {
        createUser(input: $input)
    }
   `
};