export default {
  createUser: `
    mutation RegisterNewUser($input: NewUser!) {
      createUser(input: $input)
    }
  `
};