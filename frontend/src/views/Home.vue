<template>
  <el-container class="container">
    <el-header>
      <el-row justify="space-between">
        <h1>My Todo List</h1>
        <add-task-form @add-task="addTask" />
        </el-row
      >
    </el-header>
    <el-main>
      <tasks
        :tasks="tasks"
        @delete-task="deleteTask"
        @update-Task="updateTask"
      />
    </el-main>

    <el-footer>
      <el-row justify="center">
        <p>Copyright &copy; 2021 NiC</p>
      </el-row>
    </el-footer>
  </el-container>
</template>

<script>
import Tasks from "../components/Tasks.vue";
import AddTaskForm from "../components/AddTask.vue";

export default {
  name: "Home",
  components: {
    Tasks,
    AddTaskForm,
  },
  data() {
    return {
      tasks: [],
      showAddTask: false,
    };
  },
  async created() {
    this.tasks = await this.fetchData();
  },

  methods: {
    async addTask(newTask) {
      const res = await fetch("api/tasks", {
        method: "PUT",
        body: JSON.stringify(newTask),
      });
      const result = await res.json();
      this.refresh(result.code);
    },

    async deleteTask(id) {
      const res = await fetch("api/tasks", {
        method: "DELETE",
        body: JSON.stringify({ id: [id] }),
      });
      const result = await res.json();
      this.refresh(result.code);
    },

    async updateTask(newTask) {
      const res = await fetch("api/tasks", {
        method: "POST",
        body: JSON.stringify(newTask),
      });

      const result = await res.json();
      this.refresh(result.code);
    },

    async fetchData() {
      const res = await fetch("api/tasks");
      const resData = await res.json();
      const tasks = resData.data;
      tasks.sort(function (a, b) {
        if (a.isFinish === b.isFinish) {
          if (a.isFinish) {
            return a.ddl < b.ddl ? 1 : -1;
          } else {
            return a.ddl > b.ddl ? 1 : -1;
          }
        } else {
          return a.isFinish ? 1 : -1;
        }
      });

      return tasks;
    },
    async refresh(httpCode) {
      switch (httpCode) {
        case 200:
          this.tasks = await this.fetchData();
          break;
        default:
          alert("出错了,请联系管理员");
      }
    },
  },
};
</script>

<style scoped>
@media screen and (min-width: 992px) {
  .container {
    max-width: 500px;
    margin: 30px auto;
    overflow: auto;
    min-height: 300px;
    border: 1px solid steelblue;
    padding: 30px 30px 10px 30px;
    border-radius: 5px;
  }
}
@media screen and (max-width: 991px) {
  .container {
    max-width: 500px;
    margin: 30px auto;
    overflow: auto;
    min-height: 300px;
    padding: 30px 30px 10px 30px;
    border-radius: 5px;
  }
}
.el-header {
  padding: 0;
}
.el-main {
  padding: 0;
}
.el-footer {
  height: 20px;
  font-size: 8px;
}
</style>