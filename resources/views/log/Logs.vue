<template>
  <div>
    <section class="hero is-primary">
      <div class="hero-body">
        <div class="container">
          <h1 class="title">
            Logs
          </h1>
          <h2 class="subtitle">
            My Logs are here
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
                  <div class="log-list">
                    <log-preview
                      v-for="log in logs"
                      v-bind:key="log.id"
                      v-bind="log"
                      @click.native="selectLog(log)"
                      :class="{
                        'active': isActiveLog(log),
                        'is-loading': isLoadingLogs
                      }"
                    ></log-preview>
                    <b-loading :is-full-page="false" :active.sync="isLoadingLogs" :can-cancel="true"></b-loading>
                  </div>
                </div>
              </div>
            </div>
            <div
              :class="{
                'column is-8 hero is-fullheight': true,
                'is-hidden': !hasActiveLog()
              }">
              <div class="box">
                <log-details
                  v-if="hasActiveLog()"
                  v-bind.sync="selectedLog"
                ></log-details>
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
  import LogPreview from './LogPreview'
  import LogDetails from './LogDetails'
  let ghytConfApi = process.env.GHYT_CONF_API;
  export default {
    name: 'logs',
    components: {LogDetails, LogPreview},
    data: function() {
      return {
        selectedLog: null,
        logs: [],
        isLoadingLogs: true,
        errorMessage: null
      };
    },
    mounted() {
      if(ghytConfApi === undefined){
        this.errorMessage = "Error : The API url is missing"
        return
      }

      this.axios.get(ghytConfApi + '/logs')
      .then(response => {
        this.isLoadingLogs = false
        response.data = response.data.map((element) => {
          element.created_at = new Date(element.created_at)
          return element
        })

        this.logs = response.data.sort((a, b) => b.created_at - a.created_at)
      })
      .catch(error => {
        this.isLoadingLogs = false
        this.errorMessage = error
      });
    },
    methods: {
      hasActiveLog(){
        return this.selectedLog !== null
      },
      isActiveLog(log){
        return this.selectedLog !== null && log.id === this.selectedLog.id
      },
      selectLog(log){
        this.selectedLog = log
      },
    }
  }
</script>

<style scoped>

</style>
