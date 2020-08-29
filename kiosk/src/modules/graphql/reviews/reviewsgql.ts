import ReviewQueries from "./reviewqueries";
import GQLFactory from "@/utils/gqlclient/gqlfactory";
import Review from "@/models/review";

export default class ReviewsGQL {
  async fetchReviewsForProductWithSku(
    productSku: string,
    limit: number,
    offset: number
  ): Promise<Array<Review>> {
    return new Promise<Array<Review>>(async (resolve, reject) => {
      let reviewsForProduct: Array<Review> = new Array<Review>();

      let GQLClient = new GQLFactory().newGQLClient();
      await GQLClient.request(ReviewQueries.loadReviewsForProduct, {
        productSku: productSku,
        limit: limit,
        offset: offset
      })
        .then((response: any) => {
          if (response) {
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

  async fetchTotalReviewCountForProductWithSku(
    productSku: string
  ): Promise<number> {
    return new Promise<number>(async (resolve, reject) => {
      let GQLClient = new GQLFactory().newGQLClient();
      await GQLClient.request(
        ReviewQueries.loadTotalNumberOfReviewsForProduct,
        {
          productSku
        }
      )
        .then((response: any) => {
          if (response) {
            resolve(response["loadTotalNumberOfReviewsForProduct"]);
          }
        })
        .catch(error => {
          // TODO: Fix logging error
          console.log(error);
          reject("Unable to fetch number of reviews for product");
        });
    });
  }

  public fetchCurrentReviewForUser(
    productSku: string
  ): Promise<Review | undefined> {
    return new Promise<Review | undefined>(async (resolve, reject) => {
      let GQLClient = new GQLFactory().newGQLClient();
      await GQLClient.request(ReviewQueries.loadReviewForUserWithProductSku, {
        productSku
      })
        .then((response: any) => {
          let review: Review | undefined;
          console.log(response);
          if (response["loadReviewForUserWithProductSku"]) {
            review = new Review();
            review.setReviewFromResponseObject(
              response["loadReviewForUserWithProductSku"]
            );
          }

          resolve(review);
        })
        .catch(error => {
          // TODO: Fix logging error
          console.log(error);
          reject("Unable to fetch number of reviews for product");
        });
    });
  }
}
