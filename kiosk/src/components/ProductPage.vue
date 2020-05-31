<template>
  <div class="section">
    <div class="container is-fluid box columns">
      <div class="column is-two-fifths">
        <figure class="image is-4by3 is-half">
          <img
            src="https://bulma.io/images/placeholders/1280x960.png"
            alt="Placeholder image"
          />
        </figure>
      </div>
      <div class="column hero is-three-fifths">
        <div class="hero-body">
          <div class="title level-left">
            {{ product.name }}
          </div>
          <div class="level-left">Stock remaining: {{ product.stock }}</div>
          <div class="level-left subtitle has-text-weight-bold is-size-4">
            ${{ product.price }}
          </div>
          <div class="level-left" style="margin-top: 10px">
            <b-field>
              <b-rate
                v-model="reviewRating"
                icon="star"
                :max="5"
                :show-score="true"
                :rtl="false"
                :disabled="true"
              />
              <p style="margin-left: 10px; font-size: 14px; margin-top: 2px">
                3000 Reviews
              </p>
            </b-field>
          </div>
        </div>
        <div class="hero-foot">
          <div class="level-item">
            <b-numberinput :v-model="quantity"
              controls-position="compact"
              size="is-small"
              type="is-warning"
              style="max-width: 150px"
              :editable="false"
              min="1"
            />
          </div>
          <div class="level-item">
            <b-button
              class="hero-buttons"
              @click="addToCart"
              type="is-danger"
              size="is-medium"
              :loading="getLoadingStatus()"
              icon-left="basket"
            >Add to cart</b-button>
          </div>
        </div>
      </div>
    </div>

    <div class="container is-fluid box level">
      <div class="level-left">
        <div class="title">
          Description
        </div>
        <div class="subtitle">
          {{ product.description }}
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import Product from "@/models/product";

@Component
export default class ProductPage extends Vue {
  private loadingStatus: boolean = false;
  private reviewRating: number = 3.1;
  private quantity: number = 1;

  constructor() {
    super();
  }

  private addToCart() {
    this.setLoadingStatus(true);
  }

  private setLoadingStatus(status: boolean) {
    this.loadingStatus = status;
  }

  private getLoadingStatus(): boolean {
    return this.loadingStatus;
  }

  @Prop({ type: Product }) readonly product!: Product;
}
</script>
