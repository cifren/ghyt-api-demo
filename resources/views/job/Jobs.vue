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
    <div class="container">

      <b-button @click="clickMe">Create new enty</b-button>

    </div>
  </div>
</template>

<script>
export default {
  name: 'home',
  mounted:function(){

  },
  methods: {
    clickMe() {
      this.$buefy.notification.open('Clicked!!')
    }
  }
}
</script>

<style>

</style>
