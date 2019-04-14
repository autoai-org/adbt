<template>
  <v-app id="inspire">
    <v-toolbar color="#006dad" dark fixed app clipped-right>
      <v-toolbar-side-icon @click.native="drawer = true"></v-toolbar-side-icon>
      <v-toolbar-title>Database Backup Toolkit</v-toolbar-title>
      <v-spacer></v-spacer>
    </v-toolbar>
    <drawer :triggerDrawer="drawer"></drawer>
    <v-content>
      <v-container fluid fill-height>
        <v-layout justify-center align-center>
          <v-flex shrink>
            <v-tooltip right>
              <template v-slot:activator="{ on }">
                <v-btn :href="source" icon large target="_blank" v-on="on">
                  <v-icon large>code</v-icon>
                </v-btn>
              </template>
              <span>Source</span>
            </v-tooltip>
            <v-tooltip right>
              <template v-slot:activator="{ on }">
                <v-btn
                  icon
                  large
                  href="https://codepen.io/johnjleider/pen/KQrPKJ"
                  target="_blank"
                  v-on="on"
                >
                  <v-icon large>mdi-codepen</v-icon>
                </v-btn>
              </template>
              <span>Codepen</span>
            </v-tooltip>
          </v-flex>
        </v-layout>
      </v-container>
    </v-content>
    <v-footer color="#006dad" class="white--text" app>
      <span class="bottom_footer">Made with &hearts; by AutoAI</span>
    </v-footer>
  </v-app>
</template>

<script>
import Drawer from '@/components/Drawer'
export default {
  components: {
    Drawer
  },
  data: () => ({
    drawer: null,
  }),
  props: {
    source: String
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
