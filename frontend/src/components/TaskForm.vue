<template>
  <el-form :model="task" label-width="80px">
    <el-form-item label="Context">
      <el-input v-model="newTask.text" ></el-input>
    </el-form-item>

    <el-form-item label="DDL">
      <el-date-picker v-model="newTask.ddl" type="date" placeholder="choose a date">
      </el-date-picker>
    </el-form-item>


  </el-form>
</template>

<script>
export default {
  name: "TaskForm",
  data() {
    return {
      newTask: {
        ...this.task,
        text: this.task.text,
        ddl: this.task.ddl,
      },
    };
  },
  props: {
      task : Object,
  }
};
</script>

<style>
</style>