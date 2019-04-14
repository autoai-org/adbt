<template>
  <v-navigation-drawer
    id="app-drawer"
    v-model="triggerDrawer"
    app
    dark
    floating
    persistent
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
          disabled
          class="v-list-item v-list__tile--buy"
          to="/"
        >
          <v-list-tile-action>
            <v-icon>mdi-package-up</v-icon>
          </v-list-tile-action>
          <v-list-tile-title class="font-weight-light upgrade-text">Upgrade To PRO</v-list-tile-title>
        </v-list-tile>
      </v-layout>
    </v-img>
  </v-navigation-drawer>
</template>

<script>
export default {
  props: ['triggerDrawer'],
  data: () => ({
    logo: "../assets/logo.png",
    links: [
      {
        to: "/dashboard",
        icon: "mdi-view-dashboard",
        text: "Dashboard"
      },
    ],
    responsive: false,
    drawer:false,
  }),
  watch: {
    triggerDrawer(val) {
      // eslint-disable-next-line
      console.log(val)
      this.drawer = this.triggerDrawer
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