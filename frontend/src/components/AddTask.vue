<template>
  <el-button type="success" @click="visible = true">Add Task</el-button>
  <el-dialog
    title="Add Task"
    v-model="visible"
    @close="reset('taskForm')"
    custom-class="dialog"
  >
    <el-form :model="newTask" :rules="rules" ref="taskForm" label-width="80px">
      <el-form-item label="Context" prop="text">
        <el-input v-model="newTask.text"></el-input>
      </el-form-item>

      <el-form-item label="DDL">
        <el-date-picker
          v-model="newTask.ddl"
          type="date"
          placeholder="choose a date"
        >
        </el-date-picker>
      </el-form-item>
    </el-form>

    <template #footer>
      <el-button @click="visible = false">CANCEL</el-button>
      <el-button type="primary" @click="addTask('taskForm')">SUBMIT</el-button>
    </template>
  </el-dialog>
</template>

<script>
import { filterEmptyAttr } from "../assets/utils.js";
export default {
  name: "AddTaskForm",
  data() {
    return {
      visible: false,
      newTask: {
        text: "",
        ddl: "",
      },
      rules: {
        text: [
          {
            required: true,
            message: "Please enter the task content",
            trigger: "blur",
          },
          { max: 50, message: "No more than 50 words", trigger: "blur" },
        ],
      },
    };
  },
  methods: {
    addTask(formName) {
      // From 参数验证
      this.$refs[formName].validate((vaild) => {
        if (vaild) {
          this.visible = false;
          this.newTask.ddl = this.newTask.ddl
            ? new Date(this.newTask.ddl).toISOString()
            : null;

          this.$emit("add-task", filterEmptyAttr(this.newTask));
        } else {
          return false;
        }
      });
    },
    reset(formName) {
      this.$refs[formName].resetFields();
    },
  },
  emits: ["add-task"],
};
</script>

<style>
.dialog {
  width: 100%;
  min-width: 360px;
  max-width: 400px;
}
</style>