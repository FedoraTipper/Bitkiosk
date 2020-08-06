import ProductQueries from "./productqueries";
import Product from "@/models/product";
import gqlfactory from "@/utils/gqlclient/gqlfactory";
import User from "@/models/user";

export default class ProductsGQL {
  async fetchActiveProducts(): Promise<Array<Product>> {
    return new Promise<Array<Product>>(async resolve => {
      let activeProducts: Array<Product> = new Array<Product>();

      let GQLClient = new gqlfactory().newGQLClient();

      await GQLClient.request(ProductQueries.loadActiveProducts, {})
        .then((response: any) => {
          if (response) {
            console.log(response);
            for (let productResponse of response["loadActiveProducts"]) {
              let product: Product = new Product();
              product.setProductFromResponseObject(productResponse);

              let user: User = new User();
              activeProducts.push(product);
            }
          }
        })
        .catch(error => {
          // TODO: fix error logging
        });

      resolve(activeProducts);
    });
  }
}
