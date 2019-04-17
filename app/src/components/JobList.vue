<template>
  <material-card class="card-tabs" color="green">
    <v-flex slot="header">
      <v-tabs color="transparent" slider-color="white">
        <span class="subheading font-weight-light mr-3" style="align-self: center">Scheduled Backups</span>
      </v-tabs>
    </v-flex>
    <v-tabs-items v-if="jobsLoaded">
      <v-tab-item>
        <v-list three-line>
          <v-list-tile v-for="(item, index) in jobsList" :key="index">
            <v-list-tile-avatar>
              <v-icon>icon-mongodb</v-icon>
            </v-list-tile-avatar>
            <v-list-tile-content>
              <v-list-tile-title>{{ jobsDetail[index].name }}</v-list-tile-title>
              <v-list-tile-sub-title>
                <v-tooltip top content-class="top">
                  <template v-slot:activator="{ on }">
                    <span v-on="on">
                      <B>will backup</B>
                      {{ formatTimeFromNow(item.nextRun) }} at {{ formatTimeToDay(item.nextRun) }}
                    </span>
                  </template>
                  <span>{{ formatTimeToDay(item.nextRun) }}, {{formatTimeToSec(item.nextRun)}}</span>
                </v-tooltip>
              </v-list-tile-sub-title>
            </v-list-tile-content>
            <div class="d-flex">
              <v-tooltip top content-class="top">
                <v-btn slot="activator" class="v-btn--simple" color="white" icon @click="modifyJob(jobsDetail[index])">
                  <v-icon color="primary">edit</v-icon>
                </v-btn>
                <span>Edit</span>
              </v-tooltip>
              <v-tooltip top content-class="top">
                <v-btn slot="activator" class="v-btn--simple" color="white" icon @click="runJob(item.identifier)">
                  <v-icon color="primary">sync</v-icon>
                </v-btn>
                <span>Backup Now</span>
              </v-tooltip>
              <v-tooltip top content-class="top">
                <v-btn slot="activator" class="v-btn--simple" color="white" icon @click="deleteJob(item.identifier)">
                  <v-icon color="primary">delete</v-icon>
                </v-btn>
                <span>Delete</span>
              </v-tooltip>
            </div>
          </v-list-tile>
          <v-divider/>
        </v-list>
      </v-tab-item>
    </v-tabs-items>
    <v-dialog v-model="modifyDialog" v-if="modifyDialog" fullscreen hide-overlay transition="dialog-bottom-transition">
      <modify @onDialogClose="onModifiedJob()" :item="toModifyItem" ></modify>
    </v-dialog>
  </material-card>
</template>

<script>
import MaterialCard from "@/components/Card";
import { adbtService } from "@/service/adbt";
import Modify from '@/components/Modify'
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
dayjs.extend(relativeTime);
export default {
  components: {
    MaterialCard,
    Modify
  },
  data() {
    return {
      list: {
        0: false,
        1: false,
        2: false
      },
      jobsDetail: [],
      jobsList: [],
      jobsLoaded: false,
      modifyDialog: false,
      toModifyItem: {}
    };
  },
  created() {
    this.fetchJobs();
  },
  methods: {
    fetchJobs() {
      let self = this;
      adbtService
        .getJobs()
        .then(function(res) {
          self.jobsList = res.data;
          self.fetchDetailedJobs(self.jobsList);
        })
        .catch(function(err) {
          // eslint-disable-next-line
          console.error(err);
        });
    },
    fetchDetailedJobs() {
      let self = this;
      adbtService
        .getDetailedJobs(self.jobsList)
        .then(function(res) {
          self.jobsDetail = res.map(function(res) {
            return res.data;
          });
          self.jobsLoaded = true;
        })
        .catch(function(err) {
          // eslint-disable-next-line
          console.error(err.response);
        });
    },
    modifyJob (item) {
      this.toModifyItem = item
      this.modifyDialog = true
    },
    runJob (identifier) {
      let self = this
      adbtService.runJob(identifier).then(function (res) {
        console.log(res)
      }).catch(function(err) {
        console.error(err.response)
      })
    },
    deleteJob (identifier) {
      let self = this
      adbtService.deleteJob(identifier).then(function () {
        self.fetchJobs()
      }).catch(function(err) {
          // eslint-disable-next-line
          console.error(err.response);
      })
    },
    onModifiedJob () {
      this.fetchJobs()
      this.modifyDialog = false
    },
    formatTimeFromNow(time) {
      return dayjs(time).fromNow();
    },
    formatTimeToDay(time) {
      return dayjs(time).format("DD, MMM");
    },
    formatTimeToSec(time) {
      return dayjs(time).format("HH:mm");
    }
  }
};
</script>

<style>
</style>
