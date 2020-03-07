<template>
  <div>
    <div class="card" v-if="cardForm" @click="pushToPageWithParams(routeDefinitions.productview.name, { sku: product.SKU })">
      <div class="card-image">
        <figure class="image is-4by3">
          <img src="https://bulma.io/images/placeholders/1280x960.png" alt="Placeholder image">
        </figure>
      </div>
      <header>
        <p class="card-header-title is-centered is-4 title">
          {{ product.name }}
        </p>
      </header>
      <div class="card-content">
        <div class="content">
          {{ product.description }}
        </div>
        <div class="content title is-5">
          ${{ product.price }}
        </div>
      </div>
      <footer class="card-footer">
        <span class="card-footer-item">Add to cart</span>
      </footer>
    </div>

    <div v-else class="section">
      {{ product.name }}
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import Product from "@/models/product";

@Component
export default class ProductComponent extends Vue {
  constructor() {
    super();
  }

  @Prop({ type: Product }) readonly product!: Product;
  @Prop({ default: false, type: Boolean }) readonly cardForm!: boolean;
}
</script>
