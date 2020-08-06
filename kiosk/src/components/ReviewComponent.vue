<template>
  <div></div>
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

@Component
export default class RegisterForm extends Vue {
  private _ReviewsAPI!: ReviewsAPI;
  private _reviews!: Array<Review>;

  constructor() {
    super();
  }

  created() {
    this.API.fetchReviewsForProductWithSku(this.sku)
      .then(result => {
        this._reviews = result;
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
