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
          <template v-slot:day="{ date }">
            <template v-if="logsMap[date]">
              <v-menu full-width offset-x>
                <template v-slot:activator="{ on }">
                  <v-btn
                    color="green darken-2"
                    small
                    v-if="parseRatio(logsMap[date]) >= 50"
                    v-ripple
                    class="my-event"
                    v-on="on"
                    @click="getLogsJob(logsMap[date])"
                  >
                    <p
                      style="font-size:1.5vw;margin-bottom:0"
                    >{{ parseRatio(logsMap[date]) + '% Passed' }}</p>
                  </v-btn>
                  <v-btn
                    color="red darken-2"
                    small
                    v-if="parseRatio(logsMap[date]) < 50"
                    v-ripple
                    class="my-event"
                    v-on="on"
                    @click="getLogsJob(logsMap[date])"
                  >
                    <p
                      style="font-size:1.5vw;margin-bottom:0"
                    >{{ parseRatio(logsMap[date]) + '% Passed' }}</p>
                  </v-btn>
                </template>
                <v-card color="grey lighten-4" min-width="350px" flat>
                  <v-toolbar color="primary" dark>
                    <v-toolbar-title v-html="parseRatio(logsMap[date]) + '% Passed'"></v-toolbar-title>
                  </v-toolbar>
                    <v-list three-line v-if="detailLoaded">
                      <v-list-tile v-for="(item, index) in logsMap[date]" :key="index">
                        <v-list-tile-content>
                          <v-list-tile-title>{{ logsDetail[item.identifier].name }}</v-list-tile-title>
                        <v-list-tile-sub-title v-if="!item.success">
                          <div>
                            <v-icon color="red darken-2">error</v-icon>
                            Failed at {{item.time.slice(11,20)}} {{ item.log }}.
                          </div>
                        </v-list-tile-sub-title>
                        <v-list-tile-sub-title v-if="item.success">
                          <div>
                            <v-icon color="green darken-2">check</v-icon>
                            Succeeed at {{item.time.slice(11,20)}} with no errors.
                          </div>
                        </v-list-tile-sub-title>
                        </v-list-tile-content>
                      </v-list-tile>
                    </v-list>
                  <v-card-actions>
                    <v-btn flat color="secondary">Close</v-btn>
                  </v-card-actions>
                </v-card>
              </v-menu>
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
      detailLoaded: false,
      logsLoaded: false,
      type: "month",
      logs: [],
      logsMap: {},
      logsDetail: {}
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
        let map = {};
        self.logs.forEach(e =>
          (map[e.formattedTime] = map[e.formattedTime] || []).push(e)
        );
        self.logsLoaded = true;
        self.logsMap = map;
      });
    },
    parseRatio(logs) {
      let success = 0;
      let failed = 0;
      for (var i in logs) {
        if (logs[i].success) {
          success = success + 1;
        } else {
          failed = failed + 1;
        }
      }
      return (success / (failed+success)).toFixed(2) * 100;
    },
    getLogsJob(logs) {
      let self = this;
      let seen = new Set();
      logs.map(item => {
        seen.add(item.identifier);
      });
      let logsList = [...seen];
      Promise.all(
        logsList.map(async each => {
          return await adbtService.getJobDetail(each).then(function(res) {
            self.logsDetail[each] = res.data;
          });
        })
      ).then(function() {
        self.detailLoaded = true;
      });
    }
  },
  created() {
    this.fetchData();
  }
};
</script>

<style>
</style>
