<template>
  <v-card>
    <v-toolbar dark color="primary">
      <v-btn icon dark @click="closeDialog()">
        <v-icon>close</v-icon>
      </v-btn>
      <v-toolbar-title>Settings</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-toolbar-items>
        <v-btn dark flat @click="saveJob()">
          <v-icon>done</v-icon>Save
        </v-btn>
      </v-toolbar-items>
    </v-toolbar>
    <v-form ref="form" v-model="valid" lazy-validation class="newJobForm">
      <v-text-field v-model="name" label="Name" required></v-text-field>
      <v-select
        v-model="selected_database"
        :items="databases"
        :rules="[v => !!v || 'Database is required']"
        label="Database"
        required
      ></v-select>
      <v-text-field v-model="uri" label="URI" required></v-text-field>

      <v-select
        v-model="selected_period"
        :items="periods"
        :rules="[v => !!v || 'Period is required']"
        label="Period"
        required
      ></v-select>
      <v-menu
        ref="menu"
        v-model="timepicker"
        :close-on-content-click="false"
        :nudge-right="40"
        :return-value.sync="time"
        lazy
        transition="scale-transition"
        offset-y
        full-width
        max-width="290px"
        min-width="290px"
      >
        <template v-slot:activator="{ on }">
          <v-text-field
            v-model="time"
            label="Backup Time"
            prepend-icon="access_time"
            readonly
            v-on="on"
          ></v-text-field>
        </template>
        <v-time-picker
          v-if="timepicker"
          v-model="time"
          full-width
          @click:minute="$refs.menu.save(time)"
        ></v-time-picker>
      </v-menu>

      <v-btn :disabled="!valid" color="success" @click="validate">Validate</v-btn>

      <v-btn color="error" @click="reset">Reset Form</v-btn>

      <v-btn color="warning" @click="resetValidation">Reset Validation</v-btn>
      <v-snackbar v-model="snackbar" bottom>
        {{ snacktext }}
        <v-btn color="pink" flat @click="snackbar = false">Close</v-btn>
      </v-snackbar>
    </v-form>
  </v-card>
</template>

<script>
import { adbtService } from "@/service/adbt";
export default {
  data: () => ({
    valid: true,
    name: "",
    uri: "",
    select: null,
    periods: ["Monthly", "Weekly", "Daily"],
    databases: ["MongoDB"],
    selected_period: null,
    snackbar: null,
    selected_database: null,
    time: null,
    timepicker: false,
    snacktext: null
  }),
  methods: {
    validate() {
      if (this.$refs.form.validate()) {
        let self = this;
        this.snackbar = true;
        adbtService
          .testJob(this.uri, this.selected_database, this.name)
          .then(function(res) {
            self.snackbar = true;
            self.snacktext = res.data.info;
          })
          .catch(function(err) {
            self.snackbar = true;
            self.snacktext = err.response.data.info;
          });
      }
    },
    reset() {
      this.$refs.form.reset();
    },
    resetValidation() {
      this.$refs.form.resetValidation();
    },
    saveJob() {
      let self = this;
      adbtService
        .addJobs(
          this.uri,
          this.selected_period,
          this.time,
          this.selected_database,
          this.name
        )
        .then(function(res) {
          self.snackbar = 'Successfully Added' + self.name;
          self.snacktext = res.data.info;
          self.closeDialog()
        })
        .catch(function(err) {
          self.snackbar = true;
          self.snacktext = err.response.data.info;
        });
    },
    closeDialog() {
      this.$emit("onDialogClose");
    }
  }
};
</script>

<style scoped>
.newJobForm {
  margin-left: 10px;
  margin-right: 10px;
  margin-top: 10px;
}
</style>
