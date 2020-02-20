<template>
  <div>
    <div class="title">{{title}}</div>
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
    props: ['title', 'conditions', 'actions'],
    data: function() {
      return {
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
