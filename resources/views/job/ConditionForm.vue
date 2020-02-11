<template>
  <div>
    <div
      :class="{
        'box': true,
        'has-background-light': !fillBackGround
      }"
      v-if="isEditable === false">
      <div>
        <div class="columns">
          <div class="column">
            <span class="has-text-weight-semibold has-text-primary">Name : </span> {{name}}
          </div>
          <div class="column">
            <span class="has-text-weight-semibold has-text-primary">Arguments : </span>{{args}}
          </div>
          <div class="column is-one-fifth">
            <div class="level-right">
              <b-button
                type="is-danger"
                icon-right="trash"
                @click="$emit('delete:condition')"/>
              <b-button
                class="is-primary"
                icon-right="edit"
                @click="toggleEditable()"></b-button>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div v-if="isEditable === true">
      <b-field grouped>
        <b-field label="Name">
          <b-input
            v-bind:value="name"
            @input="$emit('update:name', $event)"
            placeholder="Example : equal"></b-input>
        </b-field>
        <b-field label="Arguments in json format" expanded>
          <b-input
            v-bind:value="args"
            @input="$emit('update:args', $event)"
            placeholder="Example : {'variableName': 'event.pull_request.state', 'value': 'open'}"
            maxlength="600"
            type="textarea"></b-input>
        </b-field>
        <div class="buttons">
          <b-button
            type="is-danger"
            icon-right="trash"
            @click="$emit('delete:condition')"/>
          <b-button
            class="is-primary"
            icon-right="edit"
            @click="toggleEditable()"></b-button>
        </div>
      </b-field>
    </div>
  </div>
</template>

<script>
  export default {
    name: "conditionForm",
    props: ['name', 'args', 'fillBackGround'],
    data(){
      return {
        isEditable: false,
      }
    },
    methods: {
      toggleEditable(){
        this.isEditable = !this.isEditable
      }
    }
  }
</script>
