export default {
  loadReviewsForProduct: `
    query getReviewsForProduct($productSku: String!) {
          loadReviewsForProduct (productSku: $productSku) {
            userDisplayName
            productSku
            textReview
            rating
            createdAt
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