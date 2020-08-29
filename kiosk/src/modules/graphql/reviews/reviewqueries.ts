export default {
  loadReviewsForProduct: `
    query getReviewsForProduct($productSku: String!, $limit: Int, $offset: Int) {
          loadReviewsForProduct (productSku: $productSku, limit: $limit, offset: $offset) {
            userDisplayName
            productSku
            textReview
            rating
            createdAt
          }
    } 
    `,
  loadTotalNumberOfReviewsForProduct: `
      query getTotalNumberOfReviewsForProduct($productSku: String!) {
          loadTotalNumberOfReviewsForProduct (productSku: $productSku)
    } 
    `,
  loadReviewForUserWithProductSku: `
      query loadReviewForUserWithProductSku($productSku: String!) {
          loadReviewForUserWithProductSku (productSku: $productSku) {
            userDisplayName
            productSku
            textReview
            rating
            anonymous
            createdAt
          }
    } 
    `
}