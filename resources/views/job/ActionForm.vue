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
            <span class="has-text-weight-semibold has-text-primary">To : </span> {{to}}
          </div>
          <div class="column">
            <span class="has-text-weight-semibold has-text-primary">Name : </span> {{name}}
          </div>
          <div class="column">
            <span class="has-text-weight-semibold has-text-primary">Arguments : </span> {{args}}
          </div>
          <div class="column is-one-fifth">
            <div class="level-right">
              <b-button
                class="is-primary"
                icon-right="edit"
                @click="toggleEditable()"></b-button>
              <b-button
                type="is-danger"
                icon-right="trash"
                @click="$emit('delete:action')"/>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div
      v-if="isEditable === true">
      <b-field grouped>
        <b-field label="To">
          <b-input
            v-bind:value="to"
            @input="$emit('update:to', $event)"
            placeholder="Example : Youtrack"></b-input>
        </b-field>
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
            class="is-primary"
            icon-right="edit"
            @click="toggleEditable()"></b-button>
          <b-button
            type="is-danger"
            icon-right="trash"
            @click="$emit('delete:action')"/>
        </div>
      </b-field>
    </div>
  </div>
</template>

<script>
  export default {
    name: "actionForm",
    props: ['to', 'name', 'args', 'fillBackGround'],
    data(){
      return {
        isEditable: this.$props.name === "",
      }
    },

    methods: {
      toggleEditable(){
        this.isEditable = !this.isEditable
      }
    }
  }
</script>
