<template>
  <div class="section">
    <div class="container box">
      <div class="columns">
        <div class="column is-two-fifths">
          <img
            src="https://bulma.io/images/placeholders/1280x960.png"
            alt="Placeholder image"
            style="max-height: 450px; width: auto"
          />
        </div>
        <div class="column hero is-three-fifths">
          <div class="hero-body">
            <div class="title level-left">
              {{ product.name }}
            </div>
            <div class="level-left">
              <b-taglist attached v-if="isProductInStock">
                <b-tag type="is-success">In stock</b-tag>
                <b-tag type="is-dark">{{ product.stock }} units</b-tag>
              </b-taglist>
              <!-- TODO: SUPPORT STOCK COMING SOON FEATURE -->
              <b-taglist attached v-else-if="!isProductInStock && false">
                <b-tag type="is-warning">Stock coming soon</b-tag>
              </b-taglist>
              <b-taglist attached v-else>
                <b-tag type="is-danger">Out of stock</b-tag>
              </b-taglist>
            </div>
            <div class="level-left subtitle has-text-weight-bold is-size-4" style="margin-top: 15px">
              ${{ product.price }}
            </div>
            <div class="level-left" style="margin-top: 10px">
              <b-field>
                <b-rate
                  v-model="product.rating"
                  :max="5"
                  :show-score="true"
                  :rtl="false"
                  :disabled="true"
                />
                <p style="margin-left: 10px; font-size: 14px; margin-top: 2px">
                  {{product.getReviewDisplay()}}
                </p>
              </b-field>
            </div>
          </div>
          <div class="hero-foot">
            <div class="level-item">
              <b-numberinput
                :v-model="quantity"
                controls-position="compact"
                size="is-small"
                type="is-warning"
                style="max-width: 150px"
                :editable="false"
                min="1"
                :max="product.stock"
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
                :disabled="!isProductInStock"
                >Add to cart</b-button
              >
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="container box">
      <b-tabs>
        <b-tab-item label="Description">
          <p class="level-left">
            {{ product.description }}
          </p>
        </b-tab-item>

        <b-tab-item label="Reviews"> </b-tab-item>

        <b-tab-item label="Specifications"> </b-tab-item>
      </b-tabs>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import Product from "@/models/product";

@Component
export default class ProductPage extends Vue {
  private loadingStatus: boolean = false;
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

  private isProductInStock(): boolean {
    return this.product.stock > 0;
  }

  @Prop({ type: Product }) readonly product!: Product;
}
</script>
