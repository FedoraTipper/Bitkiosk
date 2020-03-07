<template>
  <div>
    <div class="card" v-if="cardForm">
      <div class="card-content">
        <div class="content">
          {{ product.name }}
          {{ product.price }}
          {{ product.description }}
        </div>
      </div>
      <footer class="card-footer">
        <a href="#" class="card-footer-item">Add to cart</a>
      </footer>
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
