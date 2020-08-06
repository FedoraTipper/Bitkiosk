export default {
  loadReviewsForProduct: `
    query getReviewsForProduct($productSku: String!) {
          loadReviewsForProduct (productSku: $productSku) {
            userName
            productSku
            textReview
            rating
            createAt
          }
    } 
    `,
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
}