<template>
  <div class="job card">
    <div class="card-content">
      <div class="media">
        <div class="media-content">
          <p class="title">
            <b-icon icon="briefcase" size="is-large" type="is-primary"/>
            {{title}}
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'job',
    props: ['title']
  };
</script>

<style>

</style>
