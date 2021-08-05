<template>
  <div :key="task.id" v-for="task in tasks">
    <task
      @finish-task="updateTask"
      @update-task="updateTask"
      @delete-task="$emit('delete-task', task.id)"
      :task="task"
    />
  </div>
</template>


<script>
import Task from './Task.vue';
export default {
  name: 'Tasks',
  props: {
    tasks: Array,
  },
  components: {
    Task,
  },
  emits: ['update-task','delete-task'],
  methods: {
      updateTask(newTask){
          this.$emit('update-task',newTask)
      }
  }
};
</script>
