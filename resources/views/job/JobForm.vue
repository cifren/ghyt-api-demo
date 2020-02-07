<template>
  <div class="container">
    <b-button @click="addNewCondition">Add Condition</b-button>
    <div v-for="condition in conditions">
      <ConditionForm
        :condition.sync="condition.name"
      ></ConditionForm>
    </div>
  </div>
</template>

<script>
  import ConditionForm from '../../views/job/ConditionForm.vue'

  export default {
    name: 'jobForm',
    components: {
      ConditionForm
    },
    data: function() {
      return {
        conditions: []
      };
    },
    methods: {
      addNewCondition() {
        let condition = {"name": "plop", "arguments": "lol"};
        this.conditions.push(condition)
      },
      addNewAction() {

      }
    }
  }
</script>
