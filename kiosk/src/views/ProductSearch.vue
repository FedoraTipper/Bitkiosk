<template>
  <div class="columns is-multiline is-centered is-mobile" style="margin:10px 10px 0 10px">
    <div class="column is-one-fifth-desktop is-full-mobile is-one-third-tablet" style="width: 300px;"
      v-for="product in products"
      v-bind:key="product.SKU"
    >
      <Product :product="product" />
    </div>
  </div>
</template>

<script>
import Product from "@/components/ProductCard";
import { ProductsModule } from "@/store/modules/products";
import Component, {mixins} from "vue-class-component";
import { AuthMixin } from "@/mixins/authmixin";

@Component({
  components: {
    Product
  }
})
export default class ProductSearch extends mixins(AuthMixin) {
  constructor() {
    super();
  }

  created() {
    ProductsModule.loadActiveProducts();
  }

  get products() {
    return ProductsModule.products;
  }
}
</script>
