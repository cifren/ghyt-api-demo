<template>
  <div>
    <div>
      <div>
        <div class="title">{{created_at}}</div>
        <div class="section box">
          <div class="subtitle">Request</div>
          <vue-json-pretty
            :data="request"
            :highlightMouseoverNode="true"
            :showDoubleQuotes="false"
            :deep="1">
          </vue-json-pretty>
        </div>
        <div class="section box">
          <div class="subtitle">Feedback</div>
          <vue-json-pretty
            :data="feedback"
            :highlightMouseoverNode="true"
            :showDoubleQuotes="false"
            :deep="2">
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
