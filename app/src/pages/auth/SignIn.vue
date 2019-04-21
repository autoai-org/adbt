<template>
  <section id="home" class="signInPage">
    <BaseCard>
      <VCardText class="pa-5">
        <VImg :src="require('@/assets/logo.png')" class="mb-3 mx-auto" width="200"/>
        <VWindow v-model="model" lazy>
          <VWindowItem>
            <Email @next="model++"/>
          </VWindowItem>

          <VWindowItem>
            <Password @prev="model--"/>
          </VWindowItem>
        </VWindow>
      </VCardText>
    </BaseCard>
  </section>
</template>

<script>
export default {
  components: {
    BaseCard: () => import("@/components/auth/BaseCard"),
    Email: () => import("@/components/auth/Email"),
    Password: () => import("@/components/auth/Password")
  },
  data: () => ({
    model: 0
  })
};
</script>

<style scoped>
.signInPage {
}
</style>
