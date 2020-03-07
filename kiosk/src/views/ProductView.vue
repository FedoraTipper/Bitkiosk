<template>
  <div>
    <div v-for="product in products" v-bind:key="product.SKU">
      <Product :product="product" :card-form="true"></Product>
    </div>
  </div>
</template>

<script>
import Product from "@/components/Product";
import { ProductsModule } from "@/store/modules/products";
import Component, {mixins} from "vue-class-component";
import { AuthMixin } from "@/mixins/authmixin";

@Component({
  components: {
    Product
  }
})
export default class ProductView extends mixins(AuthMixin) {
  constructor() {
    super();
  }

  mounted() {
    ProductsModule.loadActiveProducts();
  }

  get products() {
    return ProductsModule.products;
  }
}
</script>
