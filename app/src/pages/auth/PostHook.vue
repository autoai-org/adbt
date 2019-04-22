<template>
  <section id="home" class="signInPage">
    <BaseCard>
      <VCardText class="pa-5">
        <VImg :src="require('@/assets/logo.png')" class="mb-3 mx-auto" width="200"/>
            <v-flex row wrap align-center justify-center style="text-align:center">
            <p>Successfully logon as <B>{{username}}</B></p>
            <p>If you run ADBT on a remote server without GUI, you can use the CLI command
                <br/>
                <code>adbt login {{username}} {{userId}}</code>
                <br/>
                to enable log & config upload.<a href="https://github.com/unarxiv/adbt/wiki/Remote-Logging-and-Config-Upload">Help</a>
            </p>
            <v-btn flat color='indigo' @click="backHome()">Back Home</v-btn>
            </v-flex>
      </VCardText>
    </BaseCard>
  </section>
</template>

<script>
export default {
  components: {
    BaseCard: () => import("@/components/auth/BaseCard")
  },
  data: () => ({
    model: 0,
    username: '',
    userdata: {},
    userId: '',
  }),
  methods: {
    backHome() {
      this.$router.replace("/");
    },
    loadData() {
        this.username = this.$store.state.cognito.user.username
        let key= this.$store.state.cognito.user.keyPrefix + '.' + this.username + '.userData'
        this.userdata = JSON.parse(this.$store.state.cognito.user.storage[key])
        for(var idx in this.userdata.UserAttributes) {
            if (this.userdata.UserAttributes[idx].Name === 'sub') {
                this.userId = this.userdata.UserAttributes[idx].Value
            }
        }
    }
  },
  created() {
      this.loadData()
  }
};
</script>

<style>
</style>
