<template>
  <div>

  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import Product from "@/models/product";

@Component
export default class RegisterForm extends Vue {
  constructor() {
    super();
  }

  @Prop({ type: Product }) readonly product!: Product;
}
</script>
