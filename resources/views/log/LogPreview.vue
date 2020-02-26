<template>
  <div class="job card">
    <div class="card-content">
      <div class="media">
        <div class="media-content">
          <p class="title is-capitalized">
            <b-icon icon="briefcase" size="is-large" type="is-primary"/>
            {{created_at | moment("YYYY/MM/DD HH:mm:ss")}}
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'logPreview',
    props: ['created_at']
  }
</script>

<style scoped>

</style>
