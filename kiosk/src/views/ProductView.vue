<template>
  <div v-if="product != undefined">
    <ProductPage :product="product" />
  </div>
</template>

<script>
import Product from "@/models/product";
import ProductPage from "@/components/ProductPage";
import { ProductsModule } from "@/store/modules/products";
import Component, { mixins } from "vue-class-component";
import { AuthMixin } from "@/mixins/authmixin";
import {Watch} from "vue-property-decorator";

@Component({
  components: {
    ProductPage
  }
})
export default class ProductView extends mixins(AuthMixin) {
  product = new Product();

  constructor() {
    super();
  }

  created() {
    if (ProductsModule.products.length === 0) {
      ProductsModule.loadActiveProducts();
    }else {
      let SKUToFind = this.$route.params["sku"];
      ProductsModule.products.forEach(p => {
        if (p.SKU === SKUToFind) {
          this.product = p;
        }
      });
    }
  }

  get products() {
    return ProductsModule.products;
  }

  @Watch("products")
  loadProductFromList(val) {
    let SKUToFind = this.$route.params["sku"];
    ProductsModule.products.forEach(p => {
      if (p.SKU === SKUToFind) {
        this.product = p;
      }
    });
  }
}
</script>
