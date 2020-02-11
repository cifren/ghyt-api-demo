<template>
  <div class="container">
    <b-button
      class="is-primary"
      @click="addNewCondition">Add Condition</b-button>
    <div v-for="(condition, key) in conditions">
      <condition-form
        :fillBackGround="key % 2 === 0"
        v-bind.sync="condition"
        @delete:condition="deleteCondtion(key)"
      ></condition-form>
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
        let condition = {"name": "", "args": "{}"};
        this.conditions.push(condition)
      },
      deleteCondtion(key){
        this.conditions.splice(key, 1)
      },
      addNewAction() {

      },
    }
  }
</script>
