<template>
  <div>
    <strong> {{ componentTitle }}  </strong>
      <b-field >
        <b-rate
            v-model="rating"
            :max="5"
            :rtl="false"
            class="center-header"
        />
      </b-field>
      <div class="field">
        <b-checkbox v-model="anonymous">
          Anonymous
        </b-checkbox>
      </div>
    <div class="columns center-header">
      <b-field class="column is-half" label="">
        <b-input type="textarea" maxlength="250" v-model="textReview" />
      </b-field>
    </div>
    <b-button type="is-primary" @click="alert('')">{{ displayTextForButton }}</b-button>
  </div>
</template>

<style scoped>
.center-header {
  justify-content: center
}
</style>

<script lang="ts">
import {Component, Prop, Vue, Watch} from "vue-property-decorator";
import Review from "@/models/review";

@Component
export default class ReviewForm extends Vue {
  public rating: number = 0;
  public textReview: string = "";
  public anonymous: boolean = false;

  @Prop()
  private readonly productSku!: string;
  @Prop()
  private readonly currentReview!: Review | undefined;

  constructor() {
    super();
  }

  @Watch('currentReview')
  private setReviewToComponentField(review: Review) : void {
    this.rating = review.rating;
    this.textReview = review.textReview;
    this.anonymous = review.anonymous;
  }

  public get componentTitle(): string {
    if (this.currentReview === undefined)
      return "Submit your review!";
    else
      return "Your review of this product";
  }

  public get displayTextForButton(): string {
    if (this.currentReview === undefined)
      return "Submit your review";
    else
      return "Update your review";
  }

}
</script>
