<template>
  <div>
    <div>
      <div>
        <div class="title">Ran at
          <span class="has-text-primary">{{created_at | moment("YYYY/MM/DD h:mm:ss")}}</span>
        </div>
        <div class="section box">
          <div class="subtitle has-text-primary has-text-weight-bold">Request</div>
          <vue-json-pretty
            :data="request"
            :highlightMouseoverNode="true"
            :showDoubleQuotes="false"
            :deep="1">
          </vue-json-pretty>
        </div>
        <div class="section box">
          <div class="subtitle has-text-primary has-text-weight-bold">Feedback</div>
          <vue-json-pretty
            :data="feedback"
            :highlightMouseoverNode="true"
            :showDoubleQuotes="false"
            :deep="4">
          </vue-json-pretty>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import VueJsonPretty from 'vue-json-pretty'
  export default {
    name: 'logDetails',
    components: {
      VueJsonPretty
    },
    props: ['created_at', 'request', 'feedback']
  }
</script>

<style scoped>

</style>
