import {
  VuexModule,
  Module,
  getModule,
  MutationAction,
  Mutation,
  Action
} from "vuex-module-decorators";
import store from "@/store";
import { UserProfile } from "@/models/userprofile.ts";
import ProductsAPI from "@/modules/graphql/products/productsgql";
import NotificationUtil from "@/utils/notification/notificationutil";
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
