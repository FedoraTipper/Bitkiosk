import {
  VuexModule,
  Module,
  getModule,
  Mutation
} from "vuex-module-decorators";
import store from "@/store";
import ProductsAPI from "@/modules/graphql/products/productsgql";
import Product from "@/models/product";

export interface IProductsState {
  products: Array<Product>;
}

@Module({
  dynamic: true,
  namespaced: true,
  name: "Products",
  store
})
class Products extends VuexModule implements IProductsState {
  products: Array<Product> = new Array<Product>();

  @Mutation
  async loadActiveProducts() {
    new ProductsAPI().fetchActiveProducts().then(result => {
      this.products = result;
    });
  }
}

export const ProductsModule = getModule(Products);
