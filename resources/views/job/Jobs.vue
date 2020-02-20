<template>
  <div>
    <section class="hero is-primary">
      <div class="hero-body">
        <div class="container">
          <h1 class="title">
            Jobs
          </h1>
          <h2 class="subtitle">
            My Jobs are here
          </h2>
        </div>
      </div>
    </section>
    <section class="section">
      <div class="container">
        <div class="box">
          <div class="columns">
            <div class="column is-4">
              <div class="action-boutons">
                <div class="field has-addons">
                  <p class="control">
                    <b-button
                      icon-right="plus"></b-button>
                  </p>
                  <p class="control">
                    <b-button
                      icon-right="trash"></b-button>
                  </p>
                </div>
              </div>
              <div class="job-list">
                <job-preview
                  v-for="job in jobs"
                  v-bind:key="job.id"
                  :title="job.title"
                  @click.native="selectJob(job)"
                  :class="{
                    'active': isActiveJob(job)
                  }"
                ></job-preview>
              </div>
            </div>
            <div
              :class="{
                'column is-8 hero is-fullheight': true,
                'is-hidden': !hasActiveJob()
              }">
              <div class="box">
                <job-form
                  v-if="selectedJob"
                  v-bind.sync="selectedJob"></job-form>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
  import JobPreview from '../../views/job/JobPreview.vue'
  import JobForm from '../../views/job/JobForm.vue'
  let ghytConfApi = process.env.GHYT_CONF_API;

  export default {
    name: 'jobs',
    components: {
      JobPreview,
      JobForm
    },
    data: function() {
      return {
        selectedJob: null,
        jobs: [
          {
            id:1,
            title: "my job 1",
            conditions: [
              {
                "name": "equal",
                "args": {
                  "variableName": "event.pull_request.state",
                  "value": "open"
                },
                "jobId": 1,
                "id": 1
              },
              {
                "name": "regex",
                "args": {
                  "variableName": "event.pull_request.title",
                  "value": "connect-[^-][0-9]*",
                  "persistName": "yt_id"
                },
                "jobId": 1,
                "id": 2
              }
            ],
            actions: [
              {
                "to": "youtrack",
                "name": "addTag",
                "args": {
                  "youtrackId": "yt_id",
                  "tagName": "nok"
                },
                "jobId": 1,
                "id": 1
              },
              {
                "to": "youtrack",
                "name": "removeTag",
                "args": {
                  "youtrackId": "yt_id",
                  "tagName": "nok"
                },
                "jobId": 1,
                "id": 2
              }
            ]
          },
          {
            title: "my job 2",
            id:2,
            conditions: [],
            actions: []
          },
          {
            title: "my job 3",
            id:3,
            conditions: [],
            actions: []
          },
          {
            title: "my job 4",
            id:4,
            conditions: [],
            actions: []
          },
          {
            title: "my job 5",
            id:5,
            conditions: [],
            actions: []
          },
        ],
      };
    },
    mounted() {
      //this.axios.get(ghytConfApi + '/jobs').then(response => this.jobs = response.data);
    },
    computed: {
    },
    methods: {
      hasActiveJob(){
        return this.selectedJob !== null
      },
      isActiveJob(job){
        return this.selectedJob !== null && job.id === this.selectedJob.id
      },
      addNewJob() {
        this.$router.push('/job');
      },
      selectJob(job) {
        this.selectedJob = job
      }
    },
  };
</script>

<style>

</style>
