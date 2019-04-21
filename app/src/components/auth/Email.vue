<template>
  <div>
    <BaseHeading class="text-xs-center" text="Sign in" sub="to continue to Auto AI"/>

    <VTextField v-model="internalValue" autofocus label="Email/Username" name="email"/>

    <!--
    <BaseText>Forgot email?</BaseText>
    -->

    <div class="my-5">
      <div class="subheading">Want to learn more about Auto AI?</div>
      <BaseText href="https://autoai.org">Learn more</BaseText>
    </div>

    <VLayout align-center>
      <VFlex xs8>
        <BaseText to="/account/signup">Create account</BaseText>
      </VFlex>
      <VFlex text-xs-right>
        <v-btn flat color='indigo' :disabled="!internalValue" class="ma-0" @click="$emit('next')">Next</v-btn>
      </VFlex>
    </VLayout>
  </div>
</template>

<script>
import BaseText from "@/components/auth/BaseText";
import BaseHeading from "@/components/auth/BaseHeading";

import { mapMutations } from "vuex";

export default {
  name: "Email",
  components: {
    BaseText,
    BaseHeading,
  },
  computed: {
    internalValue: {
      get() {
        return this.$store.state.email;
      },
      set(val) {
        this.setEmail(val);
      }
    }
  },

  methods: {
    ...mapMutations(["setEmail"])
  }
};
</script>
