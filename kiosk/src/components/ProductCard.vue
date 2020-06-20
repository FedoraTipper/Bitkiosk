<template>
  <div>
    <div
      style="cursor: pointer"
      class="card"
      @click="
        pushToPageWithParams(routeDefinitions.productview.name, {
          sku: product.SKU
        })
      "
    >
      <div class="card-image">
        <figure class="image is-4by3">
          <img
            src="https://bulma.io/images/placeholders/1280x960.png"
            alt="Placeholder image"
          />
        </figure>
      </div>
      <header>
        <p class="card-header-title is-centered is-4 title">
          {{ product.name }}
        </p>
      </header>
      <div class="card-content">
        <div class="content">
          {{ product.shortDescription }}
        </div>
        <b-field>
          <b-rate
            v-model="product.rating"
            :max="5"
            :show-score="true"
            :rtl="false"
            :disabled="true"
          />
          <p style="margin-left: 10px; font-size: 14px; margin-top: 2px">
            {{ product.getReviewDisplay() }}
          </p>
        </b-field>
        <div class="content title is-5">${{ product.price }}</div>
      </div>
      <footer class="card-footer">
        <span class="card-footer-item">Quick add</span>
      </footer>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import Product from "@/models/product";

@Component
export default class ProductCard extends Vue {
  constructor() {
    super();
  }

  @Prop({ type: Product }) readonly product!: Product;
}
</script>
