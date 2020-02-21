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

              <div class="columns">
                <div class="column is-12">
                  <div class="action-boutons">
                    <div class="field has-addons">
                      <p class="control">
                        <b-button
                          icon-right="plus"
                          @click="addJob()"
                        ></b-button>
                      </p>
                      <p class="control">
                        <b-button
                          icon-right="trash"
                          @click="deleteJob()"
                          :disabled="!hasActiveJob()"
                        ></b-button>
                      </p>
                    </div>
                  </div>
                </div>
              </div>
              <div class="columns">
                <div class="column is-12">
                  <div class="job-list">
                    <job-preview
                      v-for="job in jobs"
                      v-bind:key="job.id"
                      :title="job.name"
                      @click.native="selectJob(job)"
                      :class="{
                        'active': isActiveJob(job),
                        'is-loading': isLoadingJobs
                      }"
                    ></job-preview>
                    <b-loading :is-full-page="false" :active.sync="isLoadingJobs" :can-cancel="true"></b-loading>
                  </div>
                </div>
              </div>
            </div>
            <div
              :class="{
                'column is-8 hero is-fullheight': true,
                'is-hidden': !hasActiveJob()
              }">
              <div class="box">
                <job-form
                  v-if="hasActiveJob()"
                  v-bind.sync="selectedJob"
                ></job-form>
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
        jobs: [],
        isLoadingJobs: true,
        islod: false,
      };
    },
    mounted() {
      this.axios.get(ghytConfApi + '/jobs')
        .then(response => {
          this.isLoadingJobs = false
          this.jobs = response.data
        });
    },
    methods: {
      hasActiveJob(){
        return this.selectedJob !== null
      },
      isActiveJob(job){
        return this.selectedJob !== null && job.id === this.selectedJob.id
      },
      addJob() {
        const id = Math.floor(Math.random() * 900)
        this.jobs.push({'name': 'new job '+id, conditions: [], actions: [], id: id})
      },
      deleteJob() {
        const index = this.jobs.findIndex(element => this.selectedJob.id === element.id)
        this.jobs.splice(index, 1)
        this.selectedJob = this.jobs.length>=1?this.jobs[index]?this.jobs[index]:this.jobs[index-1]:null
      },
      selectJob(job) {
        this.selectedJob = job
      },
    },
  };
</script>

<style>

</style>
