<template>
  <div>
    <div class="text-xs-center">
      <VChip class="mb-3" outline @click="$emit('prev')">
        <VIcon color="primary" left>mdi-account-circle</VIcon>
        <span class="body-2" v-text="email"/>
        <VIcon color="black" right>mdi-chevron-down</VIcon>
      </VChip>
    </div>

    <v-text-field
      v-model="internalValue"
      :append-icon="show ? 'mdi-eye' : 'mdi-eye-off'"
      :type="show ? 'text' : 'password'"
      class="mb-3"
      label="Password"
      name="password"
      @click:append="show = !show"
    />

    <div class="text-xs-left">
      <v-layout align-center justify-space-between>
        <BaseText>Forgot password?</BaseText>

        <VBtn
          :disabled="!internalValue"
          :loading="isLoading"
          class="text-capitalize ma-0 white--text"
          color="indigo darken-1"
          depressed
          @click="submit"
        >Next</VBtn>
      </v-layout>
    </div>
  </div>
</template>

<script>
// Utilities
import BaseText from "@/components/auth/BaseText";

import { mapActions, mapMutations, mapState } from "vuex";

export default {
  data: () => ({
    show: false
  }),
  components: {
    BaseText
  },
  computed: {
    ...mapState(["email", "password", "isLoading"]),
    internalValue: {
      get() {
        return this.$store.state.password;
      },
      set(val) {
        this.setPassword(val);
      }
    }
  },

  methods: {
    ...mapActions("cognito", ["signInUser"]),
    ...mapMutations(["setEmail", "setPassword", "setIsLoading", "setSnackbar"]),
    submit() {
      this.hasError = false;
      this.setIsLoading(true);
      this.signInUser({
        username: this.email,
        password: this.password
      })
        .then(() => {
          this.setSnackbar({
            type: "success",
            msg: `Successfully signed in user ${this.email}`
          });

          this.setEmail("");
          this.$router.replace("/post");
        })
        .catch(res => {
          this.setSnackbar({
            type: "error",
            msg: res
          });
        })
        .finally(() => {
          this.setIsLoading(false);
          this.setPassword("");
        });
    }
  }
};
</script>

<style lang="scss">
.v-chip__content {
  cursor: pointer !important;
}
</style>
