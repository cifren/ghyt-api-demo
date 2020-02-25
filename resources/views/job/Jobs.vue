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
          <div class="columns" v-if="!errorMessage">
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
          <b-notification
            v-if="errorMessage"
            type="is-danger"
            aria-close-label="notification">
            {{errorMessage}}
          </b-notification>
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
        errorMessage: null
      };
    },
    mounted() {
      if(ghytConfApi === undefined){
        this.errorMessage = "Error : The API url is missing"
        return
      }

      this.axios.get(ghytConfApi + '/jobs')
        .then(response => {
          this.isLoadingJobs = false
          this.jobs = response.data
        })
        .catch(error => {
          this.isLoadingJobs = false
          this.errorMessage = error
        });
    },
    created: function () {
      this.debouncedSaving = _.debounce(this.save, 2000)
    },
    watch: {
      selectedJob: {
        deep: true,
        handler(newJob, oldJob){
          // update only
          if(newJob === null || oldJob === null){
            return
          }
          if(newJob.id !== oldJob.id){
            return
          }
          this.debouncedSaving(newJob)
        }
      },
    },
    methods: {
      save(job){
        if(job.id === null)
          return
        this.remotePatchJob(job)
      },
      // update job
      remotePatchJob(job){
        console.log(job)
        this.axios.patch(ghytConfApi + '/jobs/' + job.id, job)
          .then(() => {
            this.$buefy.toast.open(this.getToastSuccessConf('updated'))
          })
          .catch(error => {
            this.$buefy.toast.open({
              message: 'Error Happened: ' + error.message,
              duration: 10000,
              type: 'is-danger'
            })
          });
      },
      // create job
      remotePostJob(job){
        this.axios.post(ghytConfApi + '/jobs', job)
          .then((response) => {
            this.$buefy.toast.open(this.getToastSuccessConf('added'))
            job.id = response.data.id
          })
          .catch(error => {
            this.$buefy.toast.open({
              message: 'Error Happened: ' + error.message,
              duration: 10000,
              type: 'is-danger'
            })
          });
      },
      // create job
      remoteDeleteJob(job){
        return this.axios.delete(ghytConfApi + '/jobs/' + job.id)
          .then((response) => {
            this.$buefy.toast.open(this.getToastSuccessConf('deleted'))
          })
          .catch(error => {
            this.$buefy.toast.open({
              message: 'Error Happened: ' + error.message,
              duration: 10000,
              type: 'is-danger'
            })
          });
      },
      hasActiveJob(){
        return this.selectedJob !== null
      },
      isActiveJob(job){
        return this.selectedJob !== null && job.id === this.selectedJob.id
      },
      addJob(){
        this.jobs.push({'name': 'new job', conditions: [], actions: []})
        this.selectedJob = this.jobs[this.jobs.length-1]
        this.remotePostJob(this.selectedJob)
      },
      deleteJob(){
        this
          .remoteDeleteJob(this.selectedJob)
          .then(()=>{
            const index = this.jobs.findIndex(element => this.selectedJob.id === element.id)
            this.jobs.splice(index, 1)
            this.selectedJob = this.jobs.length>=1?this.jobs[index]?this.jobs[index]:this.jobs[index-1]:null
          })
      },
      selectJob(job){
        this.selectedJob = job
      },
      getToastSuccessConf(actionMessage){
        return {
          message: 'Sucessfully ' + actionMessage,
          type: 'is-success',
          queue: false,
          position: 'is-top-right'
        }
      }
    },
  };
</script>

<style>

</style>
