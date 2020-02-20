<template>
  <div>
    <div
      :class="{
        'box': true,
        'has-background-light': !fillBackGround
      }"
      v-if="isEditable === false">
      <div>
        <div class="columns level">
          <div class="column">
            <span class="has-text-weight-semibold has-text-info is-capitalized">{{name}}</span>
          </div>
          <div class="column">
            <div class="columns level">
              <b-dropdown hoverable aria-role="list">
                <div class="has-text-info level has-text-weight-semibold " slot="trigger">
                  <span>Arguments&nbsp</span>
                  <b-icon class="button" size="is-small" icon="info"></b-icon>
                </div>

                <b-dropdown-item
                  v-for="(arg, key) in args"
                  v-bind:key="key"
                  class="level-left"
                  aria-role="listitem">
                    <strong>{{key}}</strong> <b-icon icon="angle-double-right"/> {{arg}}
                </b-dropdown-item>
              </b-dropdown>
            </div>
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
                @click="$emit('delete:condition')"/>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div
      v-if="isEditable === true"
      :class="{
        'card': true,
        'has-background-light': !fillBackGround
      }">
      <div class="card-content">
        <b-field
          customClass="has-text-info"
          label="Name" horizontal expanded>
          <p class="control is-expanded">
            <b-input
              v-bind:value="name"
              @input="$emit('update:name', $event)"
              placeholder="Example : equal"></b-input>
          </p>
        </b-field>

        <b-field horizontal>
          <p class="control">
            <b-button
              type="is-secondary"
              @click="$emit('update:args', {...args, ...{'': ''}})"
              icon-left="plus">Argument</b-button>
          </p>
        </b-field>
        <b-field
          v-for="(arg, key) in args"
          v-bind:key="key"
          label="Argument"
          customClass="has-text-info"
          horizontal>
          <b-field addons>
            <b-input
              v-bind:value="key"
              placeholder="Key"
              @blur="validKey(key, $event.target.value)?$emit('update:args', updateArgKey(args, key, $event.target.value)):false"></b-input>
            <b-button
              class="is-static has-text-primary"
              icon-right="angle-double-right"></b-button>
            <b-input
              v-bind:value="arg"
              placeholder="Value"
              @input="$emit('update:args', {...args, ...{[key]: $event}})" expanded></b-input>
            <p class="control">
              <b-button
                type="is-danger"
                @click="$emit('update:args', deleteArgs(args, key))"
                expanded
                icon-right="trash"/>
            </p>
          </b-field>
        </b-field>
      </div>
      <footer class="card-footer">
        <a
          class="card-footer-item has-text-primary"
          @click="toggleEditable()">
          <b-icon icon="save"></b-icon>Close edition
        </a>
        <a
          class="card-footer-item has-text-danger"
          @click="$emit('delete:condition')">
          <b-icon icon="trash"></b-icon>Delete
        </a>
      </footer>
    </div>
  </div>
</template>

<script>
  export default {
    name: "conditionForm",
    props: ['id', 'name', 'args', 'fillBackGround'],
    data(){
      return {
        isEditable: this.$props.name === "",
      }
    },
    methods: {
      toggleEditable(){
        this.isEditable = !this.isEditable
      },
      updateArgKey(args, key, newKey){
        return Object.fromEntries(Object.entries(args).map(function([mapKey, mapValue]){
            if (key === mapKey) {
              return [newKey, mapValue]
            }
            return [mapKey, mapValue]
          }
        ))
      },
      deleteArgs(args, key){
        delete args[key]
        return {...args}
      },
      validKey(key, newKey){
        const args = this.$props.args

        // means the key name changed
        if(key !== newKey){
          if(newKey in args){
            return false
          }
        }

        return true
      }
    }
  }
</script>
