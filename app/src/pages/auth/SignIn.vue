<template>
  <section id="home" class="signInPage">
    <BaseCard>
      <VCardText class="pa-5">
        <VImg
          :src="require('@/assets/logo_bg.png')"
          class="mb-3 mx-auto"
          width="200"
        />
        <VWindow
          v-model="model"
          lazy
        >
          <VWindowItem>
            <Email @next="model++" />
          </VWindowItem>

          <VWindowItem>
            <Password @prev="model--" />
          </VWindowItem>
        </VWindow>
      </VCardText>
    </BaseCard>
  </section>
</template>

<script>
  export default {
    components: {
      Email: () => import('@/components/Auth/Email'),
      Password: () => import('@/components/Auth/Password')
    },
    data: () => ({
      model: 0
    })
  }
</script>

<style scoped>
.signInPage {
  width: 60%;
  height: 60%;
  margin: auto;
  margin-top: 5%;
}
</style>
