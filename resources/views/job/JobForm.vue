<template>
  <div class="container">
    <div>
      <b-button
        class="is-primary"
        @click="addNewCondition">Add Condition</b-button>
      <div v-for="(condition, key) in conditions">
        <condition-form
          :fillBackGround="key % 2 === 0"
          v-bind.sync="condition"
          @delete:condition="deleteCondition(key)"
        ></condition-form>
      </div>
    </div>
    <div>
      <b-button
        class="is-primary"
        @click="addNewAction">Add Action</b-button>
      <div v-for="(action, key) in actions">
        <action-form
          :fillBackGround="key % 2 === 0"
          v-bind.sync="action"
          @delete:action="deleteAction(key)"
        ></action-form>
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
    data: function() {
      return {
        conditions: [],
        actions: []
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
    }
  }
</script>
