<template>
  <el-dialog title="Update Task" v-model="visible" custom-class="dialog">
    <el-form :model="newTask" :rules="rules" ref="taskForm" label-width="80px">
      <el-form-item label="Context" prop="text">
        <el-input v-model="newTask.text" clearable></el-input>
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
      <el-button type="primary" @click="updateTask('taskForm')"
        >SUBMIT</el-button
      >

      <el-row justify="start" align="middle" class="time-footer">
        <i class="el-icon-delete task-edit" @click="deleteTask()"></i>
        <div style="font-size: 10px">Created on {{ taskCreatedAt }}</div>
      </el-row>
    </template>
  </el-dialog>
</template>

<script>
import { fmtDate, filterSameAttr } from "../assets/utils.js";
export default {
  name: "UpdateTask",
  data() {
    return {
      visible: false,
      newTask: {
        text: this.oldTask.text,
        ddl: this.oldTask.ddl,
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
  props: {
    oldTask: Object,
  },

  methods: {
    updateTask(formName) {
      // From 参数验证
      this.$refs[formName].validate((vaild) => {
        if (vaild) {
          this.visible = false;

          const newTask = filterSameAttr({ ...this.newTask }, this.oldTask);
          newTask.id = this.oldTask.id;
          if (newTask.ddl) {
            newTask.ddl = new Date(newTask.ddl).toISOString();
          }
          this.$emit("update-submit", newTask);
        } else {
          return false;
        }
      });
    },
    deleteTask() {
      this.$confirm(
        this.oldTask.text + " will be permanently deleted",
        "Are you Sure?",
        {
          confirmButtonText: "DELETE",
          cancelButtonText: "CANCEL",
          type: "warning",
          customClass: "confirm",
        }
      )
        .then(() => {
          this.$emit("delete-submit");
        })
        .catch(() => {});
    },
  },
  computed: {
    taskCreatedAt: function () {
      return fmtDate(this.oldTask.createdAt);
    },
  },
  emits: ["update-submit", "delete-submit"],
};
</script>

<style >
.confirm {
  width: 100%;
  max-width: 418px;
}
</style>