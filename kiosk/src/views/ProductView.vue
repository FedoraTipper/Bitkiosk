<template>
  <div v-if="product != undefined">
    <ProductComponent :product="product" :card-form="false" />
  </div>
</template>

<script>
import Product from "@/models/product";
import ProductComponent from "@/components/Product";
import { ProductsModule } from "@/store/modules/products";
import Component, { mixins } from "vue-class-component";
import { AuthMixin } from "@/mixins/authmixin";
import {Watch} from "vue-property-decorator";

@Component({
  components: {
    ProductComponent
  }
})
export default class ProductView extends mixins(AuthMixin) {
  product = new Product();

  constructor() {
    super();
  }

  created() {
    console.log("asdasd")
    if (ProductsModule.products.length == 0) {
      console.log("asdasdasdasd");
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
        console.log("prodcut found");
        this.product = p;
      }
    });
  }
}
</script>
