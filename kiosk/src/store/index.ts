import Vue from "vue";
import Vuex from "vuex";

import { IUserState } from "./modules/user";
import { IProductsState } from "./modules/products";

Vue.use(Vuex);

export interface IRootState {
  user: IUserState;
  products: IProductsState;
}

export default new Vuex.Store<IRootState>({});
