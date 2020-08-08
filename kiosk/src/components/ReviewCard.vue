<template>
  <div>
    <b-field>
      <b-rate
        v-model="review.rating"
        :max="5"
        :rtl="false"
        :disabled="true"
      />
      <p style="padding-left: 20px">
        {{review.userDisplayName}}
      </p>
      <p style="margin-left: 200px">
        {{ formattedCreateDate }}
      </p>
    </b-field>
    <p style="float: left">
      {{review.textReview}}
    </p>
  </div>
</template>

<script lang="ts">
import {
  Component,
  Prop,
  Vue,
} from "vue-property-decorator";
import Review from "@/models/review";
import DateUtil from "@/utils/date/dateutil";

@Component
export default class ReviewCard extends Vue {
  @Prop()
  private readonly review!: Review;

  constructor() {
    super();
  }

  private get formattedCreateDate() {
    return DateUtil.formatDateToLongDateDisplay(this.review.createdAt);
  }
}
</script>
