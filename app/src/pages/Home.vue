<template>
  <v-container fill-height fluid grid-list-xl>
    <v-layout wrap>
      <v-flex md12 lg6>
        <job-list></job-list>
      </v-flex>
      <v-flex md12 lg6>
        <calendar></calendar>
      </v-flex>
      <!--
      <v-flex md12 sm12 lg4>
        <material-chart-card
          :data="dailySalesChart.data"
          :options="dailySalesChart.options"
          color="info"
          type="Line"
        >
          <h4 class="title font-weight-light">Daily Backups</h4>
          <p class="category d-inline-flex font-weight-light">
            <v-icon color="green" small>mdi-arrow-up</v-icon>
            <span class="green--text">55%</span>&nbsp;
            Success Backups
          </p>

          <template slot="actions">
            <v-icon class="mr-2" small>mdi-clock-outline</v-icon>
            <span class="caption grey--text font-weight-light">updated 4 minutes ago</span>
          </template>
        </material-chart-card>
      </v-flex>
      -->
    </v-layout>
  </v-container>
</template>

<script>
import JobList from "@/components/JobList";
import Calendar from '@/components/Calendar';
export default {
  components: {
    JobList,
    Calendar
  },
  data() {
    return {
    };
  },
  methods: {
    complete(index) {
      this.list[index] = !this.list[index];
    }
  }
};
</script>