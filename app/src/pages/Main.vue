<template>
  <v-app id="inspire">
    <v-toolbar color="#006dad" dark fixed app clipped-right extended>
      <v-toolbar-side-icon @click.native="triggerNavbar()"></v-toolbar-side-icon>
      <v-toolbar-title>Database Backup Toolkit</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn
        color="indigo lighten-2"
        dark
        medium
        absolute
        bottom
        right
        fab
        @click="triggerCreateDialog()"
      >
        <v-icon>add</v-icon>
      </v-btn>
    </v-toolbar>
    <drawer :triggerDrawer="drawer" @onClose="onNavbarClose()"></drawer>
    <v-content>
      <v-container fluid fill-height>
        <v-layout justify-center align-center>
          <v-flex shrink>
            <router-view :key="needRefresh"></router-view>
          </v-flex>
        </v-layout>
      </v-container>
    </v-content>
    <v-footer color="#006dad" class="white--text" app>
      <span class="bottom_footer">Made with &hearts; by AutoAI</span>
    </v-footer>
    <v-dialog v-model="addDialog" fullscreen hide-overlay transition="dialog-bottom-transition">
        <new @onDialogClose="afterCreated()" ></new>
    </v-dialog>
  </v-app>
</template>

<script>
import Drawer from "@/components/Drawer";
import New from "@/components/New";
export default {
  components: {
    Drawer,
    New
  },
  data: () => ({
    addDialog: false,
    drawer: null,
    t_drawer: false,
    needRefresh:false,
  }),
  props: {
    source: String
  },
  methods: {
    triggerNavbar() {
      this.drawer = true;
    },
    onNavbarClose() {
      this.drawer = false;
    },
    triggerCreateDialog () {
      this.addDialog = true;
    },
    afterCreated () {
      this.addDialog = false;
      this.needRefresh = true;
    }
  }
};
</script>

<style>
body {
  overflow-x: hidden;
}
.bottom_footer {
  margin-left: 10px;
}
</style>
