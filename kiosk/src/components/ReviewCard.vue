<template>
  <div>
    <div class="columns is-gapless">
      <b-field class="column is-1 is-6-mobile">
        <b-rate
          v-model="review.rating"
          :max="5"
          :rtl="false"
          :disabled="true"
        />
      </b-field>
      <strong class="column is-2" style="font-size: 14px">{{ review.userDisplayName }}</strong>
      <small class="column is-1" style="font-size: 14px">{{ formattedCreateDate }}</small>
    </div>
    <p class="level-left" style="font-size: 14px">
      {{ review.textReview }}
    </p>
  </div>
</template>

<style scoped>

</style>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
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
