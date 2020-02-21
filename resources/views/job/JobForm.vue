<template>
  <div>
    <div class="title level">
      <div
        v-if="!isEditable"
        class="level-left">
        <div class="level-item is-capitalized">{{name}}</div>
        <b-button
          class="is-primary"
          @click="isEditable = true"
          icon-left="edit"></b-button>
      </div>
      <div v-if="isEditable">
        <b-field addons>
          <b-input
            v-bind:value="name"
            placeholder="Value"
            @input="$emit('update:name', $event)" expanded></b-input>
          <p class="control">
            <b-button
              class="is-primary"
              @click="isEditable = false"
              icon-left="save"></b-button>
          </p>
        </b-field>
      </div>
    </div>
    <div class="section">
      <div class="container">
        <div class="columns">
          <div class="column is-12">
            <b-button
              class="is-primary"
              @click="addNewCondition"
              icon-left="plus">Condition</b-button>
            <div
              v-for="(condition, key) in conditions"
              v-bind:key="condition.id">
              <condition-form
                :fillBackGround="key % 2 === 0"
                v-bind.sync="condition"
                :args.sync="condition.arguments"
                @delete:condition="deleteCondition(key)"
                @update:jsonargs="updateJsonArgs($event)"
              ></condition-form>
            </div>
          </div>
        </div>

        <div class="columns" v-if="conditions.length <= 0">
          <div class="column is-12">
            <div>No conditions found</div>
          </div>
        </div>

      </div>
    </div>

    <div class="section">
      <div class="container">
        <div class="columns">
          <div class="column is-12">
            <div>
              <b-button
                class="is-primary"
                @click="addNewAction"
                icon-left="plus">Action</b-button>
              <div
                  v-for="(action, key) in actions"
                  v-bind:key="action.id">
                <action-form
                  :fillBackGround="key % 2 === 0"
                  v-bind.sync="action"
                  :args.sync="action.arguments"
                  @delete:action="deleteAction(key)"
                ></action-form>
              </div>
            </div>
          </div>
        </div>

        <div class="columns" v-if="actions.length <= 0">
          <div class="column is-12">
            <div>No actions found</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import ConditionForm from '../../views/job/ConditionForm.vue'
  import ActionForm from '../../views/job/ActionForm.vue'

  export default {
    name: 'jobForm',
    components: {
      ConditionForm,
      ActionForm
    },
    props: ['name', 'conditions', 'actions'],
    data: function() {
      return {
        isEditable: false
      };
    },
    methods: {
      addNewCondition() {
        let condition = {"name": "", "args": ""};
        this.conditions.push(condition)
      },
      deleteCondition(key){
        this.conditions.splice(key, 1)
      },
      addNewAction() {
        let action = {"name": "","to": "", "args": ""};
        this.actions.push(action)
      },
      deleteAction(key){
        this.actions.splice(key, 1)
      },
      updateJsonArgs(event){
        this.conditions[event['props'].id]
      }
    }
  }
</script>
