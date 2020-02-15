<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png" />
    <HelloWorld msg="Welcome to Your Vue.js App" />
  </div>
</template>

<script lang="ts">
// @ is an alias to /src
import Component, {mixins} from "vue-class-component";
import { AuthMixin } from "@/mixins/authmixin";
import HelloWorld from "@/components/HelloWorld.vue";

@Component({
  name: 'home',
  components: {
    HelloWorld
  }
})
export default class Home extends mixins(AuthMixin) {
  constructor() {
    super();
  }
};
</script>
