<template>
  <div class="task-container">
    <el-row justify="space-between" align="middle">
      <el-col :span="12">
        <el-row>
          <div @click="changeUpdateFormVisible()" class="task-edit">
            {{ task.text }}
          </div>
        </el-row>
        <el-row>
          <div
            @click="changeUpdateFormVisible()"
            class="task-edit ddl"
            v-if="task.ddl"
          >
            <i class="el-icon-date"> </i>
            {{ taskDDL }}
          </div>
        </el-row>
      </el-col>
      <el-col :span="2">
        <el-button
          @click="finishTask()"
          size="mini"
          :type="task.isFinish ? 'success' : null"
          :icon="task.isFinish ? 'el-icon-check' : ''"
          circle
        ></el-button>
      </el-col>
    </el-row>

    <update-task
      ref="updateForm"
      :oldTask="task"
      @update-submit="updateTask"
      @delete-submit="deleteTask"
    />
  </div>
</template>

<script>
import UpdateTask from "./UpdateTask.vue";
import { fmtDate, judgeDate } from "../assets/utils.js";

export default {
  name: "Task",
  props: {
    task: Object,
  },
  components: {
    UpdateTask,
  },

  methods: {
    finishTask() {
      this.$emit("finish-task", {
        id: this.task.id,
        isFinish: !this.task.isFinish,
      });
    },
    updateTask(newTask) {
      this.$emit("update-task", newTask);
    },
    deleteTask() {
      this.$emit("delete-task", this.task.id);
    },
    changeUpdateFormVisible() {
      this.$refs.updateForm.visible = true;
    },
  },
  computed: {
    taskDDL: function () {
      const date = judgeDate(this.task.ddl);
      if (date === this.task.ddl) {
        return fmtDate(this.task.ddl);
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
.time-footer {
  margin: 10px 0 0 0;
}
</style>