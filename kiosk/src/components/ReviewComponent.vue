<template>
  <div>
    <div v-if="reviews && reviews.length > 0">
      <div v-for="review in reviews" v-bind:key="review.productSKU">
        <ReviewCard :review="review" />
        <hr>
      </div>
    </div>
    <div v-else>
      No reviews for this product
    </div>
  </div>
</template>

<script lang="ts">
import {
  Component,
  Prop,
  Vue,
} from "vue-property-decorator";
import Review from "@/models/review";
import ReviewsAPI from "@/modules/graphql/reviews/reviewsgql";
import NotificationUtil from "@/utils/notification/notificationutil";
import ReviewCard from "@/components/ReviewCard.vue";
@Component({
  components: {ReviewCard}
})
export default class ReviewComponent extends Vue {
  private _ReviewsAPI!: ReviewsAPI;
  public reviews: Array<Review> = new Array<Review>();

  constructor() {
    super();
  }

  created() {
    this.API.fetchReviewsForProductWithSku(this.sku)
      .then(reviews => {
        this.reviews = reviews;
      })
      .catch(err => {
        new NotificationUtil().displayError(
          "Unable to fetch reviews for product :("
        );
      });
  }

  private get API(): ReviewsAPI {
    if (this._ReviewsAPI === undefined) {
      this._ReviewsAPI = new ReviewsAPI();
    }

    return this._ReviewsAPI;
  }

  @Prop()
  private readonly sku!: string;
}
</script>
