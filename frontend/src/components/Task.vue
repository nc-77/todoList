<template>
  <div class="task-container">
    <el-row justify="space-between" align="middle">
      <el-col :span="12">
        <el-row>
          <div @click="dialogVisible = true" class="task-edit">
            {{ task.text }}
          </div>
        </el-row>
        <el-row>
          <div @click="dialogVisible = true" class="task-edit ddl" v-if="task.ddl">
            <i class="el-icon-date"> </i>
            {{ taskDDL }}
          </div>
        </el-row>
      </el-col>
      <el-col :span="2">
        <el-button
          @click="finishTask"
          size="mini"
          :type="task.isFinish ? 'success' : null"
          :icon="task.isFinish ? 'el-icon-check' : ''"
          circle
        ></el-button>
      </el-col>
    </el-row>

    <el-dialog title="Update Task" v-model="dialogVisible" width="388px">
      <task-form ref="taskForm" :task="task" />
      <template #footer>
        <el-button @click="dialogVisible = false">CANCEL</el-button>
        <el-button type="primary" @click="updateTask">SUBMIT</el-button>

        <el-row justify="start" align="middle">
          <i class="el-icon-delete task-edit" @click="deleteTask"></i>
          <div style="font-size: 10px">Created on {{ taskCreatedAt }}</div>
        </el-row>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import TaskForm from "./TaskForm.vue";
export default {
  name: "Task",
  data() {
    return {
      dialogVisible: false,
    };
  },
  props: {
    task: Object,
  },
  components: {
    TaskForm,
  },
  methods: {
    finishTask() {
      this.$emit("finish-task", {
        ...this.task,
        ddl: this.task.ddl ? new Date(this.task.ddl).toISOString() : null,
        isFinish: !this.task.isFinish,
      });
    },
    updateTask() {
      this.dialogVisible = false;
      const newTask = this.$refs.taskForm.newTask;
      newTask.ddl = newTask.ddl ? new Date(newTask.ddl).toISOString() : null;
      this.$emit("update-task", newTask);
    },
    deleteTask() {
      this.$confirm(
        this.task.text + " will be permanently deleted",
        "Are you Sure?",
        {
          confirmButtonText: "DELETE",
          cancelButtonText: "CANCEL",
          type: "warning",
        }
      )
        .then(() => {
          this.$emit("delete-task", this.task.id);
        })
        .catch(() => {});
    },
    fmtDate(dateString) {
      const date = new Date(dateString);
      const dates = date.toDateString().split(" ");
      const fmtDate = dates[0] + "," + dates[1] + " " + dates[2];
      return fmtDate;
    },
    // 判断是否是今天或者明天
    judgeDate(dateString) {
      const date = new Date(dateString);
      const today = new Date();
      const tomorrow = new Date(
        today.getFullYear(),
        today.getMonth(),
        today.getDate() + 1
      );
      if (
        date.getFullYear() == today.getFullYear() &&
        date.getMonth() == today.getMonth() &&
        date.getDate() == today.getDate()
      ) {
        return "Today";
      } else if (
        date.getFullYear() == tomorrow.getFullYear() &&
        date.getMonth() == tomorrow.getMonth() &&
        date.getDate() == tomorrow.getDate()
      ) {
        return "Tomorrow";
      }
      return dateString;
    },
  },
  computed: {
    taskCreatedAt: function () {
      return this.fmtDate(this.task.createdAt);
    },
    taskDDL: function () {
      const date = this.judgeDate(this.task.ddl);
      if (date === this.task.ddl) {
        return this.fmtDate(this.task.ddl);
      }
      return date;
    },
  },
};
</script>

<style scope>
.task-container {
  background: #f4f4f4;
  margin: 5px;
  padding: 10px 20px;
}
.el-button--mini.is-circle {
  width: 34px;
  height: 34px;
  padding: 10px;
}
.el-icon-delete {
  margin: 0 8px 0 0;
}
.task-edit {
  cursor: pointer;
}
.task-edit:hover {
  color: red;
}
.task-edit.ddl {
  font-size: 14px;
}
</style>