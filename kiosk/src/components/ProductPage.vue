<template>
  <div class="section">
    <div class="level columns">
      <div class="level-left column is-two-fifths">
        <figure class="image is-4by3 is-half">
          <img
            src="https://bulma.io/images/placeholders/1280x960.png"
            alt="Placeholder image"
          />
        </figure>
        <div class="title">
          {{ product.name }}
        </div>
      </div>
    </div>
    <div class="level">
      <div class="level-right"></div>
      {{ product.price }}
      {{ product.SKU }}
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import Product from "@/models/product";

@Component
export default class ProductPage extends Vue {
  constructor() {
    super();
  }

  @Prop({ type: Product }) readonly product!: Product;
}
</script>
