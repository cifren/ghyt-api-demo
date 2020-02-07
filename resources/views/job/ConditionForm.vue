<template>
  <div>
    <b-field label="Name">
      <b-input
        v-bind:value="condition.name"
        v-on:input="$emit('update:condition', $event)"
        placeholder="Example : equal"></b-input>
    </b-field>
    <b-field label="Arguments in json format">
      <b-input
        v-bind:value="condition.arguments"
        placeholder="Example : {'variableName': 'event.pull_request.state', 'value': 'open'}"></b-input>
    </b-field>
  </div>
</template>

<script>

  export default {
    name: "conditionForm",
    props: ['condition'],
    methods: {
      plop($event) {
        console.log($event)
      }
    }
  }
</script>
