<template>
  <div>
    <div v-if="reviews && reviews.length > 0">
      <div v-for="(review, i) in reviews" v-bind:key="review.productSKU">
        <ReviewCard :review="review" />
        <hr v-if="i < reviews.length - 1" />
      </div>
    </div>
    <div v-else>
      No reviews for this product
    </div>
    <hr />
    <b-pagination
      v-if="totalReviewsForProduct > 0"
      :total="totalReviewsForProduct"
      :current.sync="currentPage"
      :range-before="rangeAfterAndBefore"
      :range-after="rangeAfterAndBefore"
      :per-page="perPage"
      order="is-centered"
      size="is-small"
      icon-prev="chevron-left"
      icon-next="chevron-right"
      aria-next-label="Next page"
      aria-previous-label="Previous page"
      aria-page-label="Page"
      aria-current-label="Current page"
    />

    <hr/>
    <ReviewForm :productSku="sku" :currentReview="currentReview"/>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from "vue-property-decorator";
import Review from "@/models/review";
import ReviewsAPI from "@/modules/graphql/reviews/reviewsgql";
import NotificationUtil from "@/utils/notification/notificationutil";
import ReviewCard from "@/components/ReviewCard.vue";
import ReviewForm from "@/components/ReviewForm.vue";
@Component({
  components: { ReviewCard, ReviewForm }
})
export default class ReviewComponent extends Vue {
  @Prop()
  private readonly sku!: string;
  private _reviewsAPI!: ReviewsAPI;

  public reviews: Array<Review> = new Array<Review>();

  public totalReviewsForProduct!: number;
  public currentPage: number = 1;
  public currentReview: Review | undefined;

  private _perPage!: number;
  private _rangeAfterAndBefore!: number;

  constructor() {
    super();
  }

  created() {
    this.setConstants();
    this.fetchReviewCount().then(() => {
      this.fetchReviewsForProduct();
    });
    this.fetchCurrentReviewForUser();
  }

  private setConstants() {
    this.totalReviewsForProduct = 0;
    this._perPage = 8;
    this._rangeAfterAndBefore = 5;
    this.currentReview = undefined;
  }

  private get API(): ReviewsAPI {
    if (this._reviewsAPI === undefined) {
      this._reviewsAPI = new ReviewsAPI();
    }

    return this._reviewsAPI;
  }

  private async fetchReviewCount() {
    await this.API.fetchTotalReviewCountForProductWithSku(this.sku)
      .then(count => {
        this.totalReviewsForProduct = count;
      })
      .catch(err => {
        this.displayStandardError();
      });
  }

  private fetchCurrentReviewForUser() {
    this.API.fetchCurrentReviewForUser(this.sku)
      .then(review => {
        if (review !== undefined) {
          console.log(review);
          this.currentReview = review;
        }
      })
      .catch(err => {
        new NotificationUtil().displayError(
            "Unable to current review for product :("
        );
      });
  }

  @Watch("currentPage")
  private fetchReviewsForProduct() {
    this.API.fetchReviewsForProductWithSku(
      this.sku,
      this.perPage,
      this.calcOffset()
    )
      .then(reviews => {
        this.reviews = reviews;
      })
      .catch(err => {
        this.displayStandardError();
      });
  }

  private calcOffset(): number {
    return this.currentPage > 1 ? (this.currentPage - 1) * this.perPage : 0;
  }

  public get perPage(): number {
    return this._perPage;
  }

  public get rangeAfterAndBefore(): number {
    return this._rangeAfterAndBefore;
  }

  private displayStandardError(): void {
    new NotificationUtil().displayError(
      "Unable to fetch reviews for product :("
    );
  }

}
</script>
