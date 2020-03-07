import ProductQueries from "./productqueries";
import Product from "@/models/product";
import IProduct from "@/models/product.d.ts";
import gqlfactory from "@/utils/gqlclient/gqlfactory";

export default class ProductsGQL {
  async fetchActiveProducts(): Promise<Array<Product>> {
    return new Promise<Array<Product>>(async resolve => {
      let activeProducts: Array<Product> = new Array<Product>();

      let GQLClient = new gqlfactory().newGQLClient();

      await GQLClient.request(ProductQueries.loadActiveProducts, {})
        .then(response => {
          if (response) {
            for (let productResponse of response["loadActiveProducts"]) {
              let product: Product = new Product();
              product.setProductFromResponseObject(productResponse);
              activeProducts.push(product);
            }
          }
        })
        .catch(error => {});

      resolve(activeProducts);
    });
  }
}
