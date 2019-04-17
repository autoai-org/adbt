<template>
  <material-card class="card-tabs" color="#006dad">
    <v-flex slot="header">
      <v-tabs color="transparent" slider-color="white">
        <span class="subheading font-weight-light mr-3" style="align-self: center">Recent Backups</span>
      </v-tabs>
    </v-flex>
    <v-tabs-items v-if="logsLoaded">
      <v-sheet height="500">
        <v-calendar ref="calendar" :type="type" color="primary">
          aa
          <template v-slot:day="{ date }">
            <template v-for="log in logsMap[date]">
              <v-chip color="red" text-color="white" :key="log.identifier">{{log.identifier}}</v-chip>
              <v-chip color="green" text-color="white" :key="log.identifier">Colored Chip</v-chip>
            </template>
          </template>
        </v-calendar>
      </v-sheet>
    </v-tabs-items>
  </material-card>
</template>

<script>
import MaterialCard from "@/components/Card";
import { adbtService } from "@/service/adbt";

export default {
  components: {
    MaterialCard
  },
  data() {
    return {
      now:"2019-04-18",
      logsLoaded: false,
      type: "month",
      logs: []
    };
  },
  methods: {
    fetchData() {
      let self = this;
      adbtService.getLogs("all").then(function(res) {
        self.logs = res.data.logs.map(function(each) {
          each.formattedTime = each.time.slice(0, 10);
          return each;
        });
        self.logsLoaded = true
      });
    }
  },
  created() {
    this.fetchData();
  },
  computed: {
    logsMap() {
      const map = {};
      this.logs.forEach(e => (map[e.formattedTime] = map[e.formattedTime] || []).push(e));
      return map;
    },
  }
};
</script>

<style>
</style>
