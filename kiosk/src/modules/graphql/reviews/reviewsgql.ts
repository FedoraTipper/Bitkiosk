import ReviewQueries from "./reviewqueries";
import GQLFactory from "@/utils/gqlclient/gqlfactory";
import Review from "@/models/review";

export default class ReviewsGQL {
  async fetchReviewsForProductWithSku(productSku: string): Promise<Array<Review>> {
    return new Promise<Array<Review>>(async (resolve, reject) => {
      let reviewsForProduct: Array<Review> = new Array<Review>();

      let GQLClient = new GQLFactory().newGQLClient();
      await GQLClient.request(ReviewQueries.loadReviewsForProduct, {
        productSku
      })
        .then((response: any) => {
          if (response) {
            console.log(response);
            for (let productResponse of response["loadReviewsForProduct"]) {
              let review: Review = new Review();
              review.setReviewFromResponseObject(productResponse);

              reviewsForProduct.push(review);
            }
          }
        })
        .catch(error => {
          // TODO: Fix logging error
          console.log(error);
          reject("Unable to fetch reviews for product");
        });

      resolve(reviewsForProduct);
    });
  }
}
