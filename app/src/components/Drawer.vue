<template>
  <v-navigation-drawer
    id="app-drawer"
    v-model="drawer"
    app
    dark
    floating
    persistent
    temporary
    mobile-break-point="991"
    width="260"
  >
    <v-img :src="'https://i.loli.net/2019/04/14/5cb2c081db35e.jpg'" height="100%">
      <v-layout class="fill-height" column>
        <v-list-tile avatar>
          <v-list-tile-avatar color="white">
            <v-img height="34" contain/>
          </v-list-tile-avatar>
          <v-list-tile-title class="title">ADBT</v-list-tile-title>
        </v-list-tile>
        <v-divider/>
        <v-list-tile v-if="responsive">
          <v-text-field class="purple-input search-input" label="Search..." color="purple"/>
        </v-list-tile>
        <v-list-tile
          v-for="(link, i) in links"
          :key="i"
          :to="link.to"
          avatar
          class="v-list-item"
        >
          <v-list-tile-action>
            <v-icon>{{ link.icon }}</v-icon>
          </v-list-tile-action>
          <v-list-tile-title v-text="link.text"/>
        </v-list-tile>
        <v-list-tile
          v-for="(link, id) in externalLinks"
          :key="'el-' + id"
          @click="navto(link.to)"
          avatar
          class="v-list-item"
        >
          <v-list-tile-action>
            <v-icon>info</v-icon>
          </v-list-tile-action>
          <v-list-tile-title v-text="link.text"/>
        </v-list-tile>
        <v-list-tile
          class="v-list-item v-list__tile--buy"
          @click="navto('https://hub.autoai.org/account/signin')"
        >
          <v-list-tile-action>
            <v-icon color="black">supervisor_account</v-icon>
          </v-list-tile-action>
          <v-list-tile-title class="font-weight-light upgrade-text">AutoAI Membership</v-list-tile-title>
        </v-list-tile>
      </v-layout>
    </v-img>
  </v-navigation-drawer>
</template>

<script>
import open from 'open'
export default {
  props: ['triggerDrawer'],
  data: () => ({
    logo: "../assets/logo.png",
    links: [
      {
        to: "/dashboard",
        icon: "dashboard",
        text: "Dashboard"
      },
    ],
    externalLinks: [
      {
        to: "https://github.com/unarxiv/CVPM/wiki/Privacy-Policy",
        text: "Privacy Policy"
      },
      {
        to: "https://github.com/unarxiv/CVPM/wiki/Terms-of-Use",
        text: "Terms of Use"
      },
    ],
    responsive: false,
    drawer: false,
  }),
  watch: {
    triggerDrawer(val) {
      this.drawer = val
    },
    drawer(val) {
      if (val === false) {
        this.$emit("onClose", true)
      }
    }
  },
  methods: {
    navto (link) {
      open(link)
    }
  }
};
</script>

<style lang="scss">
  #app-drawer {
    .v-list__tile {
      border-radius: 4px;
      &--buy {
        margin-top: auto;
        margin-bottom: 17px;
      }
    }
    .v-image__image--contain {
      top: 9px;
      height: 60%;
    }
    .search-input {
      margin-bottom: 30px !important;
      padding-left: 15px;
      padding-right: 15px;
    }
    .upgrade-text {
      color: black;
    }
  }
</style>
